package helpers

import (
	"fmt"
	"html/template"
	"strconv"
)

const (
	escape = "\x1b"
)

// TemplateHelpers ...
var TemplateHelpers = template.FuncMap{
	"color": func(fg int, bg int) string {
		foreground := escape + "[38;5;" + strconv.Itoa(fg) + "m"
		background := escape + "[48;5;" + strconv.Itoa(bg) + "m"
		return fmt.Sprintf("%s%s", foreground, background)
	},
	"fg": func(c int) string {
		foreground := escape + "[38;5;" + strconv.Itoa(c) + "m"
		return foreground
	},
	"bg": func(c int) string {
		background := escape + "[48;5;" + strconv.Itoa(c) + "m"
		return background
	},
	"reset": func() string {
		reset := escape + "[0m"
		return reset
	},
}
