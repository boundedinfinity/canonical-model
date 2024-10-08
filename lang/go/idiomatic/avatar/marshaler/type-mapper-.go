package marshaler

import (
	"reflect"

	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
	"github.com/boundedinfinity/go-commoner/idiomatic/reflecter"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type TypeMapperInfo struct {
	langName string
	typeName string
	typ      reflect.Type
}

func NewTypeMapper() *TypeMapper {
	return &TypeMapper{
		langToType: make(mapper.Mapper[string, string]),
		typeToLang: make(mapper.Mapper[string, string]),
		langToInfo: make(mapper.Mapper[string, *TypeMapperInfo]),
		typeToInfo: make(mapper.Mapper[string, *TypeMapperInfo]),
	}
}

type TypeMapper struct {
	langToType mapper.Mapper[string, string]
	typeToLang mapper.Mapper[string, string]
	langToInfo mapper.Mapper[string, *TypeMapperInfo]
	typeToInfo mapper.Mapper[string, *TypeMapperInfo]
}

func (t *TypeMapper) Slice() []TypeMapperInfo {
	return slicer.Map(func(_ int, info *TypeMapperInfo) TypeMapperInfo {
		return *info
	}, t.typeToInfo.Values()...)
}

func (t *TypeMapper) ByValue(val any) (TypeMapperInfo, bool) {
	info, ok := t.langToInfo[reflecter.InstanceQualifiedName(val)]
	return *info, ok
}

func (t *TypeMapper) ByName(name string) (TypeMapperInfo, bool) {
	var info *TypeMapperInfo
	var ok bool

	info, ok = t.langToInfo[name]

	if !ok {
		info, ok = t.typeToInfo[name]
	}

	return *info, ok
}

func (t *TypeMapper) RegisterValue(typeName string, val any) error {
	info := TypeMapperInfo{
		typeName: typeName,
		langName: reflecter.InstanceQualifiedName(val),
		typ:      reflect.TypeOf(val),
	}

	if _, ok := t.langToInfo[info.langName]; !ok {
		t.langToInfo[info.langName] = &info
	}

	if _, ok := t.typeToInfo[info.typeName]; !ok {
		t.typeToInfo[info.typeName] = &info
	}

	if _, ok := t.langToType[info.langName]; !ok {
		t.langToType[info.langName] = info.typeName
	}

	if _, ok := t.typeToLang[info.typeName]; !ok {
		t.typeToLang[info.typeName] = info.langName
	}

	return nil
}

func (t *TypeMapper) RegisterNamer(val id.TypeNamer) error {
	return t.RegisterValue(val.TypeName(), val)
}
