package file_system

import (
	"path/filepath"

	"github.com/boundedinfinity/canonical_model/idiomatic/digital/mime_type"
)

var _ PathDescriptor = &FileDescriptor{}

type FileDescriptor struct {
	Path   string
	Parent PathDescriptor

	Size int64
	Sum  string
	Type mime_type.MimeTypeModel
}

func (this FileDescriptor) GetName() string           { return filepath.Base(this.Path) }
func (this FileDescriptor) GetPath() string           { return this.Path }
func (this FileDescriptor) GetParent() PathDescriptor { return this.Parent }
func (this FileDescriptor) IsDir() bool               { return false }
func (this FileDescriptor) IsFile() bool              { return true }
