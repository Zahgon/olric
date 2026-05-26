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
	"context"
	"log"
	"sync"

	"github.com/olric-data/olric/internal/dmap"
)

type currentCursor struct {
	primary uint64
	replica uint64
}

// ClusterIterator implements distributed query on DMaps.
type ClusterIterator struct {
	mtx             sync.Mutex // protects pos and page
	routingTableMtx sync.Mutex // protects routingTable and partitionCount

	logger         *log.Logger
	dm             *ClusterDMap
	clusterClient  *ClusterClient
	pos            int
	page           []string
	route          *Route
	partitionKeys  map[string]struct{}
	cursors        map[uint64]map[string]*currentCursor
	partID         uint64 // current partition id
	routingTable   RoutingTable
	partitionCount uint64
	config         *dmap.ScanConfig
	scanner        func() error
	wg             sync.WaitGroup
	ctx            context.Context
	cancel         context.CancelFunc
}

func (i *ClusterIterator) loadRoute() { _ = "STUB: not implemented"; return }

func (i *ClusterIterator) updateCursor(owner string, cursor uint64) {
	_ = "STUB: not implemented"
	return
}

func (i *ClusterIterator) loadCursor(owner string) uint64 { _ = "STUB: not implemented"; return 0 }

func (i *ClusterIterator) updateIterator(keys []string, cursor uint64, owner string) {
	_ = "STUB: not implemented"
	return
}

func (i *ClusterIterator) getOwners() []string { _ = "STUB: not implemented"; return nil }

// Make a safe copy of the raw.

func (i *ClusterIterator) removeScannedOwner(idx int) { _ = "STUB: not implemented"; return }

func (i *ClusterIterator) scanOnOwners() error { _ = "STUB: not implemented"; return nil }

// Build a scan command here

// Fetch a Redis client for the given owner.

func (i *ClusterIterator) resetPage() { _ = "STUB: not implemented"; return }

func (i *ClusterIterator) fetchData() error { _ = "STUB: not implemented"; return nil }

func (i *ClusterIterator) reset() { _ = "STUB: not implemented"; return }

func (i *ClusterIterator) next() bool { _ = "STUB: not implemented"; return false }

// We have data on the page to read. Stop the iteration.

// We completed scanning all the owners. Stop the iteration.

// Next returns true if there is more key in the iterator implementation.
// Otherwise, it returns false
func (i *ClusterIterator) Next() bool { _ = "STUB: not implemented"; return false }

// Key returns a key name from the distributed map.
func (i *ClusterIterator) Key() string { _ = "STUB: not implemented"; return "" }

func (i *ClusterIterator) fetchRoutingTablePeriodically() { _ = "STUB: not implemented"; return }

func (i *ClusterIterator) fetchRoutingTable() error { _ = "STUB: not implemented"; return nil }

// Partition count is a constant, actually. It has to be greater than zero.

// Close stops the iteration and releases allocated resources.
func (i *ClusterIterator) Close() { _ = "STUB: not implemented"; return }

// await for routing table updater
