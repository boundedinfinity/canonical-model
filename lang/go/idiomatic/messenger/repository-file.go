package messenger

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"

	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
	"github.com/boundedinfinity/rfc3339date"
)

const (
	_MESSENGER_DIR = ".config/messenger"
)

func NewFileMessageRepository(logPath string, errPath string) (MessageRepository, error) {
	repository := &fileMessageRepository{}

	if logPath == "" {
		homeDir, err := os.UserHomeDir()

		if err != nil {
			return repository, err
		}

		logPath = pather.Join(homeDir, _MESSENGER_DIR, "repository.log")
	}

	if errPath == "" {
		homeDir, err := os.UserHomeDir()

		if err != nil {
			return repository, err
		}

		logPath = pather.Join(homeDir, _MESSENGER_DIR, "repository.err")
	}

	if file, err := os.Open(logPath); err != nil {
		return repository, err
	} else {
		repository.logPath = logPath
		repository.logFile = file
	}

	if file, err := os.Open(errPath); err != nil {
		return repository, err
	} else {
		repository.logPath = logPath
		repository.logFile = file
	}

	return repository, nil
}

type fileMessageRepository struct {
	logPath string
	logFile *os.File
	errPath string
	errFile *os.File
}

var _ MessageRepository = &fileMessageRepository{}

func (t *fileMessageRepository) Store(messages ...Message) error {
	for _, message := range messages {
		if err := t.store(message); err != nil {
			return err
		}
	}

	return nil
}

func (t *fileMessageRepository) Close() error {
	var errs []error

	closeFile := func(file *os.File) {
		if file != nil {
			if err := file.Close(); err != nil {
				errs = append(errs, err)
			}
		}
	}

	closeFile(t.logFile)
	closeFile(t.errFile)

	return errors.Join(errs...)
}

func (t *fileMessageRepository) store(message Message) error {
	bs, err := json.Marshal(message)

	if err != nil {
		return err
	}

	_, err = t.logFile.Write(bs)

	if err != nil {
		return err
	}

	return nil
}

func (t *fileMessageRepository) Load(after rfc3339date.Rfc3339DateTime) (chan Message, error) {
	ch := make(chan Message)

	go func() {
		defer close(ch)

		scanner := bufio.NewScanner(t.logFile)

		for scanner.Scan() {
			var stored Message

			if err := json.Unmarshal(scanner.Bytes(), &stored); err != nil {
				t.errFile.WriteString(err.Error())
				break
			}

			ch <- stored
		}

		if err := scanner.Err(); err != nil {
			t.errFile.WriteString(err.Error())
		}
	}()

	return ch, nil
}