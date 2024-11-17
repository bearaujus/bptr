package bptr

// ToBoolSafe converts a *bool to a bool safely.
// If the input pointer is nil, it returns false.
func ToBoolSafe(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

// ToByteSafe converts a *byte to a byte safely.
// If the input pointer is nil, it returns 0.
func ToByteSafe(v *byte) byte {
	if v == nil {
		return 0
	}
	return *v
}

// ToComplex64Safe converts a *complex64 to a complex64 safely.
// If the input pointer is nil, it returns 0.
func ToComplex64Safe(v *complex64) complex64 {
	if v == nil {
		return 0
	}
	return *v
}

// ToComplex128Safe converts a *complex128 to a complex128 safely.
// If the input pointer is nil, it returns 0.
func ToComplex128Safe(v *complex128) complex128 {
	if v == nil {
		return 0
	}
	return *v
}

// ToFloat32Safe converts a *float32 to a float32 safely.
// If the input pointer is nil, it returns 0.
func ToFloat32Safe(v *float32) float32 {
	if v == nil {
		return 0
	}
	return *v
}

// ToFloat64Safe converts a *float64 to a float64 safely.
// If the input pointer is nil, it returns 0.
func ToFloat64Safe(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}

// ToIntSafe converts a *int to an int safely.
// If the input pointer is nil, it returns 0.
func ToIntSafe(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

// ToInt8Safe converts a *int8 to an int8 safely.
// If the input pointer is nil, it returns 0.
func ToInt8Safe(v *int8) int8 {
	if v == nil {
		return 0
	}
	return *v
}

// ToInt16Safe converts a *int16 to an int16 safely.
// If the input pointer is nil, it returns 0.
func ToInt16Safe(v *int16) int16 {
	if v == nil {
		return 0
	}
	return *v
}

// ToInt32Safe converts a *int32 to an int32 safely.
// If the input pointer is nil, it returns 0.
func ToInt32Safe(v *int32) int32 {
	if v == nil {
		return 0
	}
	return *v
}

// ToInt64Safe converts a *int64 to an int64 safely.
// If the input pointer is nil, it returns 0.
func ToInt64Safe(v *int64) int64 {
	if v == nil {
		return 0
	}
	return *v
}

// ToRuneSafe converts a *rune to a rune safely.
// If the input pointer is nil, it returns 0.
func ToRuneSafe(v *rune) rune {
	if v == nil {
		return 0
	}
	return *v
}

// ToStringSafe converts a *string to a string safely.
// If the input pointer is nil, it returns an empty string.
func ToStringSafe(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

// ToUintSafe converts a *uint to a uint safely.
// If the input pointer is nil, it returns 0.
func ToUintSafe(v *uint) uint {
	if v == nil {
		return 0
	}
	return *v
}

// ToUint8Safe converts a *uint8 to a uint8 safely.
// If the input pointer is nil, it returns 0.
func ToUint8Safe(v *uint8) uint8 {
	if v == nil {
		return 0
	}
	return *v
}

// ToUint16Safe converts a *uint16 to a uint16 safely.
// If the input pointer is nil, it returns 0.
func ToUint16Safe(v *uint16) uint16 {
	if v == nil {
		return 0
	}
	return *v
}

// ToUint32Safe converts a *uint32 to a uint32 safely.
// If the input pointer is nil, it returns 0.
func ToUint32Safe(v *uint32) uint32 {
	if v == nil {
		return 0
	}
	return *v
}

// ToUint64Safe converts a *uint64 to a uint64 safely.
// If the input pointer is nil, it returns 0.
func ToUint64Safe(v *uint64) uint64 {
	if v == nil {
		return 0
	}
	return *v
}

// FromBool converts a bool to a *bool.
func FromBool(v bool) *bool {
	return &v
}

// FromByte converts a byte to a *byte.
func FromByte(v byte) *byte {
	return &v
}

// FromComplex64 converts a complex64 to a *complex64.
func FromComplex64(v complex64) *complex64 {
	return &v
}

// FromComplex128 converts a complex128 to a *complex128.
func FromComplex128(v complex128) *complex128 {
	return &v
}

// FromFloat32 converts a float32 to a *float32.
func FromFloat32(v float32) *float32 {
	return &v
}

// FromFloat64 converts a float64 to a *float64.
func FromFloat64(v float64) *float64 {
	return &v
}

// FromInt converts an int to a *int.
func FromInt(v int) *int {
	return &v
}

// FromInt8 converts an int8 to a *int8.
func FromInt8(v int8) *int8 {
	return &v
}

// FromInt16 converts an int16 to a *int16.
func FromInt16(v int16) *int16 {
	return &v
}

// FromInt32 converts an int32 to a *int32.
func FromInt32(v int32) *int32 {
	return &v
}

// FromInt64 converts an int64 to a *int64.
func FromInt64(v int64) *int64 {
	return &v
}

// FromRune converts a rune to a *rune.
func FromRune(v rune) *rune {
	return &v
}

// FromString converts a string to a *string.
func FromString(v string) *string {
	return &v
}

// FromUint converts a uint to a *uint.
func FromUint(v uint) *uint {
	return &v
}

// FromUint8 converts a uint8 to a *uint8.
func FromUint8(v uint8) *uint8 {
	return &v
}

// FromUint16 converts a uint16 to a *uint16.
func FromUint16(v uint16) *uint16 {
	return &v
}

// FromUint32 converts a uint32 to a *uint32.
func FromUint32(v uint32) *uint32 {
	return &v
}

// FromUint64 converts a uint64 to a *uint64.
func FromUint64(v uint64) *uint64 {
	return &v
}
