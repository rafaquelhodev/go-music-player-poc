package drawer

type Square struct {
	IsBeat       bool
	InitPosition []int
	BeingPlayed  bool
}

func NewSquare(isBeat bool, isBeingPlayed bool, initialPosition []int) *Square {
	return &Square{IsBeat: isBeat, InitPosition: initialPosition, BeingPlayed: isBeingPlayed}
}

func (sqr *Square) UpdatePosition(pos []int) {
	sqr.InitPosition = pos
}
