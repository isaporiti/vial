# 🧪 vial
A minimalistic test assertion library for lab Gophers.

## Features

- **`Equal(t, want, got)`**
  Asserts that two comparable values are equal. Provides human-friendly diffs for strings and runes.

- **`True(t, expr)`**
  Asserts that a boolean expression is `true`.

- **`NoError(t, err)`**
  Asserts that an error is `nil`.

All assertions automatically mark the calling line as the source of the failure (via `t.Helper()`).

## Installation

```bash
go get github.com/isaporiti/vial
```

## Examples

```go
package mypkg_test

import (
	"errors"
	"testing"

	"github.com/isaporiti/vial"
)

func TestExample(t *testing.T) {
	vial.Equal(t, 42, 42)             // ✅ pass
	vial.Equal(t, "go", "gopher")     // ❌ shows string diff
	vial.True(t, 1+1 == 2)            // ✅ pass
	vial.NoError(t, nil)              // ✅ pass
	vial.NoError(t, errors.New("x"))  // ❌ fail
}
```
