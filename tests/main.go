package main

import (
  "fmt"

  . "github.com/pandasoli/go-scroll"
  "github.com/pandasoli/goterm"
)


func render(buff []string, w, h, yscroll, xscroll int) (can_scroll_x, can_scroll_y bool, xamount, yamount int) {
  contents_w, contents_h := get_buff_sizes(buff)

  app_x := 2
  app_y := 2

  square_x := app_x + big(0, contents_w - w)
  square_y := app_y + big(0, contents_h - h)

  clear(Rect {
    big(w, contents_w - w) * 2 + w,
    big(h, contents_h - h) * 2 + h,
    app_x, 
    app_y,
  }, "0")

  clear(Rect {
    w, h,
    square_x, square_y,
  }, "42")

  for i, line := range buff {
    if i >= yscroll && i < yscroll + h {
      if len(line) > xscroll {
        var added_colors int

        cl := "\033[42;37m"
        added_colors += len(cl)

        line = line[:xscroll] + cl + line[xscroll:]

        if len(line) - added_colors > xscroll + w {
          line = line[:xscroll + w + added_colors] + "\033[0m\033[90m" + line[xscroll + w + added_colors:]
        }
      }
    }

    goterm.GoToXY(
      app_x + big(w, contents_w - xscroll) - w,
      app_y + big(h, contents_h - yscroll) - h + i,
    )
    fmt.Print("\033[0m\033[90m" + line)
  }

  can_scroll_y, yamount = ShowYScrollbar(contents_h, h, square_x + w, square_y, yscroll)
  can_scroll_x, xamount = ShowXScrollbar(contents_w, w, square_x, square_y + h, xscroll)

  return can_scroll_x, can_scroll_y, xamount, yamount
}

func main() {
  termios, _ := goterm.SetRawMode()
  goterm.HideCursor()

  defer func() {
    goterm.RestoreMode(termios)
    goterm.ShowCursor()
  }()

  buff_w := 30
  buff_h := 10
  var yscroll int
  var xscroll int

  for {
    can_scroll_x, can_scroll_y, xamount, yamount := render(
      []string {
        "item1",
        "item2",
        "item3",
        "item4",
        "item5",
        "item6",
        "item7",
        "item8",
        "item9",
        "item10",
        "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
        "item12",
      },

      buff_w,
      buff_h,

      yscroll,
      xscroll,
    )

    str, _ := goterm.Getch()

    if str == "q" { break }

    switch str {
      case "\033[A" /* Arrow up */: if yscroll > 0 { yscroll -= yamount } 
      case "\033[B" /* Arrow down */: if can_scroll_y { yscroll += yamount } 
      case "\033[C" /* Arrow right */: if can_scroll_x { xscroll += xamount } 
      case "\033[D" /* Arrow left */: if xscroll > 0 { xscroll -= xamount } 
    }
  }
}
