package bptr

// ToBoolSafe safely convert *bool to bool.
//
// If the input pointer is nil, return the default value.
func ToBoolSafe(v *bool) bool {
	// If v is nil ptr bool, return default value
	if v == nil {
		return false
	}
	// If v is not nil ptr bool, dereference it
	return *v
}

// ToByteSafe safely convert *byte to byte.
//
// If the input pointer is nil, return the default value.
func ToByteSafe(v *byte) byte {
	// If v is nil ptr byte,return default value
	if v == nil {
		return 0
	}
	// If v is not nil ptr byte, dereference it
	return *v
}

// ToComplex64Safe safely convert *complex64 to complex64.
//
// If the input pointer is nil, return the default value.
func ToComplex64Safe(v *complex64) complex64 {
	// If v is nil ptr complex64, return default value
	if v == nil {
		return 0
	}
	// If v is not nil ptr complex64, dereference it
	return *v
}

// ToComplex128Safe safely convert *complex128 to complex128.
//
// If the input pointer is nil, return the default value.
func ToComplex128Safe(v *complex128) complex128 {
	// If v is nil ptr complex128, return default value
	if v == nil {
		return 0
	}
	// If v is not nil ptr complex128, dereference it
	return *v
}

// ToFloat32Safe safely convert *float32 to float32.
//
// If the input pointer is nil, return the default value.
func ToFloat32Safe(v *float32) float32 {
	// If v is nil ptr float32, return default value
	if v == nil {
		return 0
	}
	// If v is not nil ptr float32, dereference it
	return *v
}

// ToFloat64Safe safely convert *float64 to float64.
//
// If the input pointer is nil, return the default value.
func ToFloat64Safe(v *float64) float64 {
	// If v is nil ptr float64, return default value
	if v == nil {
		return 0
	}
	// If v is not nil ptr float64, dereference it
	return *v
}

// ToIntSafe safely convert *int to int.
//
// If the input pointer is nil, return the default value.
func ToIntSafe(v *int) int {
	// If v is nil ptr int, return default value
	if v == nil {
		return 0
	}
	// If v is not nil ptr int, dereference it
	return *v
}

// ToInt8Safe safely convert *int8 to int8.
//
// If the input pointer is nil, return the default value.
func ToInt8Safe(v *int8) int8 {
	// If v is nil ptr int8, return default value
	if v == nil {
		return 0
	}
	// If v is not nil ptr int8, dereference it
	return *v
}

// ToInt16Safe safely convert *int16 to int16.
//
// If the input pointer is nil, return the default value.
func ToInt16Safe(v *int16) int16 {
	// If v is nil ptr int16, return default value
	if v == nil {
		return 0
	}
	// If v is not nil ptr int16, dereference it
	return *v
}

// ToInt32Safe safely convert *int32 to int32.
//
// If the input pointer is nil, return the default value.
func ToInt32Safe(v *int32) int32 {
	// If v is nil ptr int32, return default value
	if v == nil {
		return 0
	}
	// If v is not nil ptr int32, dereference it
	return *v
}

// ToInt64Safe safely convert *int64 to int64.
//
// If the input pointer is nil, return the default value.
func ToInt64Safe(v *int64) int64 {
	// If v is nil ptr int64, return default value
	if v == nil {
		return 0
	}
	// If v is not nil ptr int64, dereference it
	return *v
}

// ToRuneSafe safely convert *rune to rune.
//
// If the input pointer is nil, return the default value.
func ToRuneSafe(v *rune) rune {
	// If v is nil ptr rune, return default value
	if v == nil {
		return 0
	}
	// If v is not nil ptr rune, dereference it
	return *v
}

// ToStringSafe safely convert *string to string.
//
// If the input pointer is nil, return the default value.
func ToStringSafe(v *string) string {
	// If v is nil ptr string, return default value
	if v == nil {
		return ""
	}
	// If v is not nil ptr string, dereference it
	return *v
}

// ToUintSafe safely convert *uint to uint.
//
// If the input pointer is nil, return the default value.
func ToUintSafe(v *uint) uint {
	// If v is nil ptr uint, return default value
	if v == nil {
		return 0
	}
	// If v is not nil ptr uint, dereference it
	return *v
}

