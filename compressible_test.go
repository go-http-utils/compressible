package compressible

import (
	"testing"

	"github.com/GitbookIO/mimedb"
	_ "github.com/GitbookIO/mimedb/autoload"
	"github.com/stretchr/testify/suite"
)

type CompressibleSuite struct {
	suite.Suite
}

func (s CompressibleSuite) TestMimeDbKeys() {
	for _, mimeEntry := range mimedb.DB {
		s.Equal(mimeEntry.Compressible, Test(mimeEntry.ContentType))
	}
}

func (s CompressibleSuite) TestInvalidType() {
	s.False(Test("foo/bar"))
}

func (s CompressibleSuite) TestMatchScheme() {
	s.True(Test("Text/foobar"))
	s.True(Test("foo/bar+jsON"))
	s.True(Test("foo/bar+texT"))
	s.True(Test("foo/bar+xml"))
}

func TestCompressible(t *testing.T) {
	suite.Run(t, new(CompressibleSuite))
}
