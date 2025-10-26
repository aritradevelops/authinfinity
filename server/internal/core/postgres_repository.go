package core

import (
	"fmt"
	"strings"

	"github.com/aritradevelops/authinfinity/server/internal/db"
	"gorm.io/gorm"
)

type PostgresRepository[S Schema] struct {
	model Model
	db    *db.Postgres // wrapper you define around *gorm.DB
}

func NewPostgresRepository[S Schema](model Model, database db.Database) Repository[S] {
	fmt.Println(database)
	gormDb, ok := database.(*db.Postgres)
	if !ok {
		panic("Invalid db.Database passed to NewPostgresRepository, it only accepts *db.Postgres")
	}
	return &PostgresRepository[S]{
		model: model,
		db:    gormDb,
	}
}

func (r *PostgresRepository[S]) List(opts *ListOpts) (*PaginatedResponse[S], error) {
	if opts == nil {
		opts = &ListOpts{}
	}
	if opts.PerPage == 0 {
		opts.PerPage = 20
	}
	if opts.Page == 0 {
		opts.Page = 1
	}

	tx := r.db.Db().Model(new(S))

	// Filters
	for k, v := range opts.Filters {
		tx = tx.Where(fmt.Sprintf("%s = ?", k), v)
	}

	// Search
	if opts.Search != "" && len(r.model.Searchables()) > 0 {
		conds := []string{}
		args := []any{}
		for _, field := range r.model.Searchables() {
			conds = append(conds, fmt.Sprintf("%s ILIKE ?", field)) // use LIKE if MySQL
			args = append(args, "%"+opts.Search+"%")
		}
		tx = tx.Where(strings.Join(conds, " OR "), args...)
	}

	// Count
	var total int64
	if err := tx.Count(&total).Error; err != nil {
		return nil, err
	}

	// Sorting
	sortBy := opts.SortBy
	if sortBy == "" {
		sortBy = "created_at"
	}
	sortOrder := opts.SortOrder
	if sortOrder == "" {
		sortOrder = "DESC"
	}
	tx = tx.Order(fmt.Sprintf("%s %s", sortBy, sortOrder))
	if sortBy != "created_at" {
		tx = tx.Order("created_at DESC")
	}

	// Select fields
	if opts.Select != "" {
		tx = tx.Select(strings.Split(opts.Select, ","))
	}

	// Pagination
	offset := (opts.Page - 1) * opts.PerPage
	tx = tx.Limit(opts.PerPage).Offset(offset)

	// Execute
	var docs []S
	if err := tx.Find(&docs).Error; err != nil {
		return nil, err
	}

	from, to := 0, 0
	if len(docs) > 0 {
		from = offset + 1
		to = min(offset+len(docs), int(total))
	}

	return &PaginatedResponse[S]{
		Data: docs,
		Info: PaginationInfo{
			From:  from,
			To:    to,
			Total: int(total),
		},
	}, nil
}

func (r *PostgresRepository[S]) Create(data S) (string, error) {
	if err := r.db.Db().Create(&data).Error; err != nil {
		return "", err
	}

	return data.GetID(), nil
}

func (r *PostgresRepository[S]) Update(filter Filter, update S) (bool, error) {
	tx := r.db.Db().Model(new(S))
	BuildPostgresFilter(tx, filter)
	result := tx.Updates(update)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

func (r *PostgresRepository[S]) Delete(filter Filter) (bool, error) {
	tx := r.db.Db().Model(new(S))
	BuildPostgresFilter(tx, filter)
	result := tx.Delete(new(S))
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

func (r *PostgresRepository[S]) View(filter Filter, container *S) error {
	tx := r.db.Db().Model(new(S))
	BuildPostgresFilter(tx, filter)
	if err := tx.First(container).Error; err != nil {
		return err
	}
	return nil
}

func BuildPostgresFilter(tx *gorm.DB, filter Filter) {
	for k, v := range filter {
		if v == nil {
			tx.Where(fmt.Sprintf("%s IS NULL", k))
			continue
		}
		tx = tx.Where(fmt.Sprintf("%s = ?", k), v)
	}
}
