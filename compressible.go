package compressible

import (
	"mime"
	"regexp"

	"github.com/GitbookIO/mimedb"
)

// Version is this package's version
const Version = "0.3.0"

var compressibleTypeRegExp = regexp.MustCompile(`(?i)^text\/|\+json$|\+text$|\+xml$`)

// Test checks whether the given contentType is compressible, using
// https://github.com/GitbookIO/mimedb as mime database.Recommand to autoload
// all extensions and their associated content type by
// `import _ "github.com/GitbookIO/mimedb/autoload"` first. All types that not
// in mimedb but have the scheme of "text/*", "*/*+json", "*/*+text" and
// "*/*+xml" are considered as compressible.
func Test(contentType string) bool {
	dbMatched := false

	exts, err := mime.ExtensionsByType(contentType)

	if err != nil {
		return false
	}

	for _, ext := range exts {
		// all exts returned by mime.ExtensionsByType are always
		// start with "."
		if entry, ok := mimedb.DB[ext[1:]]; ok {
			dbMatched = true

			if entry.Compressible {
				return true
			}
		}
	}

	if !dbMatched && compressibleTypeRegExp.MatchString(contentType) {
		return true
	}

	return false
}
