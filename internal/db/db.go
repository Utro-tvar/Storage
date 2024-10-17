package db

import "errors"

var (
	ErrResourceNotFound = errors.New("resource not found")
	ErrResourceExists   = errors.New("resource exists")
)

const (
	File = "File"
	URL  = "URL"
)
