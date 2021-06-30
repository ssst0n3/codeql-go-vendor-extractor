# codeql-go-vendor-extractor

## WHY?

The [official extractor for golang](https://github.com/github/codeql-go/tree/cd1e14ed09f4b56229b5c4fb7797203193b93897/extractor/cli/go-extractor) only support gomod mode.

It's not the graceful for pure vendor mode.

For example, this query below will failed on lgtm.com:

