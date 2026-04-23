package file_system

import "path/filepath"

var _ PathDescriptor = &DirDescriptor{}

type DirDescriptor struct {
	Path     string
	Parent   PathDescriptor
	Visible  bool
	Children []PathDescriptor
}

func (this DirDescriptor) GetName() string           { return filepath.Base(this.Path) }
func (this DirDescriptor) GetPath() string           { return this.Path }
func (this DirDescriptor) GetParent() PathDescriptor { return this.Parent }
func (this DirDescriptor) IsVisible() bool           { return this.Visible }
func (this *DirDescriptor) SetVisible(v bool)        { this.Visible = v }
func (this DirDescriptor) IsDir() bool               { return true }
func (this DirDescriptor) IsFile() bool              { return false }
func (this DirDescriptor) Count() int                { return len(this.Children) }

func (this *DirDescriptor) GetFile(path string) (*FileDescriptor, bool) {
	for _, child := range this.Children {
		if child.IsDir() && child.GetPath() == path {
			return child.(*FileDescriptor), false
		}
	}

	found := &FileDescriptor{Path: path, Parent: this}
	this.Children = append(this.Children, found)

	return found, true
}

func (this *DirDescriptor) GetDir(path string) (*DirDescriptor, bool) {
	for _, child := range this.Children {
		if child.IsDir() && child.GetPath() == path {
			return child.(*DirDescriptor), false
		}
	}

	found := &DirDescriptor{Path: path, Parent: this}
	this.Children = append(this.Children, found)

	return found, true
}
