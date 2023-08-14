package color

import "fmt"

const BLUE = "\033[34m"
const GREEN = "\033[32m"
const RED = "\033[31m"
const RESET = "\033[0m"
const YELLOW = "\033[33m"


func Red(value string) string{
  return fmt.Sprintf("%s%s%s",RED,value, RESET)
}

func Blue(value string) string{
  return fmt.Sprintf("%s%s%s",BLUE,value, RESET)
}

func Green(value string) string{
  return fmt.Sprintf("%s%s%s",GREEN,value, RESET)
}

func Yellow(value string) string {
  return fmt.Sprintf("%s%s%s",YELLOW,value, RESET)
}
