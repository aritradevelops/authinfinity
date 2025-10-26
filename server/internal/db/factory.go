package db

type DatabaseKind int

const (
	KindPostgres DatabaseKind = 0
	KindMongoDb  DatabaseKind = 1
)

var instance Database

func New(uri string, kind DatabaseKind) Database {
	switch kind {
	case 0:
		{
			instance = NewPostgres(uri)
		}
	case 1:
		{
			// instance = NewMongoDb(uri)
		}
	}
	return instance
}

func Instance() Database {
	return instance
}
