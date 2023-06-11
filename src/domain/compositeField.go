package domain

type CompositeField struct {
	id           int64
	name         string
	description  string
	abbreviation string
	subfields    []Subfield
}

func NewCompositeField(name string) *CompositeField {
	return &CompositeField{
		name:      name,
		subfields: []Subfield{},
	}
}

func (c *CompositeField) SetId(id int64) {
	c.id = id
}

func (c *CompositeField) Id() int64 {
	return c.id
}

func (c *CompositeField) SetName(name string) {
	c.name = name
}

func (c *CompositeField) Name() string {
	return c.name
}

func (c *CompositeField) SetDescription(description string) {
	c.description = description
}

func (c *CompositeField) Description() string {
	return c.description
}

func (c *CompositeField) SetAbbreviation(abbreviation string) {
	c.abbreviation = abbreviation
}

func (c *CompositeField) Abbreviation() string {
	return c.abbreviation
}

func (c *CompositeField) IsComposite() bool {
	return true
}

func (c *CompositeField) Create() error {
	return nil
}

func (c *CompositeField) Update() error {
	return nil
}

func (c *CompositeField) AddField(field Field, order int) {
	c.subfields = append(c.subfields, *NewSubfield(field, order))
}

func (c *CompositeField) Subfields() []Subfield {
	return c.subfields
}
