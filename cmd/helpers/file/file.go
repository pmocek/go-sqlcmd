package file

import (
	"bufio"
	"fmt"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/folder"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func (f *fh) AssertExists(filename string, minSize int64) {
	fi, err := os.Stat(filename)

	if  os.IsNotExist(err) {
		panic(filepath.Join(getwd(), filename) + ": does not exist")
	}
	if fi.Size() < minSize {
		panic("filename: " + filename + " is not large enough")
	}
}

func (f *fh) CreateEmptyIfNotExists(filename string) {
	if !f.Exists(filename) {
		folder := folder.GetInstance()
		folder.MkdirAll(filepath.Base(filename))
		f, err := os.Create(filename)
		defer f.Close()
		checkErr(err)
	}
}

func (f *fh) Exists(filename string) (exists bool) {
	if _, err := os.Stat(filename); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}

func (f *fh) Remove(filename string) {
	err := os.Remove(filename)
	if err != nil {
		panic(fmt.Sprintf("Unable to remove file '%s'. %s (current working directory: '%s')", filename, err, getwd()))
	}
}

func (f *fh) OpenFile(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		panic(fmt.Sprintf("openFile: %s (current working directory: '%s')", err, getwd()))
	}

	return file
}

func (fh *fh) AppendLine(filename string, text string) {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(text + "\n")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func (fh *fh) WriteString(f *os.File, s string) {
	_, err := f.WriteString(s)
	if err != nil {
		panic("Unable to write string")
	}
}

func (fh *fh) CloseFile(f *os.File) {
	err := f.Close()
	if err != nil {
		panic(fmt.Sprintf("Unable to close file: %s", f.Name()))
	}
}

func (f *fh) GetContents(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	content := string(b)

	// Simulate symlinks for windows (enables repos created on Linux to be built on Windows)
	if runtime.GOOS == "windows" {
		newlines := 0
		for _, c := range content {
			if c == '\n' {
				newlines++
			}

			if newlines > 1 {
				break
			}
		}

		if newlines <= 1 {
			// this is a single line file, it is probably a symlink, so see if it exists, if it does follow it
			if f.Exists(content) {
				return f.GetContents(content)
			}
		}
	}

	return content
}

func (f *fh) GetBytes(filename string) []byte {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	content := string(b)

	// Simulate symlinks for windows (enables repos created on Linux to be built on Windows)
	if runtime.GOOS == "windows" {
		newlines := 0
		for _, c := range content {
			if c == '\n' {
				newlines++
			}

			if newlines > 1 {
				break
			}
		}

		content = strings.ReplaceAll(content, "/", "\\")
		d, _ := filepath.Split(filename)
	    content = filepath.Join(d, content)

		if newlines <= 1 {
			// this is a single line file, it is probably a symlink, so see if it exists, if it does follow it
			if f.Exists(content) {
				return f.GetBytes(content)
			}
		}
	}

	return b
}

// Readlines is an iterator that returns one line of a file at a time.
func (fh *fh) Readlines(path string) (*os.File, <-chan string) {
	f, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("%s (filename: %v, current working directory: %s)", err, path, getwd()))
	}

	scanner := bufio.NewScanner(f)
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	c := make(chan string)
	go func() {
		for scanner.Scan() {
			c <- scanner.Text()
		}
		close(c)
	}()

	return f, c
}

func getwd() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return path
}
