// Copyright 2018-2025 The Olric Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package table

import (
	"sync"

	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/olric-data/olric/pkg/storage"
	"github.com/pkg/errors"
)

const (
	// MaxKeyLength is the maximum allowed key size in bytes.
	MaxKeyLength = 256

	// MetadataLength is the fixed number of bytes used to store per-entry metadata
	// (TTL + Timestamp + LastAccess + ValueLength + KeyLength = 8+8+8+4+1 = 29).
	MetadataLength = 29
)

// State represents the operational state of a Table.
type State uint8

const (
	// ReadWriteState indicates the table accepts both reads and writes.
	ReadWriteState = State(iota + 1)

	// ReadOnlyState indicates the table only accepts read operations.
	ReadOnlyState

	// RecycledState indicates the table has been reset and is ready for reuse.
	RecycledState
)

var (
	// ErrNotEnoughSpace is returned when the table's pre-allocated memory buffer
	// does not have enough room to store a new entry.
	ErrNotEnoughSpace = errors.New("not enough space")

	// ErrHKeyNotFound is returned when the given hash key does not exist in the table.
	ErrHKeyNotFound = errors.New("hkey not found")
)

// Stats holds memory usage statistics and metadata for a Table.
type Stats struct {
	// Allocated is the total size of the pre-allocated memory buffer in bytes.
	Allocated uint64

	// Inuse is the number of bytes currently occupied by active entries.
	Inuse uint64

	// Garbage is the number of bytes occupied by deleted entries that have not been reclaimed.
	Garbage uint64

	// Length is the number of active entries in the table.
	Length int

	// RecycledAt is the UnixNano timestamp of the last Reset call, or zero if never recycled.
	RecycledAt int64
}

// Table is an in-memory key-value store backed by a pre-allocated byte slice.
// Entries are written sequentially into the buffer using a compact binary layout:
//
//	KEY-LENGTH(uint8) | KEY(bytes) | TTL(uint64) | TIMESTAMP(uint64) | LASTACCESS(uint64) | VALUE-LENGTH(uint32) | VALUE(bytes)
//
// A hash key (uint64) to offset mapping provides O(1) lookups. Deleted entries
// are tracked as garbage but not reclaimed until the table is compacted or recycled.
type Table struct {
	lastAccessMtx sync.RWMutex
	coefficient   uint64
	offset        uint64
	allocated     uint64
	inuse         uint64
	garbage       uint64
	recycledAt    int64
	state         State
	hkeys         map[uint64]uint64
	offsetIndex   *roaring64.Bitmap
	memory        []byte
}

// New creates a new Table with a pre-allocated memory buffer of the given size in bytes.
func New(size uint64) *Table { _ = "STUB: not implemented"; return nil }

//  From builtin.go:
//
//  The size specifies the length. The capacity of the slice is
//	equal to its length. A second integer argument may be provided to
//	specify a different capacity; it must be no smaller than the
//	length. For example, make([]int, 0, 10) allocates an underlying array
//	of size 10 and returns a slice of length 0 and capacity 10 that is
//	backed by this underlying array.

// SetCoefficient sets the coefficient value used for load-balancing and distribution purposes.
func (t *Table) SetCoefficient(cf uint64) {
	_ = "STUB: not implemented"

	// Coefficient returns the current coefficient value of the table.
	return
}

func (t *Table) Coefficient() uint64 { _ = "STUB: not implemented"; return 0 }

// SetState sets the operational state of the table.
func (t *Table) SetState(s State) {
	_ = "STUB: not implemented"

	// State returns the current operational state of the table.
	return
}

func (t *Table) State() State {
	_ = "STUB: not implemented"

	// PutRaw stores pre-encoded raw bytes into the table under the given hash key.
	// It copies the value directly into the memory buffer without any metadata encoding.
	// Returns ErrNotEnoughSpace if the buffer cannot accommodate the value.
	return *new(State)
}

func (t *Table) PutRaw(hkey uint64, value []byte) error {
	_ = "STUB: not implemented"
	// Check empty space on the allocated memory area.
	return nil
}

// Put stores a storage.Entry into the table under the given hash key. It encodes
// the entry's key, TTL, timestamp, last access time and value into the memory buffer
// using the following binary layout:
//
//	KEY-LENGTH(uint8) | KEY(bytes) | TTL(uint64) | TIMESTAMP(uint64) | LASTACCESS(uint64) | VALUE-LENGTH(uint32) | VALUE(bytes)
//
// If the hash key already exists, the previous entry is deleted first. Returns
// ErrNotEnoughSpace if the buffer cannot accommodate the entry, or storage.ErrKeyTooLarge
// if the key exceeds MaxKeyLength.
func (t *Table) Put(hkey uint64, value storage.Entry) error { _ = "STUB: not implemented"; return nil }

// Check empty space on the allocated memory area.

// TTL + Timestamp + LastAccess + + value-Length + key-Length

// If we already have the key, delete it.

// Set key length. It's 1 byte.

// Set the key.

// Set the TTL. It's 8 bytes.

// Set the Timestamp. It's 8 bytes.

// Set the last access. It's 8 bytes.

// Set the value length. It's 4 bytes.

// Set the value.

