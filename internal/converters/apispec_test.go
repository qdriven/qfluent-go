package converters

import "testing"

func TestOpenApi3ToPostman(t *testing.T) {
	type args struct {
		swaggerFilePath string
		postmanFilePath string
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OpenApi3ToPostman(tt.args.swaggerFilePath, tt.args.postmanFilePath)
		})
	}
}

func TestSwaggerToPostman(t *testing.T) {
	type args struct {
		swaggerFilePath string
		postmanFilePath string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SwaggerToPostman(tt.args.swaggerFilePath, tt.args.postmanFilePath)
		})
	}
}
