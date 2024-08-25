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
	return &Drawer{SoundOpts: soundOpts, beatWidth: 10, beatHeight: 5, beatSpacing: 4}
}

func reset() {
	fmt.Print("\033[2J\033[1;1H")

}

func (drw *Drawer) insertSquare(matrix *[][]string, sqr *Square) []int {
	var lastJ int
	for i := sqr.InitPosition[0]; i < drw.beatHeight+sqr.InitPosition[0]; i++ {
		for j := sqr.InitPosition[1]; j < drw.beatWidth+sqr.InitPosition[1]; j++ {
			if sqr.BeingPlayed || j == sqr.InitPosition[1] ||
				j == drw.beatWidth+sqr.InitPosition[1]-1 ||
				i == sqr.InitPosition[0] ||
				i == drw.beatHeight+sqr.InitPosition[0]-1 {
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

func (drw *Drawer) initializeMatrix() [][]string {
	totalSquares := *drw.SoundOpts.Beats * *drw.SoundOpts.Subdivisions

	matrix := make([][]string, drw.beatHeight)
	for i := range matrix {
		matrix[i] = make([]string, totalSquares*(drw.beatWidth+drw.beatSpacing))
	}
	return matrix
}

func (drw *Drawer) Draw() {
	var sb strings.Builder

	matrix := drw.initializeMatrix()

	totalSquares := *drw.SoundOpts.Beats * *drw.SoundOpts.Subdivisions

	square := NewSquare()

	for sqr := 0; sqr < totalSquares; sqr++ {
		nextPosition := drw.insertSquare(&matrix, square)

		square.UpdatePosition(nextPosition)
		nextPosition = drw.addSpacing(&matrix, square)

		square.UpdatePosition(nextPosition)
	}

	reset()

	for i := range matrix {
		rowStr := strings.Join(matrix[i], "")
		sb.WriteString(rowStr)
		sb.WriteString("\n")
	}

	fmt.Print(sb.String())
}
