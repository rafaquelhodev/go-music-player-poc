package drawer

import (
	"fmt"
	"strings"

	"github.com/rafaquelhodev/go-sound-player/internal/options"
)

type Drawer struct {
	SoundOpts   *options.Options
	beatWidth   int
	beatHeight  int
	beatSpacing int
}

func NewDrawer(soundOpts *options.Options) *Drawer {
	return &Drawer{SoundOpts: soundOpts, beatWidth: 12, beatHeight: 6, beatSpacing: 4}
}

func reset() {
	fmt.Print("\033[2J\033[1;1H")

}

func (drw *Drawer) insertSquare(matrix *[][]string, sqr *Square) []int {
	var lastJ int

	initialI, initialJ := sqr.InitPosition[0], sqr.InitPosition[1]

	height := drw.beatHeight
	width := drw.beatWidth

	if !sqr.IsBeat {
		initialI += 2
		initialJ += 2
		height = height / 2
		width = width / 2
	}

	for i := initialI; i < height+initialI; i++ {
		for j := initialJ; j < width+initialJ; j++ {
			if sqr.BeingPlayed || j == initialJ ||
				j == width+initialJ-1 ||
				i == initialI ||
				i == height+initialI-1 {
				(*matrix)[i][j] = "."
			} else {
				(*matrix)[i][j] = " "
			}
			lastJ = j

		}
	}

	return []int{0, lastJ + 1}
}

func (drw *Drawer) addSpacing(matrix *[][]string, sqr *Square) []int {
	var lastJ int
	for i := sqr.InitPosition[0]; i < drw.beatHeight+sqr.InitPosition[0]; i++ {
		for j := sqr.InitPosition[1]; j < drw.beatSpacing+sqr.InitPosition[1]; j++ {
			(*matrix)[i][j] = " "
			lastJ = j
		}
	}
	return []int{0, lastJ + 1}

}

func (drw *Drawer) initializeMatrix(totalSquares int) [][]string {
	matrix := make([][]string, drw.beatHeight)
	for i := range matrix {
		matrix[i] = make([]string, totalSquares*(drw.beatWidth+drw.beatSpacing))
		for j := range matrix[i] {
			matrix[i][j] = " "
		}
	}
	return matrix
}

func (drw *Drawer) Draw(beepNumber int) {
	var sb strings.Builder

	totalSquares := *drw.SoundOpts.Beats + 1 + *drw.SoundOpts.Beats*(*drw.SoundOpts.Subdivisions-1)

	matrix := drw.initializeMatrix(totalSquares)

	nextSquarePosition := []int{0, 0}
	isBeat := true
	subCount := 0
	for sqr := 1; sqr <= totalSquares; sqr++ {
		square := NewSquare(isBeat, sqr == beepNumber, nextSquarePosition)
		nextSquarePosition = drw.insertSquare(&matrix, square)

		square.UpdatePosition(nextSquarePosition)
		nextSquarePosition = drw.addSpacing(&matrix, square)

		square.UpdatePosition(nextSquarePosition)

		if *drw.SoundOpts.Subdivisions > 1 {
			if isBeat {
				isBeat = false
				subCount = 1
			} else if subCount == *drw.SoundOpts.Subdivisions-1 {
				isBeat = true
				subCount = 0
			} else {
				subCount += 1
			}
		}
	}

	reset()

	for i := range matrix {
		rowStr := strings.Join(matrix[i], "")
		sb.WriteString(rowStr)
		sb.WriteString("\n")
	}

	fmt.Print(sb.String())
}
