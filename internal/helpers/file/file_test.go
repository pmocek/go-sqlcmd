package file

import (
	"github.com/microsoft/go-sqlcmd/internal/helpers/folder"
	"path/filepath"
	"strings"
	"testing"
)

func TestCreateEmptyFileIfNotExists(t *testing.T) {
	filename := "foo.txt"
	folderName := "folder"

	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "default", args: args{filename: filename}},
		{name: "alreadyExists", args: args{filename: filename}},
		{name: "emptyInputPanic", args: args{filename: ""}},
		{name: "incFolder", args: args{filename: filepath.Join(folderName, filename)}},
	}

	cleanup(folderName, filename)
	defer cleanup(folderName, filename)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// If test name ends in 'Panic' expect a Panic
			if strings.HasSuffix(tt.name, "Panic") {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}

			CreateEmptyFileIfNotExists(tt.args.filename)
		})
	}
}

func TestExists(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name       string
		args       args
		wantExists bool
	}{
		{name: "exists", args: args{filename: "file_test.go"}, wantExists: true},
		{name: "notExists", args: args{filename: "does-not-exist.file"}, wantExists: false},
		{name: "noFilenamePanic", args: args{filename: ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// If test name ends in 'Panic' expect a Panic
			if strings.HasSuffix(tt.name, "Panic") {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}

			if gotExists := Exists(tt.args.filename); gotExists != tt.wantExists {
				t.Errorf("Exists() = %v, want %v", gotExists, tt.wantExists)
			}
		})
	}
}

func cleanup(folderName string, filename string) {
	if Exists(folderName) {
		folder.RemoveAll(folderName)
	}

	if Exists(filename) {
		Remove(filename)
	}
}
