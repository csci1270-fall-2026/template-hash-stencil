package database

import (
	"dinodb/pkg/cursor"
	"dinodb/pkg/entry"
	"dinodb/pkg/pager"
	"io"
)

// IndexType represents a Hash Table.
type IndexType string

const (
	HashIndexType  IndexType = "hash"
)

// Index interface.
type Index interface {
	Close() error
	GetName() string
	GetPager() *pager.Pager
	Find(int64) (entry.Entry, error)
	Insert(int64, int64) error
	Update(int64, int64) error
	Delete(int64) error
	Select() ([]entry.Entry, error)
	Print(io.Writer)
	PrintPN(int, io.Writer)
	CursorAtStart() (cursor.Cursor, error)
}
