package controller

import (
	"fmt"
	"html/template"
	"io"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
	"www/domain"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// TestTemplate テンプレートエンジン用
type testTemplate struct {
	templates *template.Template
}

// Render htmlにレンダリング
func (t *testTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func TestMain(m *testing.M) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%sTokyo", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), "test-db", os.Getenv("DB_DATABASE"), "%2F")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}
	db.AutoMigrate(&domain.User{})
	m.Run()
	db.Migrator().DropTable(&domain.User{})
}

func initContext(form url.Values) echo.Context {
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(form.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	rec := httptest.NewRecorder()

	e := echo.New()
	e.Renderer = initTmp()

	return e.NewContext(req, rec)
}

// initTmp テンプレートエンジンの初期化
func initTmp() *testTemplate {
	d := "/var/www/view/*.html"

	t := &testTemplate{
		templates: template.Must(template.New(d).ParseGlob(d)),
	}
	return t
}
