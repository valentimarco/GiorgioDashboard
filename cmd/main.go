package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

// data for the template
func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Count struct {
	Counting int
}

type User struct {
	Name     string
	Password string
}

func newUser(name, password string) User {
	return User{
		Name:     name,
		Password: password,
	}
}

type UserList = []User
type Data struct {
	User UserList
}

func (d *Data) hasUser(name string) bool {
	for _, user := range d.User {
		if user.Name == name {
			return true
		}
	}
	return false
}

func newData() Data {
	return Data{
		User: []User{
			newUser("admin", "admin"),
			newUser("user", "user"),
		},
	}
}

type formData struct {
	Values map[string]string
	Errors map[string]string
}

func newFormData() formData {
	return formData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = newTemplate()
	count := Count{Counting: 1}
	data := newData()
	data_index := map[string]interface{}{
		"count": count,
		"data":  data,
		// add more key-value pairs as needed
	}
	e.GET("/", func(c echo.Context) error {

		return c.Render(200, "index", data_index)
	})

	e.POST("/count", func(c echo.Context) error {
		count.Counting++
		return c.Render(200, "count", count)
	})

	e.POST("/form", func(c echo.Context) error {
		name := c.FormValue("name")
		password := c.FormValue("password")

		if data.hasUser(name) {
			form := newFormData()
			form.Errors["name"] = "User already exists"
			form.Values["name"] = name
			return c.Render(422, "form", form)
		}

		user := newUser(name, password)
		data.User = append(data.User, user)
		c.Render(200, "form", newFormData())
		return c.Render(200, "oob-users", user)
	})

	e.Logger.Fatal(e.Start(":3000"))

}
