package domain

type Subfield struct {
	field Field
	order int
}

func NewSubfield(field Field, order int) *Subfield {
	return &Subfield{
		field: field,
		order: order,
	}
}
