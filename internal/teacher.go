package internal

type Teacher struct {
	id   int
	Name string
}

func (t Teacher) ID() int { return t.id }