// GetRaw returns the raw byte representation of the entry stored under the given hash key.
// The returned slice is a copy and includes the full binary-encoded entry (key, metadata and value).
// Returns ErrHKeyNotFound if the hash key does not exist.
func (t *Table) GetRaw(hkey uint64) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// In-memory structure:
// 1                 | klen       | 8           | 8                  | 8                  | 4                    | vlen
// KEY-LENGTH(uint8) | KEY(bytes) | TTL(uint64) | TIMESTAMP(uint64)  | LASTACCESS(uint64) | VALUE-LENGTH(uint64) | VALUE(bytes)

// One byte to keep key length
// key length
// TTL
// Timestamp
// LastAccess

// 4 bytes to keep value length
// value length

// Create a copy of the requested data.

// getRawKey reads and returns the raw key bytes from the memory buffer at the given offset.
func (t *Table) getRawKey(offset uint64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetRawKey returns the raw key bytes for the given hash key.
// Returns ErrHKeyNotFound if the hash key does not exist.
func (t *Table) GetRawKey(hkey uint64) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// GetKey returns the key as a string for the given hash key.
// Returns ErrHKeyNotFound if the hash key does not exist.
func (t *Table) GetKey(hkey uint64) (string, error) { _ = "STUB: not implemented"; return "", nil }

// GetTTL returns the TTL value in nanoseconds for the given hash key.
// Returns ErrHKeyNotFound if the hash key does not exist.
func (t *Table) GetTTL(hkey uint64) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// GetLastAccess returns the last access timestamp in nanoseconds for the given hash key.
// Returns ErrHKeyNotFound if the hash key does not exist.
func (t *Table) GetLastAccess(hkey uint64) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// Key length
// Key's itself
// TTL
// Timestamp

// get decodes a storage.Entry from the memory buffer at the given offset and updates
// the entry's last access time to the current time. It is used internally by Scan methods.
func (t *Table) get(offset uint64) storage.Entry {
	_ = "STUB: not implemented"

	// In-memory structure:
	//
	// KEY-LENGTH(uint8) | KEY(bytes) | TTL(uint64) | TIMESTAMP(uint64) | LASTACCESS(uint64) | VALUE-LENGTH(uint32) | VALUE(bytes)
	return *new(storage.Entry)
}

// Every SCAN call updates the last access time. We have to serialize the access to that field.

// Update the last access field

// Get retrieves the storage.Entry for the given hash key and updates the entry's
// last access time. Returns ErrHKeyNotFound if the hash key does not exist.
func (t *Table) Get(hkey uint64) (storage.Entry, error) {
	_ = "STUB: not implemented"
	return *new(storage.Entry), nil
}

// Delete removes the entry associated with the given hash key from the table.
// The occupied memory is marked as garbage but not reclaimed. Returns
// ErrHKeyNotFound if the hash key does not exist.
func (t *Table) Delete(hkey uint64) error { _ = "STUB: not implemented"; return nil }

// Try the previous tables.

// key, 1 byte for key size, klen for key's actual length.

// Delete the offset from offsetIndex

// TTL, skip it.

// Timestamp, skip it.

// LastAccess, skip it.

// value len and its header.

// Delete it from metadata

// UpdateTTL updates the TTL and timestamp fields of the entry identified by the
// given hash key in-place, and refreshes its last access time. Returns
// ErrHKeyNotFound if the hash key does not exist.
func (t *Table) UpdateTTL(hkey uint64, value storage.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// key, 1 byte for key size, klen for key's actual length.

// Set the new TTL. It's 8 bytes.

// Set the new Timestamp. It's 8 bytes.

// Update the last access field

// Check reports whether the given hash key exists in the table.
func (t *Table) Check(hkey uint64) bool { _ = "STUB: not implemented"; return false }

// Stats returns the current memory usage statistics for the table.
func (t *Table) Stats() Stats { _ = "STUB: not implemented"; return *new(Stats) }

// Range iterates over all entries in the table, calling f for each one.
// If f returns false, iteration stops. The iteration order is non-deterministic.
func (t *Table) Range(f func(hkey uint64, e storage.Entry) bool) { _ = "STUB: not implemented"; return }

// RangeHKey iterates over all hash keys in the table without decoding entries.
// If f returns false, iteration stops. The iteration order is non-deterministic.
func (t *Table) RangeHKey(f func(hkey uint64) bool) { _ = "STUB: not implemented"; return }

// Reset clears all entries and metadata, resets memory usage counters, and
// transitions the table to RecycledState. The underlying memory buffer is
// retained for reuse.
func (t *Table) Reset() { _ = "STUB: not implemented"; return }

// Scan performs a cursor-based iteration over the table entries. Starting from
// the given cursor position, it calls f for up to count entries. It returns
// the next cursor to resume scanning, or 0 when all entries have been visited.
// If f returns false, iteration stops early.
func (t *Table) Scan(cursor uint64, count int, f func(e storage.Entry) bool) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// end of the scan

// ScanRegexMatch performs a cursor-based iteration like Scan, but only yields
// entries whose keys match the given regular expression. Returns the next cursor
// to resume scanning, or 0 when all entries have been visited. Returns an error
// if the regular expression is invalid.
func (t *Table) ScanRegexMatch(cursor uint64, expr string, count int, f func(e storage.Entry) bool) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// end of the scan
