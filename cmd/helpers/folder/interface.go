package folder

type Folder interface {
	Chdir(folder string)
	Getwd() string
	RemoveFolders(folders []string)
	RemoveAll(folder string)
	MkdirAll(folder string)
	Mkdir(folder string)
	IsDirectoryEmpty(path string) bool
	Copy(src, dst string)
}
