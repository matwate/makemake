package main

import (
	"io/fs"
	"os"
	"reflect"
	"testing"
)

func Test_getExtensions(t *testing.T) {
	type args struct {
		files []fs.DirEntry
	}
	tests := []struct {
		name string
		args args
		want map[string]bool
	}{
		{
			name: "Test Case 1",
			args: args{
				files: first_arg(getFiles()),
			},
			want: map[string]bool{
				".go": true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getExtensions(tt.args.files); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getExtensions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFiles(t *testing.T) {
	tests := []struct {
		name    string
		want    []os.DirEntry
		wantErr bool
	}{
		{
			name:    "Test Case 1",
			want:    first_arg(os.ReadDir(".")),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getFiles()
			if (err != nil) != tt.wantErr {
				t.Errorf("getFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
