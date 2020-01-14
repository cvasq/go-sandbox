package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find word")

func (d Dictionary) Search(searchKey string) (string, error) {
	definition, ok := d[searchKey]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
