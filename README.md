# go-libs

Common Go libraries used by all projects.

## Troubleshouting

Error

```bash
> go get github.com/84lab/go-libs@v0.0.2
go: downloading github.com/84lab/go-libs v0.0.2
go: github.com/84lab/go-libs@v0.0.2: verifying module: github.com/84lab/go-libs@v0.0.2: reading https://sum.golang.org/lookup/github.com/84lab/go-libs@v0.0.2: 404 Not Found
```

Solution:

```bash
go env -w GOPROXY=direct
go env -w GOSUMDB=off
```
