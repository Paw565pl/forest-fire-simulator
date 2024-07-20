package forest

type Tree struct {
	isAlive bool
}

func NewTree() *Tree {
	return &Tree{
		isAlive: true,
	}
}

func (t *Tree) String() string {
	if t.isAlive {
		return "🌳"
	}
	return "☠️"
}
