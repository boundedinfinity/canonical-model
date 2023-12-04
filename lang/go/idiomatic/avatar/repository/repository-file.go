package repository

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/avatar/marshaler"
	"github.com/boundedinfinity/schema/idiomatic/avatar/model"
)

const (
	_MESSENGER_DIR = ".config/avatar/repository"
)

func NewFileMessageRepository(rootDir string, mashaler *marshaler.EventMarshaler) (MessageRepository, error) {
	repository := &fileMessageRepository{
		mashaler: mashaler,
	}

	if rootDir == "" {
		homeDir, err := os.UserHomeDir()

		if err != nil {
			return repository, err
		}

		rootDir = homeDir
	}

	rootDir, err := pather.Dirs.AbsErr(rootDir)

	if err != nil {
		return repository, nil
	}

	projectDir := pather.Join(rootDir, _MESSENGER_DIR)

	if _, err := pather.Dirs.EnsureErr(projectDir); err != nil {
		return repository, err
	}

	repository.logPath = pather.Join(projectDir, "log.json")
	repository.errPath = pather.Join(projectDir, "err.text")

	return repository, nil
}

var _ MessageRepository = &fileMessageRepository{}

type fileMessageRepository struct {
	logPath  string
	errPath  string
	log      []model.RawEvent
	mashaler *marshaler.EventMarshaler
}

func (t *fileMessageRepository) saveLog() error {
	for _, log := range t.log {
		bs, err := json.Marshal(log)

		if err != nil {
			return err
		}

		if err := pather.Files.Appendln(t.logPath, bs, 0755); err != nil {
			panic(err)
		}
	}

	return nil
}

func (t *fileMessageRepository) saveErr(err error) {
	if err := pather.Files.AppendlnString(t.errPath, err.Error(), 0755); err != nil {
		panic(err)
	}
}

func (t *fileMessageRepository) Store(messages ...model.TypedEvent) error {
	raws := slicer.Map(func(message model.TypedEvent) model.RawEvent { return message.Raw() })
	t.log = append(t.log, raws...)
	return t.saveLog()
}

func (t *fileMessageRepository) Close() error {
	return nil
}

func (t *fileMessageRepository) Load(after rfc3339date.Rfc3339DateTime) error {
	file, err := os.Open(t.logPath)

	if err != nil {
		t.saveErr(err)
		return nil
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if err := t.mashaler.Unmarshal(scanner.Bytes()); err != nil {
			t.saveErr(err)
		}
	}

	if err := scanner.Err(); err != nil {
		t.saveErr(err)
	}

	return nil
}
