package compressible

import (
	"testing"

	"github.com/GitbookIO/mimedb"
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
	s.True(Test("text/foobar"))
}

func TestCompressible(t *testing.T) {
	suite.Run(t, new(CompressibleSuite))
}
