package input

import (
	"bufio"
	"findNGame/color"
	"findNGame/model"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Read() model.Game{

  fmt.Printf("%s%s%s\n",color.GREEN,"Welcome to find N game, the first step is that you choose a size of the table",color.RESET)

  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Introduce the number (greater or equal to two) :")
  n,_,_:= reader.ReadLine()
  nInt,errN := strconv.Atoi(string(n))


  fmt.Printf("Introduce the size of the table (It is mandatory greater or equal to %s):",n)
  size,_,_:= reader.ReadLine()
  fmt.Println(string(nInt))


  sizeInt, errSize := strconv.Atoi(string(size))
  if errN != nil && errSize != nil {
    log.Fatal("There was a error in read input")
  }


  fmt.Printf("%s%s%v%s%v%s\n",color.RED,"The rule of the game is ",nInt," and the size of the table ",sizeInt,color.RESET)

  fmt.Println()
  return model.Game{N: nInt, SizeTable: sizeInt}
}
