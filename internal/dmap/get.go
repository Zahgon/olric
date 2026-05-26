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

package dmap

import (
	"context"
	"errors"

	"github.com/olric-data/olric/internal/discovery"
	"github.com/olric-data/olric/internal/stats"
	"github.com/olric-data/olric/pkg/storage"
)

// Entry is a DMap entry with its metadata.
type Entry struct {
	Key       string
	Value     interface{}
	TTL       int64
	Timestamp int64
}

var (
	// GetMisses is the number of entries that have been requested and not found
	GetMisses = stats.NewInt64Counter()

	// GetHits is the number of entries that have been requested and found present
	GetHits = stats.NewInt64Counter()

	// EvictedTotal is the number of entries removed from cache to free memory for new entries.
	EvictedTotal = stats.NewInt64Counter()
)

// ErrReadQuorum means that read quorum cannot be reached to operate.
var ErrReadQuorum = errors.New("read quorum cannot be reached")

type version struct {
	host  *discovery.Member
	entry storage.Entry
}

// getOnFragment retrieves an entry from the associated fragment based on the provided environment details.
// It returns the found entry or an error if the key is not found, too large, or expired.
func (dm *DMap) getOnFragment(e *env) (storage.Entry, error) {
	_ = "STUB: not implemented"
	return *new(storage.Entry), nil
}

// lookupOnPreviousOwner retrieves the version of a key from a previous owner in the cluster.
// It communicates with the specified owner node and decodes the value into a version object.
func (dm *DMap) lookupOnPreviousOwner(owner *discovery.Member, key string) (*version, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dm *DMap) valueToVersion(value storage.Entry) *version { _ = "STUB: not implemented"; return nil }

// lookupOnThisNode searches for a key's version on the current node, considering
// only the primary partition owner.
func (dm *DMap) lookupOnThisNode(hkey uint64, key string) *version {
	_ = "STUB: not implemented"
	// Check on localhost, the partition owner.
	return nil
}

// still need to use "ver". just log this error.

// We found the key
//
// LRU and MaxIdleDuration eviction policies are only valid on
// the partition owner. Normally, we shouldn't need to retrieve the keys
// from the backup or the previous owners. When the fsck merge
// a fragmented partition or recover keys from a backup, Olric
// continue maintaining a reliable access log.

// lookupOnOwners collects versions of a key/value pair on the partition owner
// by including previous partition owners.
func (dm *DMap) lookupOnOwners(hkey uint64, key string) []*version {
	_ = "STUB: not implemented"
	return nil
}

// Run a query on the previous owners.
// Traverse in reverse order. Except from the latest host, this one.

// Ignore failed owners. The balancer will wipe out
// the data on those hosts.

func (dm *DMap) sortVersions(versions []*version) []*version { _ = "STUB: not implemented"; return nil }

// Explicit is better than implicit.

// sanitizeAndSortVersions removes nil versions from the input slice and sorts
// the remaining versions by recency.
func (dm *DMap) sanitizeAndSortVersions(versions []*version) []*version {
	_ = "STUB: not implemented"
	return nil

	// We use versions slice for read-repair. Clear nil values first.
}

// lookupOnReplicas retrieves data from replica nodes for the given hash key and
// key, returning a list of versioned entries.
func (dm *DMap) lookupOnReplicas(hkey uint64, key string) []*version {
	_ = "STUB: not implemented"
	// Check replicas
	return nil
}

// readRepair performs synchronization of inconsistent replicas by applying the
// winning version to out-of-sync nodes.
func (dm *DMap) readRepair(winner *version, versions []*version) { _ = "STUB: not implemented"; return }

// Check the timestamp first, we apply the "last write wins" rule here.

// Sync

// If readRepair is enabled, this function is called by every GET request.

// getOnCluster retrieves the storage.Entry for a given hashed key and key string
// from cluster nodes with read quorum. It ensures data consistency via read repair
// and returns ErrKeyNotFound or ErrReadQuorum if conditions aren't met.
func (dm *DMap) getOnCluster(hkey uint64, key string) (storage.Entry, error) {
	_ = "STUB: not implemented"
	// RUnlock should not be called with a defer statement here because
	//  the readRepair function may call putOnFragment function which needs a write
	// lock. Please remember calling RUnlock before returning here.
	return *new(storage.Entry), nil
}

// We checked everywhere, it's not here.

// The most up-to-date version of the values.

// Parallel read operations may propagate different versions of
// the same key/value pair. The rule is simple: last write wins.

// Get gets the value for the given key. It returns ErrKeyNotFound if the DB
// does not contain the key. It's thread-safe. It is safe to modify the contents
// of the returned value.
func (dm *DMap) Get(ctx context.Context, key string) (storage.Entry, error) {
	_ = "STUB: not implemented"
	return *new(storage.Entry), nil
}

// We are on the partition owner

// number of keys that have been requested and found present

// Redirect to the partition owner

// number of keys that have been requested and found present
