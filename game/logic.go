package game

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const EasyLvlDimension = 9
const EasyLvlBombsNumber = 10
const MediumLvlDimension = 16
const MediumLvlBombsNumber = 40
const HardLvlDimension = 30
const HardLvlBombsNumber = 116

type point struct {
	touched     bool
	isBomb      bool
	bombsNumber int
	hasFlag     bool
}

// Board represents minesweeper board.
type Board struct {
	bombsNumber int
	dimension   int
	field       [][]*point
	gameOver    bool
	gameWin     bool
}

func (p *point) toString() string {
	return " " + strconv.FormatBool(p.isBomb) + " neighbours " + strconv.Itoa(p.bombsNumber)
}

func (b *Board) setBoard(n int) {
	for i := 0; i < n; i++ {
		row := []*point{}
		for j := 0; j < n; j++ {
			row = append(row, new(point))
		}
		b.field = append(b.field, row)
	}
	b.setBombs()
	b.setBombsNeighbours()
}

func (b *Board) setBombs() {
	rand.Seed(time.Now().UTC().UnixNano())
	count := b.bombsNumber
	for count > 0 {
		x := rand.Intn(b.dimension)
		y := rand.Intn(b.dimension)
		if b.field[x][y].isBomb == false {
			b.field[x][y].isBomb = true
			count--
		}
	}
}

func (b *Board) setBombsNeighbours() {
	for i := 0; i < b.dimension; i++ {
		for j := 0; j < b.dimension; j++ {
			coords := []int{-1, 0, 1}
			for _, ki := range coords {
				for _, kj := range coords {
					if ki == 0 && kj == 0 {
						continue
					} else if ((ki+i >= 0) && (ki+i < b.dimension)) &&
						((kj+j >= 0) && (kj+j < b.dimension)) {
						if b.field[ki+i][kj+j].isBomb == true {
							b.field[i][j].bombsNumber++
						}
					}
				}
			}
		}
	}
	//for i := 0; i < b.dimension; i++{
	//	for j := 0; j < b.dimension; j++{
	//		if (b.field[i][j].isBomb) {
	//			fmt.Print("*"+" ")
	//		} else {
	//			fmt.Print(strconv.Itoa(b.field[i][j].bombsNumber)+" ")
	//		}
	//	}
	//	fmt.Println()
	//}
}

func (b *Board) performRightClick(colCoord int, rowCoord int) {
	newBoardState := b.field
	if newBoardState[colCoord][rowCoord].touched == true {
	} else {
		newBoardState[colCoord][rowCoord].hasFlag = true
	}
	b.updateState(newBoardState)
}

func (b *Board) performLeftClick(colCoord int, rowCoord int) {
	newBoardState := b.field
	if newBoardState[colCoord][rowCoord].touched == true {
	} else {
		if newBoardState[colCoord][rowCoord].isBomb == true {
			//TODO:make visible all bomb points
			newBoardState[colCoord][rowCoord].touched = true
			b.updateState(newBoardState)
			b.gameOver = true
			b.showAllBombs()
		} else {
			bombs := newBoardState[colCoord][rowCoord].bombsNumber
			if bombs > 0 {
				newBoardState[colCoord][rowCoord].touched = true
				b.updateState(newBoardState)
				if b.isWin() == true {
					b.gameWin = true
				}
			} else {
				//empty amount of bomb neighbours
				newBoardState[colCoord][rowCoord].touched = true
				coords := []int{-1, 0, 1}
				for _, ki := range coords {
					for _, kj := range coords {
						if ki == 0 && kj == 0 {
							continue
						} else if ((ki+colCoord >= 0) && (ki+colCoord < b.dimension)) &&
							((kj+rowCoord >= 0) && (kj+rowCoord < b.dimension)) {
							b.performLeftClick(ki+colCoord, kj+rowCoord)
						}
					}
				}
			}
		}

	}
}

func (b *Board) updateState(newBoard [][]*point) {
	b.field = newBoard
}

func (b *Board) isWin() bool {
	for i := 0; i < b.dimension; i++ {
		for j := 0; j < b.dimension; j++ {
			if (b.field[i][j].isBomb == false) &&
				(b.field[i][j].touched == false) {
				return false
			}
		}
	}
	return true
}

// continue tells whether we should keep playing.
func (b *Board) continuePlaying() bool {
	// Stub. TODO: Implement
	return !b.gameOver
}

func (b *Board) showBoard() {
	for i := 0; i < b.dimension; i++ {
		for j := 0; j < b.dimension; j++ {
			if b.field[i][j].touched {
				if b.field[i][j].isBomb {
					fmt.Print("x" + " ")

				} else {
					fmt.Print(strconv.Itoa(b.field[i][j].bombsNumber) + " ")

				}
			} else {
				fmt.Print("*" + " ")
			}
		}
		fmt.Println()
	}
}

func (b *Board) showAllBombs() {
	for i := 0; i < b.dimension; i++ {
		for j := 0; j < b.dimension; j++ {
			if b.field[i][j].isBomb {
				b.field[i][j].touched = true
			}
		}
	}
}

func (b *Board) initGame() {
	b.dimension = EasyLvlDimension
	b.bombsNumber = EasyLvlBombsNumber
	b.setBoard(b.dimension)
	fmt.Println(strconv.FormatBool(b.gameOver))
}

func (b *Board) choose(x, y int) {}

//func main() {
//	b := Board {dimension:EasyLvlDimension,
//		   bombsNumber: EasyLvlBombsNumber, }
//	b.setBoard(b.dimension)
//	fmt.Println(strconv.FormatBool(b.gameOver))
//	b.showBoard()
//	fmt.Println()
//	fmt.Println()
//	count := 0
//	for b.continuePlaying() {
//		b.performLeftClick(rand.Intn(b.dimension),rand.Intn(b.dimension))
//		b.showBoard()
//		fmt.Println()
//		fmt.Println()
//		count++
//	}
//	fmt.Println(count)
//}
