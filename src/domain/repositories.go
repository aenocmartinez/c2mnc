package domain

type FieldRepository interface {
	CreateField(field Field) error
	UpdateField(field Field) error
	FindField(idField int64) (Field, error)
	FieldList() []Field
}
