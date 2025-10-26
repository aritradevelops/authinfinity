package core

type Model interface {
	Name() string
	Searchables() []string
}

type BaseModel struct {
	name        string
	searchables []string
}

func NewModel(name string, searchables []string) Model {
	return &BaseModel{
		name:        name,
		searchables: searchables,
	}
}

func (m *BaseModel) Name() string {
	return m.name
}
func (m *BaseModel) Searchables() []string {
	return m.searchables
}
