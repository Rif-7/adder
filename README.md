# adder

The `adder` module provides a simple `Add` function that demonstrates both basic and generic arithmetic using Go.

This repository is structured to follow proper semantic import versioning, with separate import paths for major versions:

- `v1`: Basic integer-only addition  
- `v2`: Generic addition using type parameters with support for `int` and `float` types

## Versions

- `v1.0.0`: Initial release with `Add(int, int) int`
- `v1.0.1`: Added GoDoc comments
- `v2.0.0`: Breaking change – converted `Add` to a generic function using Go 1.18+ type parameters and `golang.org/x/exp/constraints`

## Install

To use version 1:

```bash
go get github.com/Rif-7/adder@v1.0.1
```

```go
import "github.com/Rif-7/adder"
```

To use version 2:

```bash
go get github.com/Rif-7/adder/v2@v2.0.0
```

```go
import "github.com/Rif-7/adder/v2"
```

## About

This project was created as a learning exercise to understand how Go modules and semantic versioning work, including how to manage backward-incompatible changes using versioned module paths.


---

## License

This project is licensed under the MIT License – see the [LICENSE](LICENSE) file for details.
