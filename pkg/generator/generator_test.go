package generator

import "testing"

func TestGenerate(t *testing.T) {
	type args struct {
		transformationsFile string
		source              string
		destination         string
		inputArgs           []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Generate(tt.args.transformationsFile, tt.args.source, tt.args.destination, tt.args.inputArgs); (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
