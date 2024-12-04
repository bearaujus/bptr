package bptr

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

// FromBoolNilAble converts a bool to a *bool. If the input value is the default value, return nil.
func FromBoolNilAble(v bool) *bool {
	if v == false {
		return nil
	}
	return &v
}

// FromByteNilAble converts a byte to a *byte. If the input value is the default value, return nil.
func FromByteNilAble(v byte) *byte {
	if v == 0 {
		return nil
	}
	return &v
}

// FromComplex64NilAble converts a complex64 to a *complex64. If the input value is the default value, return nil.
func FromComplex64NilAble(v complex64) *complex64 {
	if v == 0 { // default value for complex64
		return nil
	}
	return &v
}

// FromComplex128NilAble converts a complex128 to a *complex128. If the input value is the default value, return nil.
func FromComplex128NilAble(v complex128) *complex128 {
	if v == 0 { // default value for complex128
		return nil
	}
	return &v
}

// FromFloat32NilAble converts a float32 to a *float32. If the input value is the default value, return nil.
func FromFloat32NilAble(v float32) *float32 {
	if v == 0 { // default value for float32
		return nil
	}
	return &v
}

// FromFloat64NilAble converts a float64 to a *float64. If the input value is the default value, return nil.
func FromFloat64NilAble(v float64) *float64 {
	if v == 0 { // default value for float64
		return nil
	}
	return &v
}

// FromIntNilAble converts an int to a *int. If the input value is the default value, return nil.
func FromIntNilAble(v int) *int {
	if v == 0 { // default value for int
		return nil
	}
	return &v
}

// FromInt8NilAble converts an int8 to a *int8. If the input value is the default value, return nil.
func FromInt8NilAble(v int8) *int8 {
	if v == 0 { // default value for int8
		return nil
	}
	return &v
}

// FromInt16NilAble converts an int16 to a *int16. If the input value is the default value, return nil.
func FromInt16NilAble(v int16) *int16 {
	if v == 0 { // default value for int16
		return nil
	}
	return &v
}

// FromInt32NilAble converts an int32 to a *int32. If the input value is the default value, return nil.
func FromInt32NilAble(v int32) *int32 {
	if v == 0 { // default value for int32
		return nil
	}
	return &v
}

// FromInt64NilAble converts an int64 to a *int64. If the input value is the default value, return nil.
func FromInt64NilAble(v int64) *int64 {
	if v == 0 { // default value for int64
		return nil
	}
	return &v
}

// FromRuneNilAble converts a rune to a *rune. If the input value is the default value, return nil.
func FromRuneNilAble(v rune) *rune {
	if v == 0 { // default value for rune
		return nil
	}
	return &v
}

// FromStringNilAble converts a string to a *string. If the input value is the default value, return nil.
func FromStringNilAble(v string) *string {
	if v == "" { // default value for string
		return nil
	}
	return &v
}

// FromUintNilAble converts a uint to a *uint. If the input value is the default value, return nil.
func FromUintNilAble(v uint) *uint {
	if v == 0 { // default value for uint
		return nil
	}
	return &v
}

// FromUint8NilAble converts a uint8 to a *uint8. If the input value is the default value, return nil.
func FromUint8NilAble(v uint8) *uint8 {
	if v == 0 { // default value for uint8
		return nil
	}
	return &v
}

// FromUint16NilAble converts a uint16 to a *uint16. If the input value is the default value, return nil.
func FromUint16NilAble(v uint16) *uint16 {
	if v == 0 { // default value for uint16
		return nil
	}
	return &v
}

// FromUint32NilAble converts a uint32 to a *uint32. If the input value is the default value, return nil.
func FromUint32NilAble(v uint32) *uint32 {
	if v == 0 { // default value for uint32
		return nil
	}
	return &v
}

// FromUint64NilAble converts a uint64 to a *uint64. If the input value is the default value, return nil.
func FromUint64NilAble(v uint64) *uint64 {
	if v == 0 { // default value for uint64
		return nil
	}
	return &v
}

