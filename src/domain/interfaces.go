package domain

type Field interface {
	SetId(id int64)
	Id() int64
	SetName(name string)
	Name() string
	SetDescription(description string)
	Description() string
	SetAbbreviation(abbreviation string)
	Abbreviation() string
	IsComposite() bool
	Create() error
	Update() error
}