// ToUint8Safe safely convert *uint8 to uint8.
//
// If the input pointer is nil, return the default value.
func ToUint8Safe(v *uint8) uint8 {
	// If v is nil ptr uint8, return default value
	if v == nil {
		return 0
	}
	// If v is not nil ptr uint8, dereference it
	return *v
}

// ToUint16Safe safely convert *uint16 to uint16.
//
// If the input pointer is nil, return the default value.
func ToUint16Safe(v *uint16) uint16 {
	// If v is nil ptr uint16, return default value
	if v == nil {
		return 0
	}
	// If v is not nil ptr uint16, dereference it
	return *v
}

// ToUint32Safe safely convert *uint32 to uint32.
//
// If the input pointer is nil, return the default value.
func ToUint32Safe(v *uint32) uint32 {
	// If v is nil ptr uint32, return default value
	if v == nil {
		return 0
	}
	// If v is not nil ptr uint32, dereference it
	return *v
}

// ToUint64Safe safely convert *uint64 to uint64.
//
// If the input pointer is nil, return the default value.
func ToUint64Safe(v *uint64) uint64 {
	// If v is nil ptr uint64, return default value
	if v == nil {
		return 0
	}
	// If v is not nil ptr uint64, dereference it
	return *v
}

// FromBool convert *bool to bool.
func FromBool(v bool) *bool {
	// Reference the bool object to ptr bool
	return &v
}

// FromByte convert *byte to byte.
func FromByte(v byte) *byte {
	// Reference the byte object to ptr byte
	return &v
}

// FromComplex64 convert *complex64 to complex64.
func FromComplex64(v complex64) *complex64 {
	// Reference the complex64 object to ptr complex64
	return &v
}

// FromComplex128 convert *complex128 to complex128.
func FromComplex128(v complex128) *complex128 {
	// Reference the complex128 object to ptr complex128
	return &v
}

// FromFloat32 convert *float32 to float32.
func FromFloat32(v float32) *float32 {
	// Reference the float32 object to ptr float32
	return &v
}

// FromFloat64 convert *float64 to float64.
func FromFloat64(v float64) *float64 {
	// Reference the float64 object to ptr float64
	return &v
}

// FromInt convert *int to int.
func FromInt(v int) *int {
	// Reference the int object to ptr int
	return &v
}

// FromInt8 convert *int8 to int8.
func FromInt8(v int8) *int8 {
	// Reference the int8 object to ptr int8
	return &v
}

// FromInt16 convert *int16 to int16.
func FromInt16(v int16) *int16 {
	// Reference the int16 object to ptr int16
	return &v
}

// FromInt32 convert *int32 to int32.
func FromInt32(v int32) *int32 {
	// Reference the int32 object to ptr int32
	return &v
}

// FromInt64 convert *int64 to int64.
func FromInt64(v int64) *int64 {
	// Reference the int64 object to ptr int64
	return &v
}

// FromRune convert *rune to rune.
func FromRune(v rune) *rune {
	// Reference the rune object to ptr rune
	return &v
}

// FromString convert *string to string.
func FromString(v string) *string {
	// Reference the string object to ptr string
	return &v
}

// FromUint convert *uint to uint.
func FromUint(v uint) *uint {
	// Reference the uint object to ptr uint
	return &v
}

// FromUint8 convert *uint8 to uint8.
func FromUint8(v uint8) *uint8 {
	// Reference the uint8 object to ptr uint8
	return &v
}

// FromUint16 convert *uint16 to uint16.
func FromUint16(v uint16) *uint16 {
	// Reference the uint16 object to ptr uint16
	return &v
}

// FromUint32 convert *uint32 to uint32.
func FromUint32(v uint32) *uint32 {
	// Reference the uint32 object to ptr uint32
	return &v
}

// FromUint64 convert *uint64 to uint64.
func FromUint64(v uint64) *uint64 {
	// Reference the uint64 object to ptr uint64
	return &v
}
