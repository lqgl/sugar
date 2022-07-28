package tasty

import (
	"reflect"
	"testing"
)

func TestClamp(t *testing.T) {
	type args struct {
		value int32
		min   int32
		max   int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{

		{
			"test1",
			args{
				value: 20,
				min:   100,
				max:   200,
			},
			100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clamp(tt.args.value, tt.args.min, tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
