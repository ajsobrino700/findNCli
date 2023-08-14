package model

import (
  "findNGame/color"

  "fmt"
  "strings"
)


const ENTER byte = 13
const RIGHT byte = 67
const LEFT byte = 68

var keys = map[byte] int{
    LEFT: -1,
    RIGHT: +1,
  }


type Cursor struct {
  Position int
  LimitPosition int
}

func (t *Cursor) Init(){
  t.Position = 0
}

func (t *Cursor) Move(action byte){
  value,ok := keys[action]
  if value < 0 && t.Position >0 && ok {
    t.Position--
  }else if  value > 0 && t.Position < t.LimitPosition-1 && ok{
    t.Position++
  }
  t.Print()
}

func (t *Cursor) Print(){
  spaces:= " "
  if t.Position >=1 {
    spaces = strings.Repeat(" ",(t.Position)*2+1)
  }
  fmt.Printf("%s%s%s\n",spaces,color.Yellow("\u2B07"),t.fillSpace())
}

func (t *Cursor)fillSpace()string{
  return strings.Repeat(" ",(t.LimitPosition-t.Position)*2)
}
