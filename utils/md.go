package utils

import (
	"fmt"
	"html/template"

	"github.com/russross/blackfriday"
)

func ParseMdToHtml(args ...interface{}) template.HTML {
	input := fmt.Sprintf("%s", args...)
	html := blackfriday.MarkdownCommon([]byte(input))
	return template.HTML(html)
}
