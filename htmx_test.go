package htmx_test

import (
	"html/template"
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

func Test_spike(t *testing.T) {
	tmpl := template.New("page")
	_ = tmpl
}
