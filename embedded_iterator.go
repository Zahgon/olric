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

package olric

import (
	"sync"

	"github.com/olric-data/olric/internal/dmap"
)

// EmbeddedIterator implements distributed query on DMaps.
type EmbeddedIterator struct {
	mtx sync.Mutex

	client          *EmbeddedClient
	dm              *dmap.DMap
	clusterIterator *ClusterIterator
}

func (e *EmbeddedIterator) scanOnOwners() error { _ = "STUB: not implemented"; return nil }

// Build a scan command here

// Fetch a Redis client for the given owner.

// Next returns true if there is more key in the iterator implementation.
// Otherwise, it returns false.
func (e *EmbeddedIterator) Next() bool { _ = "STUB: not implemented"; return false }

// Key returns a key name from the distributed map.
func (e *EmbeddedIterator) Key() string { _ = "STUB: not implemented"; return "" }

// Close stops the iteration and releases allocated resources.
func (e *EmbeddedIterator) Close() { _ = "STUB: not implemented"; return }
