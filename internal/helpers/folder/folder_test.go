package folder

import "testing"

func TestRemoveAll(t *testing.T) {
	type args struct {
		folder string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "noFolderNamePanic", args: args{folder: ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RemoveAll(tt.args.folder)
		})
	}
}
