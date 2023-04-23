package goscroll

import (
  "fmt"

  "github.com/pandasoli/goterm"
)


var debug_line = 0
func debug(a ...any) {
  goterm.GoToXY(40, 2 + debug_line)
  debug_line++
  fmt.Print(a...)
}
