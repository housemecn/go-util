package version

import (
	"bytes"
	"runtime"
	"text/template"
)

// Options for app
type Options struct {
	GitCommit string
	Version   string
	BuildTime string
	GoVersion string
	Os        string
	Arch      string
}

var versionTemplate = `Version:      {{.Version}}
Go version:   {{.GoVersion}}
Built:        {{.BuildTime}}
OS/Arch:      {{.Os}}/{{.Arch}}`

// Stringify get version string
// version "1.0.0"
// BuildTime "2020/05/27"
func Stringify(version string, buildTime string) string {
	ops := Options{
		Version:   version,
		BuildTime: buildTime,
		GoVersion: runtime.Version(),
		Os:        runtime.GOOS,
		Arch:      runtime.GOARCH,
	}
	return StringifyWithOps(ops)
}

// StringifyWithOps get version string with versionOptions
func StringifyWithOps(options Options) string {
	var doc bytes.Buffer
	t, _ := template.New("version").Parse(versionTemplate)
	_ = t.Execute(&doc, options)
	return doc.String()
}
