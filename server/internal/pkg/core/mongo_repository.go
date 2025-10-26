package core

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"maps"
// 	"strings"
// 	"time"

// 	"github.com/aritradevelops/authinfinity/server/internal/pkg/db"
// 	"go.mongodb.org/mongo-driver/v2/bson"
// )

// // implements Repository
// type MongoRepository[S Schema] struct {
// 	model Model
// 	db    *db.MongoDb
// }

// func NewMongoRepository[S Schema](model Model, database db.Database) Repository[S] {
// 	mongoDb, ok := database.(*db.MongoDb)
// 	if !ok {
// 		panic("Invalid db.Database passed to NewMongoRepository, it only accepts *db.MongoDB")
// 	}
// 	return &MongoRepository[S]{
// 		model: model,
// 		db:    mongoDb,
// 	}
// }

// func (r *MongoRepository[S]) List(opts *ListOpts) (*PaginatedResponse[S], error) {
// 	if opts == nil {
// 		opts = &ListOpts{}
// 	}
// 	sanitizeFilter(opts.Filters)
// 	searchableFields := r.model.Searchables()

// 	matchStage := bson.M{}
// 	maps.Copy(matchStage, opts.Filters)
// 	if len(searchableFields) > 0 && opts.Search != "" {
// 		var orConditions []bson.M
// 		for _, field := range searchableFields {
// 			orConditions = append(orConditions, bson.M{
// 				field: bson.M{"$regex": opts.Search, "$options": "i"},
// 			})
// 		}
// 		if len(orConditions) > 0 {
// 			matchStage["$or"] = orConditions
// 		}
// 	}

// 	pipeline := []bson.M{}
// 	if len(matchStage) > 0 {
// 		pipeline = append(pipeline, bson.M{"$match": matchStage})
// 	}

// 	sortOpts := bson.M{opts.SortBy: opts.SortOrderInt()}
// 	if opts.SortBy != "created_at" {
// 		sortOpts["created_at"] = -1
// 	}
// 	pipeline = append(pipeline, bson.M{"$sort": sortOpts})

// 	skip := (opts.Page - 1) * opts.PerPage
// 	pipeline = append(pipeline,
// 		bson.M{"$skip": skip},
// 		bson.M{"$limit": opts.PerPage},
// 	)

// 	if opts.Select != "" {
// 		fields := strings.Split(opts.Select, ",")
// 		projectOpts := bson.M{}
// 		for _, f := range fields {
// 			f = strings.TrimSpace(f)
// 			if f != "" {
// 				projectOpts[f] = 1
// 			}
// 		}
// 		pipeline = append(pipeline, bson.M{"$project": projectOpts})
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	// Build count pipeline (same as main, but replace after $match with $count)
// 	countPipeline := make([]bson.M, 0, len(pipeline))
// 	for _, stage := range pipeline {
// 		// Copy until first $sort/$skip/$limit/$project
// 		k := func(m bson.M) string {
// 			for key := range m {
// 				return key
// 			}
// 			return ""
// 		}(stage)
// 		if k == "$sort" || k == "$skip" || k == "$limit" || k == "$project" {
// 			break
// 		}
// 		countPipeline = append(countPipeline, stage)
// 	}
// 	countPipeline = append(countPipeline, bson.M{"$count": "total"})

// 	countCursor, err := r.db.Collection(r.model.Name()).Aggregate(ctx, countPipeline)
// 	if err != nil {
// 		return nil, err
// 	}
// 	total := 0
// 	if countCursor.Next(ctx) {
// 		var countDoc struct{ Total int }
// 		if err := countCursor.Decode(&countDoc); err != nil {
// 			return nil, err
// 		}
// 		total = countDoc.Total
// 	}
// 	countCursor.Close(ctx)
// 	fmt.Printf("%#v\n", pipeline)
// 	cursor, err := r.db.Collection(r.model.Name()).Aggregate(ctx, pipeline)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)
// 	docs := []S{}
// 	if err := cursor.All(ctx, &docs); err != nil {
// 		return nil, err
// 	}

// 	from, to := 0, 0
// 	if len(docs) > 0 {
// 		from = skip + 1
// 		to = min(skip+len(docs), total)
// 	}

// 	return &PaginatedResponse[S]{
// 		Data: docs,
// 		Info: PaginationInfo{
// 			From:  from,
// 			To:    to,
// 			Total: total,
// 		},
// 	}, nil
// }

// func (r *MongoRepository[S]) Create(data S) (string, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()
// 	result, err := r.db.Collection(r.model.Name()).InsertOne(ctx, data)

// 	if err != nil {
// 		return "", err
// 	}
// 	oid, ok := result.InsertedID.(bson.ObjectID)
// 	if !ok {
// 		return "", errors.New("inserted ID is not ObjectID")
// 	}
// 	return oid.Hex(), nil
// }

// func (r *MongoRepository[S]) Update(filter Filter, update S) (bool, error) {
// 	sanitizeFilter(filter)
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()
// 	result, err := r.db.Collection(r.model.Name()).UpdateOne(ctx, filter, bson.M{"$set": update})
// 	if err != nil {
// 		return false, err
// 	}
// 	return result.MatchedCount > 0 && result.ModifiedCount > 0, nil
// }

// func (r *MongoRepository[S]) Delete(filter Filter) (bool, error) {
// 	sanitizeFilter(filter)
// 	result, err := r.db.Collection(r.model.Name()).DeleteOne(context.Background(), filter)
// 	if err != nil {
// 		return false, err
// 	}
// 	return result.Acknowledged, nil
// }
// func (r *MongoRepository[S]) View(filter Filter, container *S) error {
// 	sanitizeFilter(filter)
// 	fmt.Printf("%+v", filter)
// 	result := r.db.Collection(r.model.Name()).FindOne(context.Background(), filter)
// 	fmt.Printf("%+v", result)
// 	err := result.Decode(&container)

// 	fmt.Printf("%+v", container)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func sanitizeFilter(f Filter) {
// 	if f == nil {
// 		return
// 	}
// 	if val, ok := f["id"]; ok {
// 		if str, ok := val.(string); ok {
// 			if oID, err := bson.ObjectIDFromHex(str); err == nil {
// 				f["_id"] = oID
// 			}
// 		}
// 		delete(f, "id")
// 	}
// }
