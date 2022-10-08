package folder

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func (f *fh) Chdir(folder string) {

	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err := os.Mkdir(folder, 0700)
		if err != nil {
			panic(err)
		}
	}

	err := os.Chdir(folder)
	if err != nil {
		panic(err)
	}
}

//Getwd returns the current working directory
func (f *fh) Getwd() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return path
}

func (f *fh) RemoveFolders(folders []string) {
	for _, folder := range folders {
		f.RemoveAll(folder)
	}
}

func (f *fh) RemoveAll(folder string) {
	err := os.RemoveAll(folder)
	if err != nil {
		panic(fmt.Sprintf("Unable to remove file/folder '%s'. %s (current working directory: '%s')",folder, err, f.Getwd()))
	}
}

func (f *fh) MkdirAll(folder string) {
	if folder == "" {
		return
	}
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err := os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			panic(fmt.Sprintf("Unable to create folder '%v'. %v", folder, err))
		}
	}
}

func (f *fh) Mkdir(folder string) {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err := os.Mkdir(folder, 0700)
		if err != nil {
			panic(err)
		}
	}
}

// IsDirectoryEmpty returns true if there are no files in 'path'
//
// SPECIAL CASE: Consider directory empty if the only thing in the directory is
//	the 'output' directory, (created from a previous failed attempt to init)
func (f *fh) IsDirectoryEmpty(path string) bool {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	if len(files) == 1 {
		if files[0].Name() == "output" {
			return true
		}
	}

	if len(files) > 0 {
		return false
	} else {
		return true
	}
}

// Copy the src file to dst. Any existing file will be overwritten and will not
// copy file attributes.
func (f *fh) Copy(src, dst string) {
	in, err := os.Open(src)
	if err != nil {
		panic(err.Error() + "\nThe current working directory: " + f.Getwd())
	}
	defer func() {
		err := in.Close()
		if err != nil {
			panic(err)
		}
	}()

	out, err := os.Create(dst)
	if err != nil {
		panic(err.Error() + "\nThe current working directory: " + f.Getwd())
	}
	defer func() {
		err := out.Close()
		if err != nil {
			panic(err)
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		panic(err)
	}
}
