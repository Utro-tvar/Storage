package db

import "errors"

var (
	ErrURLNotFound     = errors.New("url not found")
	ErrURLExists       = errors.New("url exists")
	ErrArticleNotFound = errors.New("article not found")
	ErrArticleExists   = errors.New("article exists")
)
