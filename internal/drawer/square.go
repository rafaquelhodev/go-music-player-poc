package drawer

type Square struct {
	IsBeat       bool
	InitPosition []int
	BeingPlayed  bool
}

func NewSquare() *Square {
	return &Square{IsBeat: true, InitPosition: []int{0, 0}, BeingPlayed: false}
}

func (sqr *Square) UpdatePosition(pos []int) {
	sqr.InitPosition = pos
}
