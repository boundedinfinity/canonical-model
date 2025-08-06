package main

import (
	"errors"
	"os"
	"runtime"

	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
	"gopkg.in/yaml.v3"
)

func writeFile(path string, data []byte) error {
	rootDir, err := getRootDir()

	if err != nil {
		return err
	}

	filePath := pather.Dirs.Join(rootDir, path)

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return err
	}

	return nil
}

func loadYaml(path string, val any) error {
	rootDir, err := getRootDir()

	if err != nil {
		return err
	}

	yamlPath := pather.Dirs.Join(rootDir, path)

	bs, err := os.ReadFile(yamlPath)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(bs, val)

	if err != nil {
		return err
	}

	return nil
}

func writeYaml(path string, val any) error {
	bs, err := yaml.Marshal(val)

	if err != nil {
		return err
	}

	return writeFile(path, bs)
}

func getRootDir() (string, error) {
	cwd, err := os.Getwd()

	if err != nil {
		return "", err
	}

	return cwd, nil
}

func getWorkingDir() (string, error) {
	// https://stackoverflow.com/a/70050843
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		return "", errors.New("can't determine current directory")
	}

	return pather.Files.Dir(filename), nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
