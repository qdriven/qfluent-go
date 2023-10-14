package utils

import "testing"

func TestDirectoryExist(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DirectoryExist(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("DirectoryExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DirectoryExist() got = %v, want %v", got, tt.want)
			}
		})
	}
}
