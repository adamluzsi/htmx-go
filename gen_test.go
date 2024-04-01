//go:build gogenerate

package htmx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go.llib.dev/testcase/assert"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

var hxVersions = []string{
	"1.3.0",
	"1.3.1",
	"v0.0.1",
	"v0.0.3",
	"v0.0.4",
	"v0.0.5",
	"v0.0.6",
	"v0.0.7",
	"v0.0.8",
	"v0.1.0",
	"v0.1.1",
	"v0.1.2",
	"v0.2.0",
	"v0.3.0",
	"v0.4.0",
	"v0.4.1",
	"v1.0.0",
	"v1.0.1",
	"v1.0.2",
	"v1.1.0",
	"v1.2.0",
	"v1.2.1",
	"v1.3.2",
	"v1.3.3",
	"v1.4.0",
	"v1.4.1",
	"v1.5.0",
	"v1.6.0",
	"v1.6.1",
	"v1.7.0",
	"v1.8.0",
	"v1.8.1",
	"v1.8.2",
	"v1.8.3",
	"v1.8.4",
	"v1.8.5",
	"v1.8.6",
	"v1.9.0",
	"v1.9.1",
	"v1.9.10",
	"v1.9.11",
	"v1.9.2",
	"v1.9.3",
	"v1.9.4",
	"v1.9.5",
	"v1.9.6",
	"v1.9.7",
	"v1.9.8",
	"v1.9.9",
	"v2.0.0-alpha2",
	"v2.0.0-beta1",
	"v2.0.0-beta2",
}

func Test_unpack(t *testing.T) {
	const unpackMetaDirName = ".unpack"
	assert.NoError(t, os.MkdirAll(unpackMetaDirName, 0700))
	for _, version := range hxVersions {
		metaFilePath := filepath.Join(unpackMetaDirName, version+".json")
		if _, err := os.Stat(metaFilePath); err == nil {
			continue
		}
		t.Logf("fetching %s version meta from unpack", version)
		resp, err := http.Get(fmt.Sprintf(unpackURL(version)+"?meta", version))
		assert.NoError(t, err)
		if resp.StatusCode == http.StatusNotFound {
			t.Logf("unpack doesn't recognise %s version", version)
			continue
		}
		bs, err := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		assert.NoError(t, err)
		t.Logf("response: %s", string(bs))
		var buf bytes.Buffer
		assert.NoError(t, json.Indent(&buf, bs, "", "  "))
		assert.NoError(t, os.WriteFile(metaFilePath, buf.Bytes(), 0644))
	}
}