// ToBool converts a *bool to a bool. Returns an error if the pointer is nil.
func ToBool(v *bool) (bool, error) {
	if v == nil {
		return false, ErrNilPointer
	}
	return *v, nil
}

// ToByte converts a *byte to a byte. Returns an error if the pointer is nil.
func ToByte(v *byte) (byte, error) {
	if v == nil {
		return 0, ErrNilPointer
	}
	return *v, nil
}

// ToComplex64 converts a *complex64 to a complex64. Returns an error if the pointer is nil.
func ToComplex64(v *complex64) (complex64, error) {
	if v == nil {
		return 0, ErrNilPointer
	}
	return *v, nil
}

// ToComplex128 converts a *complex128 to a complex128. Returns an error if the pointer is nil.
func ToComplex128(v *complex128) (complex128, error) {
	if v == nil {
		return 0, ErrNilPointer
	}
	return *v, nil
}

// ToFloat32 converts a *float32 to a float32. Returns an error if the pointer is nil.
func ToFloat32(v *float32) (float32, error) {
	if v == nil {
		return 0, ErrNilPointer
	}
	return *v, nil
}

// ToFloat64 converts a *float64 to a float64. Returns an error if the pointer is nil.
func ToFloat64(v *float64) (float64, error) {
	if v == nil {
		return 0, ErrNilPointer
	}
	return *v, nil
}

// ToInt converts a *int to an int. Returns an error if the pointer is nil.
func ToInt(v *int) (int, error) {
	if v == nil {
		return 0, ErrNilPointer
	}
	return *v, nil
}

// ToInt8 converts a *int8 to an int8. Returns an error if the pointer is nil.
func ToInt8(v *int8) (int8, error) {
	if v == nil {
		return 0, ErrNilPointer
	}
	return *v, nil
}

// ToInt16 converts a *int16 to an int16. Returns an error if the pointer is nil.
func ToInt16(v *int16) (int16, error) {
	if v == nil {
		return 0, ErrNilPointer
	}
	return *v, nil
}

// ToInt32 converts a *int32 to an int32. Returns an error if the pointer is nil.
func ToInt32(v *int32) (int32, error) {
	if v == nil {
		return 0, ErrNilPointer
	}
	return *v, nil
}

// ToInt64 converts a *int64 to an int64. Returns an error if the pointer is nil.
func ToInt64(v *int64) (int64, error) {
	if v == nil {
		return 0, ErrNilPointer
	}
	return *v, nil
}

// ToRune converts a *rune to a rune. Returns an error if the pointer is nil.
func ToRune(v *rune) (rune, error) {
	if v == nil {
		return 0, ErrNilPointer
	}
	return *v, nil
}

// ToString converts a *string to a string. Returns an error if the pointer is nil.
func ToString(v *string) (string, error) {
	if v == nil {
		return "", ErrNilPointer
	}
	return *v, nil
}

// ToUint converts a *uint to a uint. Returns an error if the pointer is nil.
func ToUint(v *uint) (uint, error) {
	if v == nil {
		return 0, ErrNilPointer
	}
	return *v, nil
}

// ToUint8 converts a *uint8 to a uint8. Returns an error if the pointer is nil.
func ToUint8(v *uint8) (uint8, error) {
	if v == nil {
		return 0, ErrNilPointer
	}
	return *v, nil
}

// ToUint16 converts a *uint16 to a uint16. Returns an error if the pointer is nil.
func ToUint16(v *uint16) (uint16, error) {
	if v == nil {
		return 0, ErrNilPointer
	}
	return *v, nil
}

// ToUint32 converts a *uint32 to a uint32. Returns an error if the pointer is nil.
func ToUint32(v *uint32) (uint32, error) {
	if v == nil {
		return 0, ErrNilPointer
	}
	return *v, nil
}

// ToUint64 converts a *uint64 to a uint64. Returns an error if the pointer is nil.
func ToUint64(v *uint64) (uint64, error) {
	if v == nil {
		return 0, ErrNilPointer
	}
	return *v, nil
}

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
