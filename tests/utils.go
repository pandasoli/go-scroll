package main

import (
  "fmt"
  "strings"

  "github.com/pandasoli/goterm"
)


type Rect struct {
  W, H,
  X, Y int
}


func get_buff_sizes(buff []string) (w, h int) {
  h = len(buff)

  for _, line := range buff {
    if w < len(line) {
      w = len(line)
    }
  }

  return w, h
}

func clear(bounds Rect, color string) {
  for i := range make([]int, bounds.H) {
    goterm.GoToXY(bounds.X, bounds.Y + i)
    fmt.Print("\033[0m\033[" + color + "m", strings.Repeat(" ", bounds.W))
  }
}

func less(a, b int) int { if a < b { return a } else { return b } }
func big(a, b int) int { if a > b { return a } else { return b } }

// func makeBorder(bounds Rect) {
  // goterm.GoToXY(bounds.X - 1, bounds.Y - 1)
  // fmt.Print(strings.Repeat("─", bounds.W + 2))

  // goterm.GoToXY(bounds.X - 1, bounds.Y + bounds.H)
  // fmt.Print(strings.Repeat("─", bounds.W + 2))

  // for y := range make([]int, bounds.H) {
  //   y += bounds.Y

  //   goterm.GoToXY(bounds.X - 1, y)
  //   fmt.Printf("│")

  //   goterm.GoToXY(bounds.X + bounds.W, y)
  //   fmt.Printf("│")
  // }

  // goterm.GoToXY(bounds.X - 1, bounds.Y - 1)
  // fmt.Print("╭")

  // goterm.GoToXY(bounds.X + bounds.W, bounds.Y - 1)
  // fmt.Print("╮")

  // goterm.GoToXY(bounds.X - 1, bounds.Y + bounds.H)
  // fmt.Print("╰")

  // goterm.GoToXY(bounds.X + bounds.W, bounds.Y + bounds.H)
  // fmt.Print("╯")
// }
