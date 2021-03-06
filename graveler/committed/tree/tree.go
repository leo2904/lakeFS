package tree

import (
	"github.com/treeverse/lakefs/graveler"
	"github.com/treeverse/lakefs/graveler/committed"
)

// part is the basic building stone of a tree
// nolint: structcheck, unused
type Part struct {
	Name   committed.ID
	MaxKey graveler.Key
}

// Tree is a sorted slice of parts with no overlapping between the parts
// nolint: structcheck, unused
type Tree struct {
	ID    graveler.TreeID
	Parts []Part
}

// Repo is an abstraction for a repository of trees that exposes operations on them
type Repo interface {
	GetTree(treeID graveler.TreeID) (*Tree, error)

	// GetValue finds the matching graveler.ValueRecord in the tree with the treeID
	GetValue(treeID graveler.TreeID, key graveler.Key) (*graveler.ValueRecord, error)

	// NewTreeWriter returns a writer that is used for creating new trees
	NewTreeWriter() Writer

	// NewIterator accepts a tree ID, and returns an iterator
	// over the tree from the first value GE than the from
	NewIterator(treeID graveler.TreeID, from graveler.Key) (graveler.ValueIterator, error)

	// NewIteratorFromTree accept a tree in memory, returns an iterator
	// over the tree from the first value GE than the from
	NewIteratorFromTree(tree Tree, from graveler.Key) (graveler.ValueIterator, error)

	// GetIterForPart accepts a tree ID and a reading start point. it returns an iterator
	// positioned at the start point. When Next() will be called, first value that is GE
	// than the from key will be returned
	NewPartIterator(partID committed.ID, from graveler.Key) (graveler.ValueIterator, error)

	// RemoveCommonParts accepts the left and right trees of the diff, and finds the common parts which
	// exist in both trees.
	// it returns the left and right trees with common parts filtered.
	RemoveCommonParts(Left, Right Tree) (*Tree, *Tree, error)
}

// Writer is an abstraction for creating new trees
type Writer interface {
	// WriteRecord adds a record to the tree. The key key must be greater than any other key that was written
	// (in other words - values must be entered sorted by key order).
	// If the most recent insertion was using AddParts, the key must be greater than any key in the added parts.
	WriteRecord(record graveler.ValueRecord) error

	// AddParts adds complete parts to the tree at the current insertion point.
	// Added parts must not contain keys smaller than last previously written value.
	AddParts(parts []Part) error

	// FlushIterToTree writes the content of an iterator to the tree.
	FlushIterToTree(iter graveler.ValueIterator) error

	// SaveTree finalizes the tree creation. It's invalid to add records after calling this method.
	// During tree writing, parts are closed asynchronously and copied by tierFS
	// while writing continues. SaveTree waits until closing and copying all parts.
	SaveTree() (*graveler.TreeID, error)
}
