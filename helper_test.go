package htmx_test

import (
	"bytes"
	"go.llib.dev/testcase/assert"
	"go.llib.dev/testcase/random"
	"html/template"
	"testing"
)

var rnd = random.New(random.CryptoSeed{})

type MyEntity struct {
	ID  string `json:"id,omitempty"`
	Foo string `json:"foo"`
	Bar int    `json:"bar"`
	Baz bool   `json:"baz"`
}

func ExecuteTemplate(tb testing.TB, tmpl *template.Template, data any) string {
	var buf bytes.Buffer
	assert.NoError(tb, tmpl.Execute(&buf, data))
	return buf.String()
}
