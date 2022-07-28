package tasty

import (
	"reflect"
	"testing"
)

func TestAssign(t *testing.T) {
	type args struct {
		maps []map[string]int32
	}
	tests := []struct {
		name string
		args args
		want map[string]int32
	}{

		{
			"test1",
			args{maps: []map[string]int32{
				{"a": 1, "b": 2},
				{"b": 5, "d": 6},
			}},
			map[string]int32{"a": 1, "b": 5, "d": 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Assign(tt.args.maps...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Assign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntriesToMap(t *testing.T) {
	type args struct {
		entries []Entry[string, int32]
	}
	tests := []struct {
		name string
		args args
		want map[string]int32
	}{

		{
			"test1",
			args{
				[]Entry[string, int32]{{"a", 1}, {"b", 2}, {"c", 3}},
			},
			map[string]int32{"a": 1, "b": 2, "c": 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EntriesToMap(tt.args.entries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EntriesToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFiltrateBy(t *testing.T) {
	type args struct {
		in       map[int32]string
		filtrate func(int32, string) bool
	}
	tests := []struct {
		name string
		args args
		want map[int32]string
	}{

		{
			"test1",
			args{
				in: map[int32]string{1: "a", 2: "b", 3: "c"},
				filtrate: func(i int32, s string) bool {
					return i > 2
				},
			},
			map[int32]string{3: "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FiltrateBy(tt.args.in, tt.args.filtrate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FiltrateBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFiltrateByKeys(t *testing.T) {
	type args struct {
		in   map[int32]string
		keys []int32
	}
	tests := []struct {
		name string
		args args
		want map[int32]string
	}{

		{
			"test1",
			args{
				in:   map[int32]string{1: "a", 2: "b", 3: "c"},
				keys: []int32{2, 3},
			},
			map[int32]string{2: "b", 3: "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FiltrateByKeys(tt.args.in, tt.args.keys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FiltrateByKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFiltrateByValues(t *testing.T) {
	type args struct {
		in     map[string]int32
		values []int32
	}
	tests := []struct {
		name string
		args args
		want map[string]int32
	}{

		{
			"test1",
			args{
				in:     map[string]int32{"a": 1, "b": 2, "c": 3},
				values: []int32{2, 3},
			},
			map[string]int32{"b": 2, "c": 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FiltrateByValues(tt.args.in, tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FiltrateByValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInvert(t *testing.T) {
	type args struct {
		in map[string]int32
	}
	tests := []struct {
		name string
		args args
		want map[int32]string
	}{

		{
			"test1",
			args{in: map[string]int32{"a": 1, "b": 2, "c": 3}},
			map[int32]string{1: "a", 2: "b", 3: "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Invert(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Invert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeys(t *testing.T) {
	type args struct {
		in map[string]int32
	}
	tests := []struct {
		name string
		args args
		want []string
	}{

		{
			"test1",
			args{in: map[string]int32{"a": 1, "b": 2, "c": 3}},
			[]string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Keys(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapToEntries(t *testing.T) {
	type args struct {
		in map[string]int32
	}
	tests := []struct {
		name string
		args args
		want []Entry[string, int32]
	}{

		{
			"test1",
			args{map[string]int32{"a": 1, "b": 2, "c": 3}},
			[]Entry[string, int32]{{"a", 1}, {"b", 2}, {"c", 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapToEntries(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapToEntries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapUpdateKeys(t *testing.T) {
	type args struct {
		in       map[int32]int32
		iteratee func(int32, int32) int32
	}
	tests := []struct {
		name string
		args args
		want map[int32]int32
	}{

		{
			"test1",
			args{
				in: map[int32]int32{2: 1, 4: 2, 6: 3, 8: 4},
				iteratee: func(i int32, j int32) int32 {
					return i / 2
				},
			},
			map[int32]int32{1: 1, 2: 2, 3: 3, 4: 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapUpdateKeys(tt.args.in, tt.args.iteratee); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapUpdateKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapUpdateValues(t *testing.T) {
	type args struct {
		in       map[int32]int32
		iteratee func(int32, int32) int32
	}
	tests := []struct {
		name string
		args args
		want map[int32]int32
	}{

		{
			"map val multiple 2 test",
			args{
				in: map[int32]int32{1: 1, 2: 2, 3: 3, 4: 4},
				iteratee: func(i int32, j int32) int32 {
					return j * 2
				},
			},
			map[int32]int32{1: 2, 2: 4, 3: 6, 4: 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapUpdateValues(tt.args.in, tt.args.iteratee); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapUpdateValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValues(t *testing.T) {
	type args struct {
		in map[string]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{

		{
			"test1",
			args{
				map[string]int{"0": 1, "1": 1, "2": 2},
			},
			[]int{1, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Values(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}
