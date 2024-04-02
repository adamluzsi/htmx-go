package htmx

import (
	"bytes"
	"go.llib.dev/frameless/pkg/pathkit"
	"go.llib.dev/frameless/pkg/reflectkit"
	"go.llib.dev/frameless/pkg/zerokit"
	"html/template"
	"net/http"
	"reflect"
	"sync"
)

func New() *HTMX { return &HTMX{} }

type HTMX struct {
	// Version of HTMX
	Version string `pattern:"^\\d+\\.\\d+\\.\\d+$"`
	// SourceURL is the url where the htmx js suppose to be available.
	// If left empty, then unpack.com will be used as the source.
	SourceURL string `is:"url"`

	// MountPoint is the Mount MountPoint where the HTMX.Handler is supposed to be mounted.
	// The template helpers use this heavily to
	//
	// default: /htmx
	MountPoint string

	once     sync.Once
	registry []regrec
}

type httpServeMux interface {
	Handle(pattern string, handler http.Handler)
}

func Register[T any](hx *HTMX, id conststring) {
	hx.register(reflectkit.TypeOf[T](), string(id))

}

func (hx *HTMX) Apply(tmpl *template.Template) *template.Template {
	tmpl = tmpl.Funcs(hx.Funcs())
	return tmpl
}

func (hx *HTMX) sourceURL() string {
	if zerokit.IsZero(hx.SourceURL) {
		return unpackURL(hx.Version)
	}
	return hx.SourceURL
}

func (hx *HTMX) helperFormFor(Entity any) *template.Template {
	tmpl := template.New("hx-form")

	return tmpl
}

const scriptTemplateText = `<script 
	type="text/javascript" 
	src="{{.Source}}"
	crossorigin="anonymous"
{{ if .Integrity }}
	integrity="{{.Integrity}}"
{{ end }}
></script>
`

var scriptTemplate = template.Must(
	template.New("script-source").
		Parse(scriptTemplateText))

var integrity = map[string] /* src */ string /* integrity hash */ {}

func (hx *HTMX) htmxScriptHTML() (template.HTML, error) {
	type Data struct {
		Source    string
		Integrity string
	}
	src := hx.sourceURL()
	var data = Data{
		Source:    src,
		Integrity: integrity[src],
	}
	var buf bytes.Buffer
	err := scriptTemplate.Execute(&buf, data)
	return template.HTML(buf.String()), err
}

func (hx *HTMX) Funcs() map[string]any {
	return map[string]any{
		"htmx": hx.htmxScriptHTML,
	}
}

type regrec struct {
	Type reflect.Type
	ID   string
}

func (hx *HTMX) register(typ reflect.Type, id string) {
	hx.init()
	hx.registry = append(hx.registry, regrec{
		Type: typ,
		ID:   id,
	})
}

func (hx *HTMX) init() {
	hx.once.Do(func() {
		hx.registry = make([]regrec, 0)
	})
}

func (hx *HTMX) EditForm(v any, opts ...string) {

}

func (hx *HTMX) getMountPoint() string {
	const defaultMountPoint = "/htmx"
	return pathkit.Clean(zerokit.Coalesce(hx.MountPoint, defaultMountPoint))
}

func (hx *HTMX) Mount(mux httpServeMux) {
	mux.Handle(hx.getMountPoint(), hx.Handler())
}

func (hx *HTMX) Handler() http.Handler {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		part, path := pathkit.Unshift(r.URL.Path)
		if pathkit.Canonical(part) == pathkit.Canonical(hx.MountPoint) {
			part, path = pathkit.Unshift(path)
		}

		for typ, reg := range hx.registry {

		}
	})
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix(hx.getMountPoint(), handler).ServeHTTP(w, r)
	})
}

type conststring string

type Attributes struct {
	// Get Issues a GET request to the given URL
	Get string `htmx:"hx-get"`
	// Post Issues a POST request to the given URL
	Post string `htmx:"hx-post"`
	// Put Issues a PUT request to the given URL
	Put string `htmx:"hx-put"`
	// Patch Issues a PATCH request to the given URL
	Patch string `htmx:"hx-patch"`
	// Delete Issues a DELETE request to the given URL
	Delete string `htmx:"hx-delete"`

	// On handle events with inline scripts on elements
	On string `htmx:"hx-on*"`
	// PushURL push a URL into the browser location bar to create history
	PushURL string `htmx:"hx-push-url"`
	// Select will select content to swap in from a response
	Select string `htmx:"hx-select"`
	// SelectOutOfBand will select content to swap in from a response, somewhere other than the target (out of band)
	SelectOutOfBand string `htmx:"hx-select-oob"`
	// Swap controls how content will swap in (outerHTML, beforeend, afterend, …)
	Swap string `htmx:"hx-swap"`
	// SwapOutOfBand mark element to swap in from a response (out of band)
	SwapOutOfBand string `htmx:"hx-swap-oob"`
	// Target specifies the target element to be swapped
	Target string `htmx:"hx-target"`
	// Trigger specifies the event that triggers the request.
	//
	// 1. Standard events: These are event names such as "click", "mouseover", etc. For example:
	// ```html
	// <div hx-get="/example" hx-trigger="click">Click me</div>
	// ```
	// 2. Standard events with filters: Events can be filtered using a boolean JavaScript expression enclosed in square brackets after the event name. For example:
	// ```html
	// <div hx-get="/example" hx-trigger="click[ctrlKey]">Control click me</div>
	// ```
	// 3. Polling definition: An element can be set to poll periodically using the `every` keyword followed by a timing declaration. For example:
	// ```html
	// <div hx-get="/latest_updates" hx-trigger="every 1s">Latest updates</div>
	// ```
	// 4. Multiple triggers: Multiple triggers can be provided, separated by commas. Each trigger gets its own options. For example:
	// ```html
	// <div hx-get="/news" hx-trigger="load, click delay:1s"></div>
	// ```
	// 5. Non-standard events: There are some additional non-standard events that htmx supports, such as "load", "revealed", and "intersect". For example:
	// ```html
	// <img src="/example.jpg" hx-trigger="revealed from:window">
	// ```
	// 6. Triggering via the `HX-Trigger` header: If you're trying to fire an event from `HX-Trigger` response header, you will likely want to use the `from:body` modifier. For example:
	// ```html
	// <div hx-get="/example" hx-trigger="my-custom-event from:body">
	//   Triggered by HX-Trigger header...
	// </div>
	// ```
	//
	// In addition to these, there are also standard event modifiers such as `once`, `changed`, `delay`, `throttle`, `from`, `target`, and `consume`.
	// These can be used to modify how the events behave. For example:
	// ```html
	// <input type="text" hx-get="/search" hx-trigger="keyup changed delay:1s target:#results">
	// ```
	//This will trigger an AJAX request on every keyup event, but only if the value of the input has changed since the last keyup event. The response will be swapped into the element with the ID "results".
	Trigger string `htmx:"hx-trigger"`
	// Vals add values to submit with the request (JSON format)
	Vals string `htmx:"hx-vals"`
}
