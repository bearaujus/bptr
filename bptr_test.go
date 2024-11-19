package bptr_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/bearaujus/bptr"
)

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
			if got := bptr.FromBool(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
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
			if got := bptr.FromByte(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
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
			if got := bptr.FromComplex64(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
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
			if got := bptr.FromComplex128(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
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
			if got := bptr.FromFloat32(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
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
			if got := bptr.FromFloat64(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
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
			if got := bptr.FromInt(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
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
			if got := bptr.FromInt8(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
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
			if got := bptr.FromInt16(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
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
			if got := bptr.FromInt32(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
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
			if got := bptr.FromInt64(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
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
			if got := bptr.FromRune(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
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
			if got := bptr.FromString(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
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
			if got := bptr.FromUint(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
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
			if got := bptr.FromUint8(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
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
			if got := bptr.FromUint16(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
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
			if got := bptr.FromUint32(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
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
			if got := bptr.FromUint64(tt.args.v); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("FromUint64() = %v, want %v", got, &tt.want)
			}
		})
	}
}

func TestToBool(t *testing.T) {
	type args struct {
		v *bool
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr error
	}{
		{
			name:    "nil",
			args:    args{v: nil},
			want:    false,
			wantErr: bptr.ErrNilPointer,
		},
		{
			name:    "default",
			args:    args{v: bptr.FromBool(false)},
			want:    false,
			wantErr: nil,
		},
		{
			name:    "with value",
			args:    args{v: bptr.FromBool(true)},
			want:    true,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bptr.ToBool(tt.args.v)
			if got != tt.want || !errors.Is(err, tt.wantErr) {
				t.Errorf("ToBool() = %v, %v; want %v, %v", got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestToByte(t *testing.T) {
	type args struct {
		v *byte
	}
	tests := []struct {
		name    string
		args    args
		want    byte
		wantErr error
	}{
		{
			name:    "nil",
			args:    args{v: nil},
			want:    0,
			wantErr: bptr.ErrNilPointer,
		},
		{
			name:    "default",
			args:    args{v: bptr.FromByte(0)},
			want:    0,
			wantErr: nil,
		},
		{
			name:    "with value",
			args:    args{v: bptr.FromByte(1)},
			want:    1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bptr.ToByte(tt.args.v)
			if got != tt.want || !errors.Is(err, tt.wantErr) {
				t.Errorf("ToByte() = %v, %v; want %v, %v", got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestToComplex64(t *testing.T) {
	type args struct {
		v *complex64
	}
	tests := []struct {
		name    string
		args    args
		want    complex64
		wantErr error
	}{
		{
			name:    "nil",
			args:    args{v: nil},
			want:    0,
			wantErr: bptr.ErrNilPointer,
		},
		{
			name:    "default",
			args:    args{v: bptr.FromComplex64(0)},
			want:    0,
			wantErr: nil,
		},
		{
			name:    "with value",
			args:    args{v: bptr.FromComplex64(1)},
			want:    1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bptr.ToComplex64(tt.args.v)
			if got != tt.want || !errors.Is(err, tt.wantErr) {
				t.Errorf("ToComplex64() = %v, %v; want %v, %v", got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestToComplex128(t *testing.T) {
	type args struct {
		v *complex128
	}
	tests := []struct {
		name    string
		args    args
		want    complex128
		wantErr error
	}{
		{
			name:    "nil",
			args:    args{v: nil},
			want:    0,
			wantErr: bptr.ErrNilPointer,
		},
		{
			name:    "default",
			args:    args{v: bptr.FromComplex128(0)},
			want:    0,
			wantErr: nil,
		},
		{
			name:    "with value",
			args:    args{v: bptr.FromComplex128(1)},
			want:    1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bptr.ToComplex128(tt.args.v)
			if got != tt.want || !errors.Is(err, tt.wantErr) {
				t.Errorf("ToComplex128() = %v, %v; want %v, %v", got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestToFloat32(t *testing.T) {
	type args struct {
		v *float32
	}
	tests := []struct {
		name    string
		args    args
		want    float32
		wantErr error
	}{
		{
			name:    "nil",
			args:    args{v: nil},
			want:    0,
			wantErr: bptr.ErrNilPointer,
		},
		{
			name:    "default",
			args:    args{v: bptr.FromFloat32(0)},
			want:    0,
			wantErr: nil,
		},
		{
			name:    "with value",
			args:    args{v: bptr.FromFloat32(1)},
			want:    1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bptr.ToFloat32(tt.args.v)
			if got != tt.want || !errors.Is(err, tt.wantErr) {
				t.Errorf("ToFloat32() = %v, %v; want %v, %v", got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestToFloat64(t *testing.T) {
	type args struct {
		v *float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr error
	}{
		{
			name:    "nil",
			args:    args{v: nil},
			want:    0,
			wantErr: bptr.ErrNilPointer,
		},
		{
			name:    "default",
			args:    args{v: bptr.FromFloat64(0)},
			want:    0,
			wantErr: nil,
		},
		{
			name:    "with value",
			args:    args{v: bptr.FromFloat64(1)},
			want:    1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bptr.ToFloat64(tt.args.v)
			if got != tt.want || !errors.Is(err, tt.wantErr) {
				t.Errorf("ToFloat64() = %v, %v; want %v, %v", got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestToInt(t *testing.T) {
	type args struct {
		v *int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{
			name:    "nil",
			args:    args{v: nil},
			want:    0,
			wantErr: bptr.ErrNilPointer,
		},
		{
			name:    "default",
			args:    args{v: bptr.FromInt(0)},
			want:    0,
			wantErr: nil,
		},
		{
			name:    "with value",
			args:    args{v: bptr.FromInt(1)},
			want:    1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bptr.ToInt(tt.args.v)
			if got != tt.want || !errors.Is(err, tt.wantErr) {
				t.Errorf("ToInt() = %v, %v; want %v, %v", got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestToInt8(t *testing.T) {
	type args struct {
		v *int8
	}
	tests := []struct {
		name    string
		args    args
		want    int8
		wantErr error
	}{
		{
			name:    "nil",
			args:    args{v: nil},
			want:    0,
			wantErr: bptr.ErrNilPointer,
		},
		{
			name:    "default",
			args:    args{v: bptr.FromInt8(0)},
			want:    0,
			wantErr: nil,
		},
		{
			name:    "with value",
			args:    args{v: bptr.FromInt8(-128)},
			want:    -128,
			wantErr: nil,
		},
		{
			name:    "positive value",
			args:    args{v: bptr.FromInt8(127)},
			want:    127,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bptr.ToInt8(tt.args.v)
			if got != tt.want || !errors.Is(err, tt.wantErr) {
				t.Errorf("ToInt8() = %v, %v; want %v, %v", got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestToInt16(t *testing.T) {
	type args struct {
		v *int16
	}
	tests := []struct {
		name    string
		args    args
		want    int16
		wantErr error
	}{
		{
			name:    "nil",
			args:    args{v: nil},
			want:    0,
			wantErr: bptr.ErrNilPointer,
		},
		{
			name:    "default",
			args:    args{v: bptr.FromInt16(0)},
			want:    0,
			wantErr: nil,
		},
		{
			name:    "with value",
			args:    args{v: bptr.FromInt16(-32768)},
			want:    -32768,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bptr.ToInt16(tt.args.v)
			if got != tt.want || !errors.Is(err, tt.wantErr) {
				t.Errorf("ToInt16() = %v, %v; want %v, %v", got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestToInt32(t *testing.T) {
	type args struct {
		v *int32
	}
	tests := []struct {
		name    string
		args    args
		want    int32
		wantErr error
	}{
		{
			name:    "nil",
			args:    args{v: nil},
			want:    0,
			wantErr: bptr.ErrNilPointer,
		},
		{
			name:    "default",
			args:    args{v: bptr.FromInt32(0)},
			want:    0,
			wantErr: nil,
		},
		{
			name:    "with value",
			args:    args{v: bptr.FromInt32(1)},
			want:    1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bptr.ToInt32(tt.args.v)
			if got != tt.want || !errors.Is(err, tt.wantErr) {
				t.Errorf("ToInt32() = %v, %v; want %v, %v", got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestToInt64(t *testing.T) {
	type args struct {
		v *int64
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr error
	}{
		{
			name:    "nil",
			args:    args{v: nil},
			want:    0,
			wantErr: bptr.ErrNilPointer,
		},
		{
			name:    "default",
			args:    args{v: bptr.FromInt64(0)},
			want:    0,
			wantErr: nil,
		},
		{
			name:    "with value",
			args:    args{v: bptr.FromInt64(1)},
			want:    1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bptr.ToInt64(tt.args.v)
			if got != tt.want || !errors.Is(err, tt.wantErr) {
				t.Errorf("ToInt64() = %v, %v; want %v, %v", got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestToUint(t *testing.T) {
	type args struct {
		v *uint
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr error
	}{
		{
			name:    "nil",
			args:    args{v: nil},
			want:    0,
			wantErr: bptr.ErrNilPointer,
		},
		{
			name:    "default",
			args:    args{v: bptr.FromUint(0)},
			want:    0,
			wantErr: nil,
		},
		{
			name:    "with value",
			args:    args{v: bptr.FromUint(1)},
			want:    1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bptr.ToUint(tt.args.v)
			if got != tt.want || !errors.Is(err, tt.wantErr) {
				t.Errorf("ToUint() = %v, %v; want %v, %v", got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestToUint8(t *testing.T) {
	type args struct {
		v *uint8
	}
	tests := []struct {
		name    string
		args    args
		want    uint8
		wantErr error
	}{
		{
			name:    "nil",
			args:    args{v: nil},
			want:    0,
			wantErr: bptr.ErrNilPointer,
		},
		{
			name:    "default",
			args:    args{v: bptr.FromUint8(0)},
			want:    0,
			wantErr: nil,
		},
		{
			name:    "with value",
			args:    args{v: bptr.FromUint8(1)},
			want:    1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bptr.ToUint8(tt.args.v)
			if got != tt.want || !errors.Is(err, tt.wantErr) {
				t.Errorf("ToUint8() = %v, %v; want %v, %v", got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestToUint16(t *testing.T) {
	type args struct {
		v *uint16
	}
	tests := []struct {
		name    string
		args    args
		want    uint16
		wantErr error
	}{
		{
			name:    "nil",
			args:    args{v: nil},
			want:    0,
			wantErr: bptr.ErrNilPointer,
		},
		{
			name:    "default",
			args:    args{v: bptr.FromUint16(0)},
			want:    0,
			wantErr: nil,
		},
		{
			name:    "with value",
			args:    args{v: bptr.FromUint16(1)},
			want:    1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bptr.ToUint16(tt.args.v)
			if got != tt.want || !errors.Is(err, tt.wantErr) {
				t.Errorf("ToUint16() = %v, %v; want %v, %v", got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestToUint32(t *testing.T) {
	type args struct {
		v *uint32
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr error
	}{
		{
			name:    "nil",
			args:    args{v: nil},
			want:    0,
			wantErr: bptr.ErrNilPointer,
		},
		{
			name:    "default",
			args:    args{v: bptr.FromUint32(0)},
			want:    0,
			wantErr: nil,
		},
		{
			name:    "with value",
			args:    args{v: bptr.FromUint32(1)},
			want:    1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bptr.ToUint32(tt.args.v)
			if got != tt.want || !errors.Is(err, tt.wantErr) {
				t.Errorf("ToUint32() = %v, %v; want %v, %v", got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestToUint64(t *testing.T) {
	type args struct {
		v *uint64
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr error
	}{
		{
			name:    "nil",
			args:    args{v: nil},
			want:    0,
			wantErr: bptr.ErrNilPointer,
		},
		{
			name:    "default",
			args:    args{v: bptr.FromUint64(0)},
			want:    0,
			wantErr: nil,
		},
		{
			name:    "with value",
			args:    args{v: bptr.FromUint64(1)},
			want:    1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bptr.ToUint64(tt.args.v)
			if got != tt.want || !errors.Is(err, tt.wantErr) {
				t.Errorf("ToUint64() = %v, %v; want %v, %v", got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestToRune(t *testing.T) {
	type args struct {
		v *rune
	}
	tests := []struct {
		name    string
		args    args
		want    rune
		wantErr error
	}{
		{
			name:    "nil",
			args:    args{v: nil},
			want:    0,
			wantErr: bptr.ErrNilPointer,
		},
		{
			name:    "default",
			args:    args{v: bptr.FromRune(0)},
			want:    0,
			wantErr: nil,
		},
		{
			name:    "with value",
			args:    args{v: bptr.FromRune('a')},
			want:    'a',
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bptr.ToRune(tt.args.v)
			if got != tt.want || !errors.Is(err, tt.wantErr) {
				t.Errorf("ToRune() = %v, %v; want %v, %v", got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestToString(t *testing.T) {
	type args struct {
		v *string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name:    "nil",
			args:    args{v: nil},
			want:    "",
			wantErr: bptr.ErrNilPointer,
		},
		{
			name:    "default",
			args:    args{v: bptr.FromString("")},
			want:    "",
			wantErr: nil,
		},
		{
			name:    "with value",
			args:    args{v: bptr.FromString("hello")},
			want:    "hello",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bptr.ToString(tt.args.v)
			if got != tt.want || !errors.Is(err, tt.wantErr) {
				t.Errorf("ToString() = %v, %v; want %v, %v", got, err, tt.want, tt.wantErr)
			}
		})
	}
}

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
			args: args{v: bptr.FromBool(false)},
			want: false,
		},
		{
			name: "with value",
			args: args{v: bptr.FromBool(true)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bptr.ToBoolSafe(tt.args.v); got != tt.want {
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
			args: args{v: bptr.FromByte(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: bptr.FromByte(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bptr.ToByteSafe(tt.args.v); got != tt.want {
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
			args: args{v: bptr.FromComplex64(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: bptr.FromComplex64(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bptr.ToComplex64Safe(tt.args.v); got != tt.want {
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
			args: args{v: bptr.FromComplex128(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: bptr.FromComplex128(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bptr.ToComplex128Safe(tt.args.v); got != tt.want {
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
			args: args{v: bptr.FromFloat32(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: bptr.FromFloat32(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bptr.ToFloat32Safe(tt.args.v); got != tt.want {
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
			args: args{v: bptr.FromFloat64(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: bptr.FromFloat64(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bptr.ToFloat64Safe(tt.args.v); got != tt.want {
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
			args: args{v: bptr.FromInt(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: bptr.FromInt(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bptr.ToIntSafe(tt.args.v); got != tt.want {
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
			args: args{v: bptr.FromInt8(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: bptr.FromInt8(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bptr.ToInt8Safe(tt.args.v); got != tt.want {
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
			args: args{v: bptr.FromInt16(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: bptr.FromInt16(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bptr.ToInt16Safe(tt.args.v); got != tt.want {
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
			args: args{v: bptr.FromInt32(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: bptr.FromInt32(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bptr.ToInt32Safe(tt.args.v); got != tt.want {
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
			args: args{v: bptr.FromInt64(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: bptr.FromInt64(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bptr.ToInt64Safe(tt.args.v); got != tt.want {
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
			args: args{v: bptr.FromRune(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: bptr.FromRune(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bptr.ToRuneSafe(tt.args.v); got != tt.want {
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
			args: args{v: bptr.FromString("")},
			want: "",
		},
		{
			name: "with value",
			args: args{v: bptr.FromString("1")},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bptr.ToStringSafe(tt.args.v); got != tt.want {
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
			args: args{v: bptr.FromUint(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: bptr.FromUint(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bptr.ToUintSafe(tt.args.v); got != tt.want {
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
			args: args{v: bptr.FromUint8(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: bptr.FromUint8(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bptr.ToUint8Safe(tt.args.v); got != tt.want {
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
			args: args{v: bptr.FromUint16(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: bptr.FromUint16(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bptr.ToUint16Safe(tt.args.v); got != tt.want {
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
			args: args{v: bptr.FromUint32(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: bptr.FromUint32(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bptr.ToUint32Safe(tt.args.v); got != tt.want {
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
			args: args{v: bptr.FromUint64(0)},
			want: 0,
		},
		{
			name: "with value",
			args: args{v: bptr.FromUint64(1)},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bptr.ToUint64Safe(tt.args.v); got != tt.want {
				t.Errorf("ToUint64Safe() = %v, want %v", got, tt.want)
			}
		})
	}
}
