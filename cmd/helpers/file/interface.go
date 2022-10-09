package file

import "os"

type File interface {
	AppendLine(filename string, text string)
	AssertExists(filename string, minSize int64)
	CreateEmptyIfNotExists(filename string)
	Exists(filename string) (exists bool)
	Remove(filename string)
	OpenFile(filename string) *os.File
	WriteString(f *os.File, s string)
	CloseFile(f *os.File)
	GetContents(filename string) string
	GetBytes(filename string) []byte
	Readlines(path string) (*os.File, <-chan string)
}
