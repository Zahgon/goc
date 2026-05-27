// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package edit implements buffered position-based editing of byte slices.
package tool

// A Buffer is a queue of edits to apply to a given byte slice.
type Buffer struct {
	old []byte
	q   edits
}

// An edit records a single text modification: change the bytes in [start,end) to new.
type edit struct {
	start int
	end   int
	new   string
}

// An edits is a list of edits that is sortable by start offset, breaking ties by end offset.
type edits []edit

func (x edits) Len() int           { _ = "STUB: not implemented"; return 0 }
func (x edits) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (x edits) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// NewBuffer returns a new buffer to accumulate changes to an initial data slice.
// The returned buffer maintains a reference to the data, so the caller must ensure
// the data is not modified until after the Buffer is done being used.
func NewBuffer(data []byte) *Buffer { _ = "STUB: not implemented"; return nil }

func (b *Buffer) Insert(pos int, new string) { _ = "STUB: not implemented"; return }

func (b *Buffer) Delete(start, end int) { _ = "STUB: not implemented"; return }

func (b *Buffer) Replace(start, end int, new string) { _ = "STUB: not implemented"; return }

// Bytes returns a new byte slice containing the original data
// with the queued edits applied.
func (b *Buffer) Bytes() []byte {
	_ = "STUB: not implemented"
	// Sort edits by starting position and then by ending position.
	// Breaking ties by ending position allows insertions at point x
	// to be applied before a replacement of the text at [x, y).
	return nil
}

// String returns a string containing the original data
// with the queued edits applied.
func (b *Buffer) String() string { _ = "STUB: not implemented"; return "" }
