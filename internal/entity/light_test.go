package entity_test

import (
	"pixie_adapter/internal/entity"
	"reflect"
	"testing"
)

func TestRGBColorFromSlice(t *testing.T) {
	type args struct {
		values []uint8
	}
	tests := []struct {
		name string
		args args
		want entity.RGBColor
	}{
		{
			name: "RGBColorFromSlice",
			args: args{
				values: []uint8{10, 20, 30},
			},
			want: entity.RGBColor{
				R: 10,
				G: 20,
				B: 30,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := entity.RGBColorFromSlice(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RGBColorFromSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLightPowerStateFromBool(t *testing.T) {
	type args struct {
		enabled bool
	}
	tests := []struct {
		name string
		args args
		want entity.LightPowerState
	}{
		{
			name: "LightPowerStateFromBool",
			args: args{
				enabled: true,
			},
			want: entity.LightPowerStateOn,
		},
		{
			name: "LightPowerStateFromBool",
			args: args{
				enabled: false,
			},
			want: entity.LightPowerStateOff,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := entity.LightPowerStateFromBool(tt.args.enabled); got != tt.want {
				t.Errorf("LightPowerStateFromBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLightPowerStateToBool(t *testing.T) {
	type args struct {
		state entity.LightPowerState
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "LightPowerStateToBool",
			args: args{
				state: entity.LightPowerStateOn,
			},
			want: true,
		},
		{
			name: "LightPowerStateToBool",
			args: args{
				state: entity.LightPowerStateOff,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.state.Bool(); got != tt.want {
				t.Errorf("LightPowerStateToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLightEffectFromCode(t *testing.T) {
	type args struct {
		code uint8
	}
	tests := []struct {
		name string
		args args
		want entity.LightEffect
	}{
		{
			name: "LightEffectFromCode",
			args: args{
				code: 0,
			},
			want: entity.LightEffectStatic,
		},
		{
			name: "LightEffectFromCode",
			args: args{
				code: 1,
			},
			want: entity.LightEffectSmooth,
		},
		{
			name: "LightEffectFromCode",
			args: args{
				code: 2,
			},
			want: entity.LightEffectZoom,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := entity.LightEffectFromCode(tt.args.code); got != tt.want {
				t.Errorf("LightEffectFromCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLightEffectCode(t *testing.T) {
	type args struct {
		effect entity.LightEffect
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			name: "Static light effect code",
			args: args{
				effect: entity.LightEffectStatic,
			},
			want: 0,
		},
		{
			name: "Smooth light effect code",
			args: args{
				effect: entity.LightEffectSmooth,
			},
			want: 1,
		},
		{
			name: "Zoom light effect code",
			args: args{
				effect: entity.LightEffectZoom,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.effect.Code(); got != tt.want {
				t.Errorf("LightEffectCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLightEffectToCode(t *testing.T) {
	type args struct {
		effect entity.LightEffect
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			name: "LightEffectToCode",
			args: args{
				effect: entity.LightEffectStatic,
			},
			want: 0,
		},
		{
			name: "LightEffectToCode",
			args: args{
				effect: entity.LightEffectSmooth,
			},
			want: 1,
		},
		{
			name: "LightEffectToCode",
			args: args{
				effect: entity.LightEffectZoom,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.effect.Code(); got != tt.want {
				t.Errorf("LightEffectToCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
