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

/*
Package ramblock implements a GC-friendly in-memory storage engine by using
built-in maps and byte slices. It also supports compaction.
*/
package ramblock

import (
	"log"
	"time"

	"github.com/olric-data/olric/internal/ramblock/table"
	"github.com/olric-data/olric/pkg/storage"
)

const (
	maxGarbageRatio = 0.40
	// 1MB
	defaultTableSize = uint64(1 << 20)

	defaultMaxIdleTableTimeout = 15 * time.Minute
)

// RamBlock implements an in-memory storage engine.
type RamBlock struct {
	coefficient         uint64
	tableSize           uint64
	tablesByCoefficient map[uint64]*table.Table
	tables              []*table.Table
	config              *storage.Config
}

func DefaultConfig() *storage.Config { _ = "STUB: not implemented"; return nil }

func New(c *storage.Config) (*RamBlock, error) { _ = "STUB: not implemented"; return nil, nil }

func (rb *RamBlock) SetConfig(c *storage.Config) { _ = "STUB: not implemented"; return }

func (rb *RamBlock) makeTable() error { _ = "STUB: not implemented"; return nil }

func (rb *RamBlock) SetLogger(_ *log.Logger) { _ = "STUB: not implemented"; return }

func (rb *RamBlock) Start() error { _ = "STUB: not implemented"; return nil }

func requiredSizeForAnEntry(e storage.Entry) uint64 { _ = "STUB: not implemented"; return 0 }

func prepareTableSize(raw interface{}) (size uint64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Fork creates a new RamBlock instance.
func (rb *RamBlock) Fork(c *storage.Config) (storage.Engine, error) {
	_ = "STUB: not implemented"
	return *new(storage.Engine), nil
}

func (rb *RamBlock) Name() string { _ = "STUB: not implemented"; return "" }

func (rb *RamBlock) NewEntry() storage.Entry {
	_ = "STUB: not implemented"

	// putWithRetry ensures at least one table exists and retries the given write
	// function on a new table when the current one runs out of space.
	return *new(storage.Entry)
}

func (rb *RamBlock) putWithRetry(writeFn func(t *table.Table) error) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the last value, storage only calls Put on the last created table.

// try again

// everything is ok

// PutRaw sets the raw value for the given key.
func (rb *RamBlock) PutRaw(hkey uint64, value []byte) error { _ = "STUB: not implemented"; return nil }

// Put sets the value for the given key. It overwrites any previous value for that key
func (rb *RamBlock) Put(hkey uint64, value storage.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// GetRaw extracts encoded value for the given hkey. This is useful for merging tables.
func (rb *RamBlock) GetRaw(hkey uint64) ([]byte, error) {
	_ = "STUB: not implemented"
	// Scan available tables by starting the last added table.
	return nil, nil
}

// Try out the other tables.

// Found the key, return the stored value with its metadata.

// Nothing here.

// Get gets the value for the given key. It returns storage.ErrKeyNotFound if the DB
// does not contain the key. The returned Entry is its own copy,
// it is safe to modify the contents of the returned slice.
func (rb *RamBlock) Get(hkey uint64) (storage.Entry, error) {
	_ = "STUB: not implemented"
	// Scan available tables by starting the last added table.
	return *new(storage.Entry), nil
}

// Try out the other tables.

// Found the key, return the stored value with its metadata.

// Nothing here.

// GetTTL gets the timeout for the given key. It returns storage.ErrKeyNotFound if the DB
// does not contain the key.
func (rb *RamBlock) GetTTL(hkey uint64) (int64, error) {
	_ = "STUB: not implemented"
	// Scan available tables by starting the last added table.
	return 0, nil
}

// Try out the other tables.

// Found the key, return its ttl

// Nothing here.

func (rb *RamBlock) GetLastAccess(hkey uint64) (int64, error) {
	_ = "STUB: not implemented"
	// Scan available tables by starting the last added table.
	return 0, nil
}

// Try out the other tables.

// Found the key, return its ttl

// Nothing here.

// GetKey gets the key for the given hkey. It returns storage.ErrKeyNotFound if the DB
// does not contain the key.
func (rb *RamBlock) GetKey(hkey uint64) (string, error) {
	_ = "STUB: not implemented"
	// Scan available tables by starting the last added table.
	return "", nil
}

// Try out the other tables.

// Found the key, return its ttl

// Nothing here.

// Delete deletes the value for the given key. Delete will not returns error if key doesn't exist.
func (rb *RamBlock) Delete(hkey uint64) error {
	_ = "STUB: not implemented"
	// Scan available tables by starting the last added table.
	return nil
}

// Try out the other tables.

// UpdateTTL updates the expiry for the given key.
func (rb *RamBlock) UpdateTTL(hkey uint64, data storage.Entry) error {
	_ = "STUB: not implemented"
	// Scan available tables by starting the last added table.
	return nil
}

// Try out the other tables.

// Found the key, return the stored value with its metadata.

// Nothing here.

// Stats is a function which provides memory allocation and garbage ratio of a storage instance.
func (rb *RamBlock) Stats() storage.Stats { _ = "STUB: not implemented"; return *new(storage.Stats) }

// Check checks the key existence.
func (rb *RamBlock) Check(hkey uint64) bool {
	_ = "STUB: not implemented"
	// Scan available tables by starting the last added table.
	return false
}

// Nothing there.

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration. Range may be O(N) with
// the number of elements in the map even if f returns false after a constant
// number of calls.
func (rb *RamBlock) Range(f func(hkey uint64, e storage.Entry) bool) {
	_ = "STUB: not implemented"
	// Scan available tables by starting the last added table.
	return
}

// RangeHKey calls f sequentially for each key present in the map.
// If f returns false, range stops the iteration. Range may be O(N) with
// the number of elements in the map even if f returns false after a constant
// number of calls.
func (rb *RamBlock) RangeHKey(f func(hkey uint64) bool) {
	_ = "STUB: not implemented"
	// Scan available tables by starting the last added table.
	return
}

func (rb *RamBlock) findCoefficient(coefficient uint64) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (rb *RamBlock) scanCommon(cursor uint64, expr string, count int, f func(e storage.Entry) bool) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Invalid cursor

// Invalid cursor

// findCoefficient already returns the next valid coefficient

// The next table

func (rb *RamBlock) Scan(cursor uint64, count int, f func(e storage.Entry) bool) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (rb *RamBlock) ScanRegexMatch(cursor uint64, expr string, count int, f func(e storage.Entry) bool) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (rb *RamBlock) Close() error { _ = "STUB: not implemented"; return nil }

func (rb *RamBlock) Destroy() error { _ = "STUB: not implemented"; return nil }

var _ storage.Engine = (*RamBlock)(nil)
