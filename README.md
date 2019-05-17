graphql-ast-helpers (golang)
---

Convenience functions for the [`github.com/graphql-go/graphql/language/ast`](https://godoc.org/github.com/graphql-go/graphql/language/ast) package.

## Contents

### [**Bagger**](./bagger/bagger.go)
`Bagger` helps store a collection of known types so you can prevent duplication.

### [**Creator**](./creates/creates.go)
Lots of shorthand functions for creating correct `ast` types.

### [**Directives**](./directives/directives.go)
Helpers for accessing directive properties.

### [**Mergers**](./mergers/mergers.go)
Sane deep merging of types.

### [**Sorters**](./sorters/sorters.go)
Sorting of various `ast` collections (useful in conjuction with `Bagger`).