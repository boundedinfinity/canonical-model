package file_system

type PathDescriptor interface {
	GetName() string
	GetPath() string
	GetParent() PathDescriptor
	IsDir() bool
	IsFile() bool
}
