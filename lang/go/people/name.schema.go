package main

import (
	"fmt"
	"strings"
)

type Name struct {
	Prefixes []string
	Names    []string
	Suffixes []string
}

func (t Name) Validate() error {
	for i, s := range t.Prefixes {
		if len(s) < 1 || len(s) > 10 {
			return fmt.Errorf("prefixes[%v] must be between 1 and 10 characters in lenght", i)
		}
	}

	if len(t.Names) < 1 {
		return fmt.Errorf("names must contain at least 1 item")
	}

	for i, s := range t.Names {
		if len(s) < 1 || len(s) > 10 {
			return fmt.Errorf("names[%v] must be between 1 and 10 characters in lenght", i)
		}
	}

	for i, s := range t.Suffixes {
		if len(s) < 1 || len(s) > 10 {
			return fmt.Errorf("suffixes[%v] must be between 1 and 10 characters in lenght", i)
		}
	}

	return nil
}

func (t Name) FullName() string {
	var name string

	if len(t.Prefixes) > 0 {
		name += strings.Join(t.Prefixes, " ")
		name += " "
	}

	if len(t.Names) > 0 {
		name += strings.Join(t.Names, " ")
		name += " "
	}

	if len(t.Suffixes) > 0 {
		name += strings.Join(t.Suffixes, " ")
	}

	return name
}
