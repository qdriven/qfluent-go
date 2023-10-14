package logging

import (
	"reflect"
	"testing"
)

func TestNewDefaultZapLogger(t *testing.T) {
	tests := []struct {
		name string
		want *zap.Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultZapLogger(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultZapLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}
