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
	"time"

	"github.com/olric-data/olric/internal/bufpool"
	"github.com/olric-data/olric/internal/discovery"
	"github.com/olric-data/olric/internal/stats"
	"github.com/olric-data/olric/pkg/storage"
	"github.com/redis/go-redis/v9"
)

var pool = bufpool.New()

// EntriesTotal is the total number of entries(including replicas)
// stored during the life of this instance.
var EntriesTotal = stats.NewInt64Counter()

var (
	ErrKeyFound      = errors.New("key found")
	ErrWriteQuorum   = errors.New("write quorum cannot be reached")
	ErrKeyTooLarge   = errors.New("key too large")
	ErrEntryTooLarge = errors.New("entry too large for the configured table size")
)

func prepareTTL(e *env) int64 { _ = "STUB: not implemented"; return 0 }

// putOnFragment calls underlying storage engine's Put method to store the key/value pair. It's not thread-safe.
func (dm *DMap) putEntryOnFragment(e *env, nt storage.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// total number of entries stored during the life of this instance.

func (dm *DMap) prepareEntry(e *env) storage.Entry {
	_ = "STUB: not implemented"
	return *new(storage.Entry)
}

func (dm *DMap) putOnReplicaFragment(e *env) error { _ = "STUB: not implemented"; return nil }

// total number of entries stored during the life of this instance.

func (dm *DMap) asyncPutOnBackup(e *env, data []byte, owner discovery.Member) {
	_ = "STUB: not implemented"
	return
}

func (dm *DMap) asyncPutOnCluster(e *env, nt storage.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Fire and forget mode.

func (dm *DMap) syncPutOnCluster(e *env, nt storage.Entry) error {
	_ = "STUB: not implemented"
	// Quorum based replication.
	return nil
}

func (dm *DMap) setLRUEvictionStats(e *env) error {
	_ = "STUB: not implemented"
	// Try to make room for the new item, if it's required.
	// MaxKeys and MaxInuse properties of LRU can be used in the same time.
	// But I think that it's good to use only one of time in a production system.
	// Because it should be easy to understand and debug.
	return nil
}

// This works for every request if you enabled LRU.
// But loading a number from memory should be very cheap.
// ownedPartitionCount changes in the case of node join or leave.

// Routing table is an eventually consistent data structure. In order to prevent a panic in prod,
// check the owned partition count before doing math.

// MaxKeys controls maximum key count owned by this node.
// We need ownedPartitionCount property because every partition
// manages itself independently. So if you set MaxKeys=70 and
// your partition count is 7, every partition 10 keys at maximum.

// MaxInuse controls maximum in-use memory of partitions on this node.
// We need ownedPartitionCount property because every partition
// manages itself independently. So if you set MaxInuse=70M(in bytes) and
// your partition count is 7, every partition consumes 10M in-use space at maximum.
// WARNING: Actual allocated memory can be different.

func (dm *DMap) checkPutConditions(e *env) error {
	_ = "STUB: not implemented"
	// Only set the key if it does not already exist.
	return nil
}

// Only set the key if it already exists.

func (dm *DMap) putOnCluster(e *env) error { _ = "STUB: not implemented"; return nil }

// Fire and forget mode. Calls PutBackup command in different goroutines
// and stores the key/value pair on local storage instance.

// Quorum based replication.

// single replica

func (dm *DMap) writePutCommand(e *env) (*redis.StatusCmd, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// put controls every write operation in Olric. It redirects the requests to its owner,
// if the key belongs to another host.
func (dm *DMap) put(e *env) error { _ = "STUB: not implemented"; return nil }

// We are on the partition owner.

// Redirect to the partition owner.

type PutConfig struct {
	HasEX         bool
	EX            time.Duration
	HasPX         bool
	PX            time.Duration
	HasEXAT       bool
	EXAT          time.Duration
	HasPXAT       bool
	PXAT          time.Duration
	HasNX         bool
	HasXX         bool
	OnlyUpdateTTL bool
}

// Put sets the value for the given key. It overwrites any previous value
// for that key, and it's thread-safe. The key has to be a string. value type
// is arbitrary. It is safe to modify the contents of the arguments after
// Put returns but not before.
func (dm *DMap) Put(ctx context.Context, key string, value interface{}, cfg *PutConfig) error {
	_ = "STUB: not implemented"
	return nil
}
