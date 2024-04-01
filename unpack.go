package htmx

import (
	"embed"
	"encoding/json"
	"fmt"
	"go.llib.dev/frameless/pkg/zerokit"
	"io/fs"
	"path/filepath"
	"strings"
)

func init() {
	err := fs.WalkDir(unpackMetaFS, ".", func(name string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		data, err := unpackMetaFS.ReadFile(name)
		if err != nil {
			return err
		}
		var meta unpackMetaDTO
		if err := json.Unmarshal(data, &meta); err != nil {
			return err
		}

		baseName := filepath.Base(name)
		version := strings.TrimSuffix(baseName, ".json")
		integrity[unpackURL(version)] = meta.Integrity
		return nil
	})
	if err != nil {
		panic(err.Error())
	}
}

type unpackMetaDTO struct {
	Path         string `json:"path"`
	Type         string `json:"type"`
	ContentType  string `json:"contentType"`
	Integrity    string `json:"integrity"`
	LastModified string `json:"lastModified"`
	Size         int    `json:"size"`
}

//go:embed .unpack
var unpackMetaFS embed.FS

func unpackURL(version string) string {
	return fmt.Sprintf("https://unpkg.com/htmx.org@%s/dist/htmx.min.js",
		zerokit.Coalesce(version, "latest"))
}
