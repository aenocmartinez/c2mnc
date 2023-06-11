package domain

type SingleField struct {
	id           int64
	name         string
	description  string
	abbreviation string
}

func (c *SingleField) SetId(id int64) {
	c.id = id
}

func (c *SingleField) Id() int64 {
	return c.id
}

func (c *SingleField) SetName(name string) {
	c.name = name
}

func (c *SingleField) Name() string {
	return c.name
}

func (c *SingleField) SetDescription(description string) {
	c.description = description
}

func (c *SingleField) Description() string {
	return c.description
}

func (c *SingleField) SetAbbreviation(abbreviation string) {
	c.abbreviation = abbreviation
}

func (c *SingleField) Abbreviation() string {
	return c.abbreviation
}

func (c *SingleField) IsComposite() bool {
	return false
}

func (c *SingleField) Create() error {
	return nil
}

func (c *SingleField) Update() error {
	return nil
}
