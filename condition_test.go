package tasty

import (
	"reflect"
	"testing"
)

func TestIf(t *testing.T) {
	type args struct {
		condition bool
		result    int32
	}
	tests := []struct {
		name string
		args args
		want *IfElse[int32]
	}{

		{
			"test1",
			args{
				condition: 1 == 1,
				result:    10,
			},
			&IfElse[int32]{10, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := If(tt.args.condition, tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("If() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTernary(t *testing.T) {
	type args struct {
		condition  bool
		ifOutput   int32
		elseOutput int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{

		{
			"test1",
			args{
				condition:  1 == 1,
				ifOutput:   5,
				elseOutput: 10,
			},
			5,
		},
		{
			"test1",
			args{
				condition:  1 == 2,
				ifOutput:   5,
				elseOutput: 10,
			},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ternary(tt.args.condition, tt.args.ifOutput, tt.args.elseOutput); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ternary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIfElse_Else(t *testing.T) {
	type fields struct {
		Result int32
		Ok     bool
	}
	type args struct {
		result int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int32
	}{

		{
			"test1",
			fields{
				Result: 50,
				Ok:     false,
			},
			args{result: 100},
			100,
		},
		{
			"test1",
			fields{
				Result: 50,
				Ok:     true,
			},
			args{result: 100},
			50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IfElse[int32]{
				Result: tt.fields.Result,
				Ok:     tt.fields.Ok,
			}
			if got := i.Else(tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Else() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIfElse_ElseFn(t *testing.T) {
	type fields struct {
		Result int32
		Ok     bool
	}
	type args struct {
		resultFn func() int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int32
	}{

		{
			"test1",
			fields{
				Result: 50,
				Ok:     false,
			},
			args{
				resultFn: func() int32 {
					return 100
				},
			},
			100,
		},
		{
			"test1",
			fields{
				Result: 50,
				Ok:     true,
			},
			args{
				resultFn: func() int32 {
					return 100
				},
			},
			50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IfElse[int32]{
				Result: tt.fields.Result,
				Ok:     tt.fields.Ok,
			}
			if got := i.ElseFn(tt.args.resultFn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ElseFn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIfElse_ElseIf(t *testing.T) {
	type fields struct {
		Result int32
		Ok     bool
	}
	type args struct {
		condition bool
		result    int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *IfElse[int32]
	}{

		{
			"test1",
			fields{
				Result: 0,
				Ok:     false,
			},
			args{
				condition: true,
				result:    100,
			},
			&IfElse[int32]{100, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IfElse[int32]{
				Result: tt.fields.Result,
				Ok:     tt.fields.Ok,
			}
			if got := i.ElseIf(tt.args.condition, tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ElseIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIfElse_ElseIfFn(t *testing.T) {
	type fields struct {
		Result int32
		Ok     bool
	}
	type args struct {
		condition bool
		resultFn  func() int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *IfElse[int32]
	}{

		{
			"test1",
			fields{
				Result: 10,
				Ok:     false,
			},
			args{
				condition: true,
				resultFn: func() int32 {
					return 100
				},
			},
			&IfElse[int32]{100, true},
		},
		{
			"test2",
			fields{
				Result: 50,
				Ok:     true,
			},
			args{
				condition: false,
				resultFn: func() int32 {
					return 100
				},
			},
			&IfElse[int32]{50, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IfElse[int32]{
				Result: tt.fields.Result,
				Ok:     tt.fields.Ok,
			}
			if got := i.ElseIfFn(tt.args.condition, tt.args.resultFn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ElseIfFn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIfFn(t *testing.T) {
	type args struct {
		condition bool
		resultFn  func() int32
	}
	tests := []struct {
		name string
		args args
		want *IfElse[int32]
	}{

		{
			"test1",
			args{
				condition: false,
				resultFn: func() int32 {
					return 0
				},
			},
			&IfElse[int32]{0, false},
		},
		{
			"test1",
			args{
				condition: true,
				resultFn: func() int32 {
					return 10
				},
			},
			&IfElse[int32]{10, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IfFn[int32](tt.args.condition, tt.args.resultFn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IfFn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwitch(t *testing.T) {
	type args struct {
		predicate int32
	}
	tests := []struct {
		name string
		args args
		want *SwitchCase[int32, int32]
	}{

		{
			"test1",
			args{10},
			&SwitchCase[int32, int32]{10, 0, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Switch[int32, int32](tt.args.predicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Switch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwitchCase_Case(t *testing.T) {
	type fields struct {
		Predicate int32
		Result    int32
		Ok        bool
	}
	type args struct {
		value  int32
		result int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SwitchCase[int32, int32]
	}{

		{
			"test1",
			fields{
				Predicate: 15,
				Result:    10,
				Ok:        false,
			},
			args{
				15,
				100,
			},
			&SwitchCase[int32, int32]{15, 100, true},
		},
		{
			"test2",
			fields{
				Predicate: 15,
				Result:    10,
				Ok:        true,
			},
			args{
				int32(0),
				15,
			},
			&SwitchCase[int32, int32]{15, 10, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SwitchCase[int32, int32]{
				Predicate: tt.fields.Predicate,
				Result:    tt.fields.Result,
				Ok:        tt.fields.Ok,
			}
			if got := s.Case(tt.args.value, tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Case() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwitchCase_CaseFn(t *testing.T) {
	type fields struct {
		Predicate int32
		Result    int32
		Ok        bool
	}
	type args struct {
		value    int32
		resultFn func() int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SwitchCase[int32, int32]
	}{

		{
			"test1",
			fields{
				Predicate: 15,
				Result:    10,
				Ok:        true,
			},
			args{
				int32(0),
				func() int32 {
					return 0
				},
			},
			&SwitchCase[int32, int32]{15, 10, true},
		},
		{
			"test2",
			fields{
				Predicate: 15,
				Result:    10,
				Ok:        false,
			},
			args{
				int32(0),
				func() int32 {
					return 0
				},
			},
			&SwitchCase[int32, int32]{15, 10, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SwitchCase[int32, int32]{
				Predicate: tt.fields.Predicate,
				Result:    tt.fields.Result,
				Ok:        tt.fields.Ok,
			}
			if got := s.CaseFn(tt.args.value, tt.args.resultFn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CaseFn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwitchCase_Default(t *testing.T) {
	type fields struct {
		Predicate int32
		Result    int32
		Ok        bool
	}
	type args struct {
		result int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int32
	}{

		{
			"test1",
			fields{
				Predicate: 15,
				Result:    10,
				Ok:        true,
			},
			args{
				int32(0),
			},
			10,
		},
		{
			"test2",
			fields{
				Predicate: 15,
				Result:    10,
				Ok:        false,
			},
			args{
				int32(0),
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SwitchCase[int32, int32]{
				Predicate: tt.fields.Predicate,
				Result:    tt.fields.Result,
				Ok:        tt.fields.Ok,
			}
			if got := s.Default(tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Default() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwitchCase_DefaultFn(t *testing.T) {
	type fields struct {
		Predicate int32
		Result    int32
		Ok        bool
	}
	type args struct {
		resultFn func() int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int32
	}{

		{
			"test1",
			fields{
				Predicate: 15,
				Result:    10,
				Ok:        true,
			},
			args{
				func() int32 {
					return 0
				},
			},
			10,
		},
		{
			"test2",
			fields{
				Predicate: 15,
				Result:    10,
				Ok:        false,
			},
			args{
				func() int32 {
					return 10
				},
			},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SwitchCase[int32, int32]{
				Predicate: tt.fields.Predicate,
				Result:    tt.fields.Result,
				Ok:        tt.fields.Ok,
			}
			if got := s.DefaultFn(tt.args.resultFn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultFn() = %v, want %v", got, tt.want)
			}
		})
	}
}
