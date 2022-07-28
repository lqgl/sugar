package tasty

import (
	"reflect"
	"testing"
)

func TestCheckInSlice(t *testing.T) {
	type args struct {
		a int32
		s []int32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{"test1", args{a: 5, s: []int32{1, 3, 5, 7, 9}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckInSlice(tt.args.a, tt.args.s); got != tt.want {
				t.Errorf("CheckInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelOneInSlice(t *testing.T) {
	type args struct {
		a   int32
		old []int32
	}
	tests := []struct {
		name    string
		args    args
		wantNew []int32
	}{

		{"test1", args{a: 5, old: []int32{2, 4, 7, 9, 5}}, []int32{2, 4, 7, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNew := DelOneInSlice(tt.args.a, tt.args.old); !reflect.DeepEqual(gotNew, tt.wantNew) {
				t.Errorf("DelOneInSlice() = %v, want %v", gotNew, tt.wantNew)
			}
		})
	}
}

func TestSliceFiltrate(t *testing.T) {
	type args struct {
		collection []int32
		filtrate   func(int32, int) bool
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{

		{
			"test1",
			args{
				collection: []int32{2, 4, 6, 9, 10},
				filtrate: func(s int32, i int) bool {
					if s > 5 {
						return true
					}
					return false
				},
			},
			[]int32{6, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceFiltrate(tt.args.collection, tt.args.filtrate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceFiltrate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceGroupBy(t *testing.T) {
	type args struct {
		collection []int32
		iteratee   func(int32) string
	}
	tests := []struct {
		name string
		args args
		want map[string][]int32
	}{

		{
			"test1",
			args{
				collection: []int32{2, 4, 6, 100, 101, 200},
				iteratee: func(i int32) string {
					if i >= 100 {
						return "gt 100"
					}
					return "lt 100"
				},
			},
			map[string][]int32{"gt 100": {100, 101, 200}, "lt 100": {2, 4, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceGroupBy(tt.args.collection, tt.args.iteratee); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceGroupBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceUniq(t *testing.T) {
	type args struct {
		collection []string
		iteratee   func(string) int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{

		{
			"test1",
			args{
				collection: []string{"abc", "cct", "screw", "glitter"},
				iteratee: func(s string) int {
					return len(s)
				},
			},
			[]string{"abc", "screw", "glitter"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceUniq(tt.args.collection, tt.args.iteratee); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceUniq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceUpdateElement(t *testing.T) {
	type args struct {
		collection []int
		iteratee   func(int, int) int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{

		{
			"test1",
			args{
				collection: []int{2, 1, 4, 5, 52, 51},
				iteratee: func(i int, j int) int {
					return i * 2
				},
			},
			[]int{4, 2, 8, 10, 104, 102},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceUpdateElement(tt.args.collection, tt.args.iteratee); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceUpdateElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
