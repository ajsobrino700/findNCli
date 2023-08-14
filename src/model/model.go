package model

import (
	"findNGame/color"
	"fmt"
	"log"
	"strings"

	"github.com/pkg/term"
)

type Game struct {
	N         int
	SizeTable int
	table     [][]int
	turn      bool
}

func (this *Game) Init() {
	table := make([][]int, this.SizeTable)
	for i := 0; i < this.SizeTable; i++ {
		row := make([]int, this.SizeTable)
		table[i] = row
	}
	this.table = table

  var row, column int=-1,-1
  var player string

  for i:=0;!this.check(row,column);i++{
      if this.turn {
        player = "two"
      }else{
        player = "one"
      }
      fmt.Println("The turn of player",player)
      row,column = this.fillCell(this.runTurn())
  }

  this.Print()
  fmt.Println("The winner is the player",player)

}

func (t *Game) fillCell(column int)(int,int) {
	for i := 0; i < t.SizeTable; i++ {
		if ((t.SizeTable-1) == i && t.table[i][column] == 0) || (t.SizeTable-1 != i && t.table[i+1][column]!= 0) && t.table[i][column]==0{
      t.table[i][column] = t.getValuTurn()
      return i,column
		}
	}
  return -1,-1
}

func (this *Game) Print() {
	for i := 0; i < this.SizeTable; i++ {
    fmt.Print("|")
		for j := 0; j < this.SizeTable; j++ {
        fmt.Printf("%s|", getIcon(this.table[i][j]))
		}
		fmt.Println(strings.Repeat(" ",15))
	}

}

func (t *Game) getValuTurn()int{
  var result int = -1
  if t.turn{
    result = 1
  }
  return result
}


func getIcon(value int)string{
  var result string
  if value < 0 {
   result = color.Red("\u25A0")
  }else if value > 0 {
    return color.Yellow("\u00BA")
  }else{
   result = " "
  }
  return result
}


func (this *Game) check(row int, column int) bool {

  if row == -1 || column == -1 {
    return false
  }
  value := this.table[row][column]
  verticalAlign := this.vertical(row,column,value)+1
  horizontalAlign := this.horizontal(row,column,value)
  mainDiagonal := this.mainDiagonal(row,column, value)
  inverseDiagonal := this.inverseDiagonal(row,column,value)
	return verticalAlign == this.N || horizontalAlign == this.N || mainDiagonal == this.N || inverseDiagonal == this.N
}


func (t *Game) inverseDiagonal(row int,column int,value int)int{
  return t.inverseDiagonalRight(row,column,value)+t.inverseDiagonalLeft(row,column,value)+1
}


func (t *Game) inverseDiagonalRight(row int, column int, value int) int {
  if row < t.SizeTable-1 && column < t.SizeTable -1 && t.checkPosition(row+1,column+1,value){
    return t.inverseDiagonalRight(row+1,column+1,value)+1
  }

  return 0
}

func (t *Game) inverseDiagonalLeft(row int, column int, value int) int {
  if row > 0 && column > 0 && t.checkPosition(row-1,column-1,value){
    return t.inverseDiagonalLeft(row-1,column-1,value)+1
  }

  return 0
}


func (t *Game) mainDiagonal(row int, column int, value int) int {
  return t.mainDiagonalLeft(row,column,value)+t.mainDiagonalRight(row,column,value)+1
}


func (t *Game) mainDiagonalRight(row int, column int, value int) int {
  if row >0 && column < t.SizeTable-1 && t.checkPosition(row-1,column+1,value){
    return t.mainDiagonalRight(row-1,column+1,value)+1
  }
  return 0
}

func (t *Game) mainDiagonalLeft(row int, column int, value int) int {
  if row < t.SizeTable-1 && column > 0 && t.checkPosition(row+1,column-1,value){
    return t.mainDiagonalLeft(row+1,column-1,value)+1
  }
  return 0
}


func (t *Game) horizontal(row int, column int, value int) int {
  return t.horizontalRight(row,column,value)+t.horizontalLeft(row,column,value)+1
}


func (t *Game) horizontalLeft(row int, column int, value int) int {
  if column > 0 && t.checkPosition(row,column-1,value){
    return t.horizontalLeft(row,column-1,value)+1
  }
  return 0
}

func (t *Game) horizontalRight(row int, column int, value int) int {
  if column < t.SizeTable-1 && t.checkPosition(row,column+1,value){
    return t.horizontalRight(row,column+1,value)+1
  }
  return 0
}

func (t *Game ) vertical(row int, column int, value int) int {
  if row < t.SizeTable-1 && t.checkPosition(row+1,column,value){
    return t.vertical(row+1,column,value)+1
  }
  return 0
}

func (t *Game) checkPosition(row int, column int, value int) bool {
  if t.table[row][column] == value {
    return true
  }
  return false
}


func (this *Game) runTurn() int {

	cursor := Cursor{LimitPosition: this.SizeTable}
	cursor.Init()

  cursor.Print()
  this.Print()

	var input byte = 0
	for input != 65 {
		input = getInput()
		if input != 65 {
			fmt.Printf("\033[%dA", this.SizeTable+1)
      cursor.Move(input)
      this.Print()
		}
	}

	fmt.Printf("\033[%dA", this.SizeTable+2)
  this.turn = !this.turn
	return cursor.Position
}

func getInput() byte {
	t, _ := term.Open("/dev/tty")
	err := term.RawMode(t)
	if err != nil {
		log.Fatal(err)
	}
	var read int
	readBytes := make([]byte, 3)
	read, err = t.Read(readBytes)
	if err != nil {
		log.Fatal(err)
	} else if read != 3 {
		return readBytes[0]
	}

	t.Restore()
	t.Close()

	return readBytes[2]
}
