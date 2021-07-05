package codeql_go_vendor_extractor

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestExtractUniverseScope(t *testing.T) {
	assert.NoError(t, os.Setenv("CODEQL_EXTRACTOR_GO_TRAP_DIR", "/tmp/trap/go"))
	ExtractUniverseScope()
}
