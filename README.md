# codeql-go-vendor-extractor

## WHY?

The [official extractor for golang](https://github.com/github/codeql-go/tree/cd1e14ed09f4b56229b5c4fb7797203193b93897/extractor/cli/go-extractor) only support gomod mode.

It's not graceful for pure vendor mode.

For example, the result of this query below will be empty:

https://lgtm.com/query/8418405387172037343/

```
import go

from CallExpr e
where e.getTarget().getName()="panic"
select e
```

But it should find the function call here:

https://github.com/ssst0n3/go-vendor-test/blob/main/vendor/st0n3/st0n3.go

```
package st0n3

func Crash() {
    panic("crash")
}
```



## How to Use?

```
GO111MODULE=off codeql database create -l go <DATABASE_NAME> -c "vendor_extractor --package <PACKAGE>"
```

And then, you will get the correct database.

For example:

// TODO

get vendor-extractor binary
```
go get -u github.com/ssst0n3/go-vendor-test/cmd/vendor_extractor
```

download source code
```
cd $GOPATH/src/
mkdir -p github.com/ssst0n3/
cd github.com/ssst0n3/
git clone https://github.com/ssst0n3/go-vendor-test.git
```

create database
```
codeql database create -l go /tmp/go-vendor-test -c "vendor_extractor --package ."
```
