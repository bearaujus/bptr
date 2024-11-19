# BPtr - Safe Pointer Conversions in Go

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/bearaujus/bptr/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/bearaujus/bptr)](https://goreportcard.com/report/github.com/bearaujus/bptr)

BPtr provides a set of utilities to safely handle pointer conversions in Go, preventing nil pointer dereferences and simplifying code where pointers to primitives are used.

## Installation

Install BPtr using:

```shell
go get github.com/bearaujus/bptr
```

Import BPtr using:
```shell
import "github.com/bearaujus/bptr"
```

## Functions

### 1. `To*Safe` Functions

The To*Safe functions safely convert pointers to primitive types. If the input pointer is nil, a default value for the type is returned:

Default values:
- bool → false
- Numeric types (int, float, etc.) → 0
- string → ""

```go
func ToBoolSafe(v *bool) bool
func ToByteSafe(v *byte) byte
func ToComplex64Safe(v *complex64) complex64
func ToComplex128Safe(v *complex128) complex128
func ToFloat32Safe(v *float32) float32
func ToFloat64Safe(v *float64) float64
func ToIntSafe(v *int) int
func ToInt8Safe(v *int8) int8
func ToInt16Safe(v *int16) int16
func ToInt32Safe(v *int32) int32
func ToInt64Safe(v *int64) int64
func ToRuneSafe(v *rune) rune
func ToStringSafe(v *string) string
func ToUintSafe(v *uint) uint
func ToUint8Safe(v *uint8) uint8
func ToUint16Safe(v *uint16) uint16
func ToUint32Safe(v *uint32) uint32
func ToUint64Safe(v *uint64) uint64
```

### 2. `To*` Functions

The To* functions convert a pointer to its primitive type but return an error (ErrNilPointer) if the input pointer is nil.

```go
func ToBool(v *bool) (bool, error)
func ToByte(v *byte) (byte, error)
func ToComplex64(v *complex64) (complex64, error)
func ToComplex128(v *complex128) (complex128, error)
func ToFloat32(v *float32) (float32, error)
func ToFloat64(v *float64) (float64, error)
func ToInt(v *int) (int, error)
func ToInt8(v *int8) (int8, error)
func ToInt16(v *int16) (int16, error)
func ToInt32(v *int32) (int32, error)
func ToInt64(v *int64) (int64, error)
func ToRune(v *rune) (rune, error)
func ToString(v *string) (string, error)
func ToUint(v *uint) (uint, error)
func ToUint8(v *uint8) (uint8, error)
func ToUint16(v *uint16) (uint16, error)
func ToUint32(v *uint32) (uint32, error)
func ToUint64(v *uint64) (uint64, error)
```

### 3. `From*` Functions

The From* functions are utility functions that take a primitive type and return a pointer to that type.

```go
func FromBool(v bool) *bool
func FromByte(v byte) *byte
func FromComplex64(v complex64) *complex64
func FromComplex128(v complex128) *complex128
func FromFloat32(v float32) *float32
func FromFloat64(v float64) *float64
func FromInt(v int) *int
func FromInt8(v int8) *int8
func FromInt16(v int16) *int16
func FromInt32(v int32) *int32
func FromInt64(v int64) *int64
func FromRune(v rune) *rune
func FromString(v string) *string
func FromUint(v uint) *uint
func FromUint8(v uint8) *uint8
func FromUint16(v uint16) *uint16
func FromUint32(v uint32) *uint32
func FromUint64(v uint64) *uint64
```

## Usage Example

```go
var bptrBool *bool = nil
boolValue := bptr.ToBoolSafe(bptrBool) // boolValue will be false

var bptrByte *byte = nil
byteValue := bptr.ToByteSafe(bptrByte) // byteValue will be 0

var bptrComplex64 *complex64 = nil
complex64Value := bptr.ToComplex64Safe(bptrComplex64) // complex64Value will be 0

// ... and so on for other types
```

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/bearaujus/bptr/blob/master/LICENSE) file for details.
