package bptr

import (
	"reflect"
	"testing"
)

func TestToBoolSafe(t *testing.T) {
	type args struct {
		v *bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil",
			args: args{v: nil},
			want: false,
		},
		{
			name: "default",
			args: args{v: FromBool(false)},
			want: false,
		},
		{
			name: "with value",
			args: args{v: FromBool(true)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToBoolSafe(tt.args.v); got != tt.want {
				t.Errorf("ToBoolSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToByteSafe(t *testing.T) {
	type args struct {
		v *byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "nil",
			args: args{v: nil},
			want: 0,
		},
		{
			name: "default",
			args: args{v: FromByte(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: FromByte(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToByteSafe(tt.args.v); got != tt.want {
				t.Errorf("ToByteSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToComplex64Safe(t *testing.T) {
	type args struct {
		v *complex64
	}
	tests := []struct {
		name string
		args args
		want complex64
	}{
		{
			name: "nil",
			args: args{v: nil},
			want: 0,
		},
		{
			name: "default",
			args: args{v: FromComplex64(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: FromComplex64(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToComplex64Safe(tt.args.v); got != tt.want {
				t.Errorf("ToComplex64Safe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToComplex128Safe(t *testing.T) {
	type args struct {
		v *complex128
	}
	tests := []struct {
		name string
		args args
		want complex128
	}{
		{
			name: "nil",
			args: args{v: nil},
			want: 0,
		},
		{
			name: "default",
			args: args{v: FromComplex128(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: FromComplex128(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToComplex128Safe(tt.args.v); got != tt.want {
				t.Errorf("ToComplex128Safe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat32Safe(t *testing.T) {
	type args struct {
		v *float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "nil",
			args: args{v: nil},
			want: 0,
		},
		{
			name: "default",
			args: args{v: FromFloat32(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: FromFloat32(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFloat32Safe(tt.args.v); got != tt.want {
				t.Errorf("ToFloat32Safe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat64Safe(t *testing.T) {
	type args struct {
		v *float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "nil",
			args: args{v: nil},
			want: 0,
		},
		{
			name: "default",
			args: args{v: FromFloat64(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: FromFloat64(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFloat64Safe(tt.args.v); got != tt.want {
				t.Errorf("ToFloat64Safe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToIntSafe(t *testing.T) {
	type args struct {
		v *int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nil",
			args: args{v: nil},
			want: 0,
		},
		{
			name: "default",
			args: args{v: FromInt(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: FromInt(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToIntSafe(tt.args.v); got != tt.want {
				t.Errorf("ToIntSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt8Safe(t *testing.T) {
	type args struct {
		v *int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{
			name: "nil",
			args: args{v: nil},
			want: 0,
		},
		{
			name: "default",
			args: args{v: FromInt8(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: FromInt8(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt8Safe(tt.args.v); got != tt.want {
				t.Errorf("ToInt8Safe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt16Safe(t *testing.T) {
	type args struct {
		v *int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			name: "nil",
			args: args{v: nil},
			want: 0,
		},
		{
			name: "default",
			args: args{v: FromInt16(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: FromInt16(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt16Safe(tt.args.v); got != tt.want {
				t.Errorf("ToInt16Safe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt32Safe(t *testing.T) {
	type args struct {
		v *int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "nil",
			args: args{v: nil},
			want: 0,
		},
		{
			name: "default",
			args: args{v: FromInt32(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: FromInt32(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt32Safe(tt.args.v); got != tt.want {
				t.Errorf("ToInt32Safe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt64Safe(t *testing.T) {
	type args struct {
		v *int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "nil",
			args: args{v: nil},
			want: 0,
		},
		{
			name: "default",
			args: args{v: FromInt64(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: FromInt64(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt64Safe(tt.args.v); got != tt.want {
				t.Errorf("ToInt64Safe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToRuneSafe(t *testing.T) {
	type args struct {
		v *rune
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{
			name: "nil",
			args: args{v: nil},
			want: 0,
		},
		{
			name: "default",
			args: args{v: FromRune(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: FromRune(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToRuneSafe(tt.args.v); got != tt.want {
				t.Errorf("ToRuneSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToStringSafe(t *testing.T) {
	type args struct {
		v *string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "nil",
			args: args{v: nil},
			want: "",
		},
		{
			name: "default",
			args: args{v: FromString("")},
			want: "",
		},
		{
			name: "with value",
			args: args{v: FromString("1")},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToStringSafe(tt.args.v); got != tt.want {
				t.Errorf("ToStringSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUintSafe(t *testing.T) {
	type args struct {
		v *uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "nil",
			args: args{v: nil},
			want: 0,
		},
		{
			name: "default",
			args: args{v: FromUint(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: FromUint(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUintSafe(tt.args.v); got != tt.want {
				t.Errorf("ToUintSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint8Safe(t *testing.T) {
	type args struct {
		v *uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			name: "nil",
			args: args{v: nil},
			want: 0,
		},
		{
			name: "default",
			args: args{v: FromUint8(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: FromUint8(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint8Safe(tt.args.v); got != tt.want {
				t.Errorf("ToUint8Safe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint16Safe(t *testing.T) {
	type args struct {
		v *uint16
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "nil",
			args: args{v: nil},
			want: 0,
		},
		{
			name: "default",
			args: args{v: FromUint16(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: FromUint16(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint16Safe(tt.args.v); got != tt.want {
				t.Errorf("ToUint16Safe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint32Safe(t *testing.T) {
	type args struct {
		v *uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "nil",
			args: args{v: nil},
			want: 0,
		},
		{
			name: "default",
			args: args{v: FromUint32(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: FromUint32(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint32Safe(tt.args.v); got != tt.want {
				t.Errorf("ToUint32Safe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint64Safe(t *testing.T) {
	type args struct {
		v *uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "nil",
			args: args{v: nil},
			want: 0,
		},
		{
			name: "default",
			args: args{v: FromUint64(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: FromUint64(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUint64Safe(tt.args.v); got != tt.want {
				t.Errorf("ToUint64Safe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromBool(t *testing.T) {
	type args struct {
		v bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "default",
			args: args{v: false},
			want: false,
		},
		{
			name: "with value",
			args: args{v: true},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromBool(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("FromBool() = %v, want %v", got, &tt.want)
			}
		})
	}
}

func TestFromByte(t *testing.T) {
	type args struct {
		v byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "default",
			args: args{v: 0},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromByte(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("FromByte() = %v, want %v", got, &tt.want)
			}
		})
	}
}

func TestFromComplex64(t *testing.T) {
	type args struct {
		v complex64
	}
	tests := []struct {
		name string
		args args
		want complex64
	}{
		{
			name: "default",
			args: args{v: 0},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromComplex64(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("FromComplex64() = %v, want %v", got, &tt.want)
			}
		})
	}
}

func TestFromComplex128(t *testing.T) {
	type args struct {
		v complex128
	}
	tests := []struct {
		name string
		args args
		want complex128
	}{
		{
			name: "default",
			args: args{v: 0},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromComplex128(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("FromComplex128() = %v, want %v", got, &tt.want)
			}
		})
	}
}

func TestFromFloat32(t *testing.T) {
	type args struct {
		v float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "default",
			args: args{v: 0},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromFloat32(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("FromFloat32() = %v, want %v", got, &tt.want)
			}
		})
	}
}

func TestFromFloat64(t *testing.T) {
	type args struct {
		v float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "default",
			args: args{v: 0},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromFloat64(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("FromFloat64() = %v, want %v", got, &tt.want)
			}
		})
	}
}

func TestFromInt(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "default",
			args: args{v: 0},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromInt(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("FromInt() = %v, want %v", got, &tt.want)
			}
		})
	}
}

func TestFromInt8(t *testing.T) {
	type args struct {
		v int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{
			name: "default",
			args: args{v: 0},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromInt8(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("FromInt8() = %v, want %v", got, &tt.want)
			}
		})
	}
}

func TestFromInt16(t *testing.T) {
	type args struct {
		v int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			name: "default",
			args: args{v: 0},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromInt16(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("FromInt16() = %v, want %v", got, &tt.want)
			}
		})
	}
}

func TestFromInt32(t *testing.T) {
	type args struct {
		v int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "default",
			args: args{v: 0},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromInt32(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("FromInt32() = %v, want %v", got, &tt.want)
			}
		})
	}
}

func TestFromInt64(t *testing.T) {
	type args struct {
		v int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "default",
			args: args{v: 0},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromInt64(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("FromInt64() = %v, want %v", got, &tt.want)
			}
		})
	}
}

func TestFromRune(t *testing.T) {
	type args struct {
		v rune
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{
			name: "default",
			args: args{v: 0},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromRune(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("FromRune() = %v, want %v", got, &tt.want)
			}
		})
	}
}

func TestFromString(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{v: ""},
			want: "",
		},
		{
			name: "with value",
			args: args{v: "1"},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromString(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("FromString() = %v, want %v", got, &tt.want)
			}
		})
	}
}

func TestFromUint(t *testing.T) {
	type args struct {
		v uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "default",
			args: args{v: 0},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromUint(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("FromUint() = %v, want %v", got, &tt.want)
			}
		})
	}
}

func TestFromUint8(t *testing.T) {
	type args struct {
		v uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			name: "default",
			args: args{v: 0},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromUint8(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("FromUint8() = %v, want %v", got, &tt.want)
			}
		})
	}
}

func TestFromUint16(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "default",
			args: args{v: 0},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromUint16(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("FromUint16() = %v, want %v", got, &tt.want)
			}
		})
	}
}

func TestFromUint32(t *testing.T) {
	type args struct {
		v uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "default",
			args: args{v: 0},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromUint32(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("FromUint32() = %v, want %v", got, &tt.want)
			}
		})
	}
}

func TestFromUint64(t *testing.T) {
	type args struct {
		v uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "default",
			args: args{v: 0},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromUint64(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("FromUint64() = %v, want %v", got, &tt.want)
			}
		})
	}
}
