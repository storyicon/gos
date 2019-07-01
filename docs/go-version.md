## About `go: cannot find main module`

When using golang `under` version `1.12`, the following errors may occur when pulling out repositories that are not managed by `go module`: 
```
go: cannot find main module; see'Go help modules', 
```
so use golang `at` or `above` `1.12` to ensure the normal operation of `go module`.

Related issue: [issue#24250](https://github.com/golang/go/issues/24250)
