package compressible

import (
	"mime"
	"regexp"

	"github.com/GitbookIO/mimedb"
)

func init() {
	// ensure all extensions and their associated content type of "mimedb" package
	// are stored in "mime" package, ignore the potential returned error of
	// mime.AddExtensionType.
	for ext, mimeEntry := range mimedb.DB {
		mime.AddExtensionType("."+ext, mimeEntry.ContentType)
	}
}

// Version is this package's version
const Version = "0.1.1"

var compressibleTypeRegExp = regexp.MustCompile(`^text\/|\+json$|\+text$|\+xml$`)

// Test checks whether the given contentType is compressible, using
// https://github.com/GitbookIO/mimedb as mime database. All types that not in
// mimedb but have the scheme of "text/*", "*/*+json", "*/*+text", "*/*+xml" are
// considered as compressible.
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
