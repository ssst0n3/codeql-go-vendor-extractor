package codeql_go_vendor_extractor

import (
	_ "github.com/github/codeql-go/extractor"
	_ "unsafe"
)

/*
borrowed from https://github.com/github/codeql-go/blob/v1.27.0/extractor/extractor.go#L213
*/

//go:linkname ExtractUniverseScope github.com/github/codeql-go/extractor.extractUniverseScope
func ExtractUniverseScope()
