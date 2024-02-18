// Package version has all its values either pulled from the runtime or set via
// flags in make at compile time.
//
// DO NOT EDIT
package version

import (
	"fmt"
	"runtime"
)

var (
	// GitCommit is the SHA that this was compiled with
	GitCommit string //nolint:gochecknoglobals

	// Version is the git tag version number that's running now
	Version string //nolint:gochecknoglobals

	// BuildDate is the date built
	BuildDate string //nolint:gochecknoglobals

	// GoVersion is the version of the go binary
	GoVersion = runtime.Version() //nolint:gochecknoglobals

	// OsArch is the os and arch targeted during build
	OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH) //nolint:gochecknoglobals
)
