package utils

import (
	"runtime/debug"
	"strings"

	"github.com/cloudflare/cloudflare-go"
)

// FindGoModuleVersion digs into the build information and extracts the version
// of a module for use without the prefixed `v` (should it exist).
func FindGoModuleVersion(modulePath string) *string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		// shouldn't ever happen but just in case we aren't using modules
		return nil
	}

	for _, mod := range info.Deps {
		if mod.Path != modulePath {
			continue
		}

		version := mod.Version
		if strings.HasPrefix(version, "v") {
			version = strings.TrimPrefix(version, "v")
		}

		return cloudflare.StringPtr(version)
	}

	return nil
}
