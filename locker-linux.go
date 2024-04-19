//go:build !tinygo

package locker

import "embed"

//go:embed css go.mod *.go html images js template
var fs embed.FS
