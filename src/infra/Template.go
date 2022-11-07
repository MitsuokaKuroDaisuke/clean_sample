package infra

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// Template テンプレートエンジン用
type Template struct {
	templates *template.Template
}

// Render htmlにレンダリング
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// InitTmp テンプレートエンジンの初期化
func InitTmp() *Template {
	d := "/var/www/src/view/*.html"

	t := &Template{
		templates: template.Must(template.New(d).ParseGlob(d)),
	}
	return t
}
