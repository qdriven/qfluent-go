package logging

import (
	"go.uber.org/zap"
	"reflect"
	"testing"
)

func TestLogger(t *testing.T) {
	tests := []struct {
		name string
		want *zap.Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Logger(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Logger() = %v, want %v", got, tt.want)
			}
			Logger().Info("this is default logger")
		})
	}
}

func TestNamedLogger(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *zap.Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NamedLogger(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NamedLogger() = %v, want %v", got, tt.want)
				NamedLogger("named").Info("this is error message")
				NamedLogger("named").Info("this is error message", zap.String("msg", "failed"))
			}
		})
	}
}
