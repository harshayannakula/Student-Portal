package internal

type Company struct {
	id     int
	name   string
	drives []*Drive
}

func (c *Company) ID() int {
	return c.id
}

func (c *Company) Name() string {
	return c.name
}

func (c *Company) Drives() []*Drive {
	return c.drives
}

func (c *Company) AddDrive(drive *Drive) {
	c.drives = append(c.drives, drive)
}
