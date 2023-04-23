package goscroll

import (
	"fmt"

	"github.com/pandasoli/goterm"
)


func render(content_size, track_size, scroll int) (thumb_size, thumb_pos, amount int, can_scroll bool) {
  viewport_size := track_size

  thumb_size = big(1, int(float64(viewport_size) * (float64(viewport_size) / float64(content_size))))
  thumb_pos = int(float64(scroll) / float64(content_size - viewport_size) * float64(viewport_size - thumb_size))
  amount = int(float64(1) / float64(track_size) * float64(content_size))

  can_scroll = scroll + viewport_size < content_size

  return thumb_size, thumb_pos, amount, can_scroll
}


func makeOutChars(thumb_pos, thumb_size, buff_size int, ch string) (out_lines []string) {
  thumb_ch := "\033[37m" + ch
  track_ch := "\033[30m" + ch

  for range make([]int, thumb_pos) { out_lines = append(out_lines, track_ch) }
  for range make([]int, thumb_size) { out_lines = append(out_lines, thumb_ch) }
  for range make([]int, buff_size - thumb_pos - thumb_size) { out_lines = append(out_lines, track_ch) }

  return out_lines
}


func ShowXScrollbar(contents_w, buff_w, x, y, scroll_position int) (can_scroll bool, amount int) {
  if contents_w <= buff_w {
    return false, 0
  }

  thumb_size, thumb_pos, amount, can_scroll := render(contents_w, buff_w, scroll_position)

  out_chars := makeOutChars(thumb_pos, thumb_size, buff_w, "━")

  for i, ch := range out_chars {
    goterm.GoToXY(x + i, y)
    fmt.Printf("\033[0m%s\033[0m", ch)
  }

  return can_scroll, amount
}

func ShowYScrollbar(contents_h, buff_h, x, y, scroll_position int) (can_scroll bool, amount int) {
  if contents_h <= buff_h {
    return false, 0
  }

  thumb_size, thumb_pos, amount, can_scroll := render(contents_h, buff_h, scroll_position)

  out_chars := makeOutChars(thumb_pos, thumb_size, buff_h, "┃")

  for i, ch := range out_chars {
    goterm.GoToXY(x, y + i)
    fmt.Printf("\033[0m%s\033[0m", ch)
  }

  return can_scroll, amount
}
