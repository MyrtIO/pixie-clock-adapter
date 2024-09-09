package pointer_test

import (
	"pixie_adapter/pkg/pointer"
	"testing"
)

func TestFrom(t *testing.T) {
	type args struct {
		value interface{}
	}
	intValue := 1
	stringValue := "test"
	boolValue := true
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "int",
			args: args{
				value: intValue,
			},
			want: intValue,
		},
		{
			name: "string",
			args: args{
				value: stringValue,
			},
			want: stringValue,
		},
		{
			name: "bool",
			args: args{
				value: boolValue,
			},
			want: boolValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pointer.From(tt.args.value)
			if *got != tt.want {
				t.Errorf("From() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestToValue(t *testing.T) {
	intValue := 1
	stringValue := "test"
	boolValue := true

	if pointer.ToValue(&intValue) != intValue {
		t.Errorf("ToValue() = %v, want %v", pointer.ToValue(&intValue), intValue)
	}
	if pointer.ToValue(&stringValue) != stringValue {
		t.Errorf("ToValue() = %v, want %v", pointer.ToValue(&stringValue), stringValue)
	}
	if pointer.ToValue(&boolValue) != boolValue {
		t.Errorf("ToValue() = %v, want %v", pointer.ToValue(&boolValue), boolValue)
	}

	var dummy *int
	if pointer.ToValue(dummy) != 0 {
		t.Errorf("ToValue() = %v, want %v", pointer.ToValue(dummy), 0)
	}

	type testStruct struct {
		a int
		b string
		c bool
	}

	var testPtr *testStruct
	testValue := pointer.ToValue(testPtr)
	if testValue.a != 0 || testValue.b != "" || testValue.c != false {
		t.Errorf("ToValue() = %v, want %v", testValue, 0)
	}
}
