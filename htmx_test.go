package htmx_test

import (
	"go.llib.dev/htmx"
	"html/template"
	"net/http"
	"testing"
)

const exampleTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>My Website</title>
    <link rel="stylesheet" href="https://unpkg.com/purecss@2.0.6/build/pure-min.css" integrity="sha384-Uu6IeWbM+gzNVXJcM9XV3SohHtmWE+3VGi496jvgX1jyvDTXfdK+rfZc8C1Aehk5" crossorigin="anonymous">
</head>
<body>
    <div class="pure-g">
        <div class="pure-u-1-2">
            <h1>My Website</h1>
            <ul class="pure-menu pure-menu-open pure-menu-horizontal">
                <li class="pure-menu-item"><a href="#" class="pure-menu-link">Home</a></li>
                <li class="pure-menu-item"><a href="#" class="pure-menu-link">About</a></li>
                <li class="pure-menu-item"><a href="#" class="pure-menu-link">Contact</a></li>
            </ul>
        </div>
        <div class="pure-u-1-2">
            <h2>List of Records</h2>
            <ul>
                <li>Record 1</li>
                <li>Record 2</li>
                <li>Record 3</li>
                <li>Record 4</li>
                <li>Record 5</li>
            </ul>
        </div>
    </div>
</body>
</html>

`

const editformTemplate = `
{{ editform ".Ent" }}"
`

const editformTemplateExpected = `
<form hx-post="/htmx/my-ent/{{ .ID }}" hx-swap="outerHTML">
  <input type="hidden" name="id" value="{{ .ID }}">
  <label for="foo">Foo:</label>
  <input type="text" name="foo" value="{{ .Foo }}" required>
  <br>
  <label for="bar">Bar:</label>
  <input type="number" name="bar" value="{{ .Bar }}" required>
  <br>
  <button type="submit">Save Changes</button>
</form>
`

func TestHTMX_EditForm(t *testing.T) {
	var (
		mux  = http.NewServeMux()
		tmpl = template.New("page")
	)

	hx := &htmx.HTMX{}
	htmx.Register[MyEntity](hx, "my-ent")
	tmpl = hx.Apply(tmpl)
	hx.Mount(mux)

	var ent = MyEntity{
		ID:  "42",
		Foo: "foo",
		Bar: 42,
		Baz: true,
	}

	type Data struct{ Ent MyEntity }
	ExecuteTemplate(t, tmpl, Data{Ent: ent})
}
