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
	"sync"
	"time"

	"github.com/olric-data/olric/internal/discovery"
	"github.com/olric-data/olric/internal/dmap"
	"github.com/olric-data/olric/stats"
)

// EmbeddedLockContext is returned by Lock and LockWithTimeout methods.
// It should be stored in a proper way to release the lock.
type EmbeddedLockContext struct {
	key   string
	token []byte
	dm    *EmbeddedDMap
}

// Unlock releases the lock.
func (l *EmbeddedLockContext) Unlock(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Lease takes the duration to update the expiry for the given Lock.
func (l *EmbeddedLockContext) Lease(ctx context.Context, duration time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// EmbeddedClient is an Olric client implementation for embedded-member scenario.
type EmbeddedClient struct {
	db *Olric
}

// EmbeddedDMap is an DMap client implementation for embedded-member scenario.
type EmbeddedDMap struct {
	mtx           sync.RWMutex
	clusterClient *ClusterClient
	config        *dmapConfig
	member        discovery.Member
	dm            *dmap.DMap
	client        *EmbeddedClient
	name          string
}

func (dm *EmbeddedDMap) setOrGetClusterClient() (Client, error) {
	_ = "STUB: not implemented"
	// Acquire the read lock and try to access the cluster client, if any.
	return *new(Client), nil
}

// The cluster client is unset, try to create a new one.

// Check the existing value last time. There can be another running instances
// of this function.

// Create a new cluster client here.

// Pipeline is a mechanism to realise Redis Pipeline technique.
//
// Pipelining is a technique to extremely speed up processing by packing
// operations to batches, send them at once to Redis and read a replies in a
// singe step.
// See https://redis.io/topics/pipelining
//
// Pay attention, that Pipeline is not a transaction, so you can get unexpected
// results in case of big pipelines and small read/write timeouts.
// Redis client has retransmission logic in case of timeouts, pipeline
// can be retransmitted and commands can be executed more than once.
func (dm *EmbeddedDMap) Pipeline(opts ...PipelineOption) (*DMapPipeline, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RefreshMetadata fetches a list of available members and the latest routing
// table version. It also closes stale clients, if there are any. EmbeddedClient has
// this method to implement the Client interface. It doesn't need to refresh metadata manually.
func (e *EmbeddedClient) RefreshMetadata(_ context.Context) error {
	_ = "STUB: not implemented"
	// EmbeddedClient already has the latest metadata.
	return nil
}

// Scan returns an iterator to loop over the keys.
//
// Available scan options:
//
// * Count
// * Match
func (dm *EmbeddedDMap) Scan(ctx context.Context, options ...ScanOption) (Iterator, error) {
	_ = "STUB: not implemented"
	return *new(Iterator), nil
}

// Lock sets a lock for the given key. Acquired lock is only for the key in
// this dmap.
//
// It returns immediately if it acquires the lock for the given key. Otherwise,
// it waits until deadline.
//
// You should know that the locks are approximate, and only to be used for
// non-critical purposes.
func (dm *EmbeddedDMap) Lock(ctx context.Context, key string, deadline time.Duration) (LockContext, error) {
	_ = "STUB: not implemented"
	return *new(LockContext), nil
}

// LockWithTimeout sets a lock for the given key. If the lock is still unreleased
// the end of given period of time,
// it automatically releases the lock. Acquired lock is only for the key in
// this dmap.
//
// It returns immediately if it acquires the lock for the given key. Otherwise,
// it waits until deadline.
//
// You should know that the locks are approximate, and only to be used for
// non-critical purposes.
func (dm *EmbeddedDMap) LockWithTimeout(ctx context.Context, key string, timeout, deadline time.Duration) (LockContext, error) {
	_ = "STUB: not implemented"
	return *new(LockContext), nil
}

// Destroy flushes the given DMap on the cluster. You should know that there
// is no global lock on DMaps. So if you call Put/PutEx and Destroy methods
// concurrently on the cluster, Put call may set new values to the DMap.
func (dm *EmbeddedDMap) Destroy(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Expire updates the expiry for the given key. It returns ErrKeyNotFound if
// the DB does not contain the key. It's thread-safe.
func (dm *EmbeddedDMap) Expire(ctx context.Context, key string, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Name exposes name of the DMap.
func (dm *EmbeddedDMap) Name() string {
	_ = "STUB: not implemented"

	// GetPut atomically sets the key to value and returns the old value stored at key. It returns nil if there is no
	// previous value.
	return ""
}

func (dm *EmbeddedDMap) GetPut(ctx context.Context, key string, value interface{}) (*GetResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Decr atomically decrements the key by delta. The return value is the new value
// after being decremented or an error.
func (dm *EmbeddedDMap) Decr(ctx context.Context, key string, delta int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Incr atomically increments the key by delta. The return value is the new value
// after being incremented or an error.
func (dm *EmbeddedDMap) Incr(ctx context.Context, key string, delta int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// IncrByFloat atomically increments the key by delta. The return value is the new value after being incremented or an error.
func (dm *EmbeddedDMap) IncrByFloat(ctx context.Context, key string, delta float64) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Delete deletes values for the given keys. Delete will not return error
// if key doesn't exist. It's thread-safe. It is safe to modify the contents
// of the argument after Delete returns.
func (dm *EmbeddedDMap) Delete(ctx context.Context, keys ...string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Get gets the value for the given key. It returns ErrKeyNotFound if the DB
// does not contain the key. It's thread-safe. It is safe to modify the contents
// of the returned value. See GetResponse for the details.
func (dm *EmbeddedDMap) Get(ctx context.Context, key string) (*GetResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Put sets the value for the given key. It overwrites any previous value for
// that key, and it's thread-safe. The key has to be a string. value type is arbitrary.
// It is safe to modify the contents of the arguments after Put returns but not before.
func (dm *EmbeddedDMap) Put(ctx context.Context, key string, value interface{}, options ...PutOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Close stops background routines and frees allocated resources.
func (dm *EmbeddedDMap) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *EmbeddedClient) NewDMap(name string, options ...DMapOption) (DMap, error) {
	_ = "STUB: not implemented"
	return *new(DMap), nil
}

// Stats exposes some useful metrics to monitor an Olric node.
func (e *EmbeddedClient) Stats(ctx context.Context, address string, options ...StatsOption) (stats.Stats, error) {
	_ = "STUB: not implemented"
	return *new(stats.Stats), nil
}

// this node is not bootstrapped yet.

// Close stops background routines and frees allocated resources.
func (e *EmbeddedClient) Close(_ context.Context) error {
	_ = "STUB: not implemented"

	// Ping sends a ping message to an Olric node. Returns PONG if message is empty,
	// otherwise return a copy of the message as a bulk. This command is often used to test
	// if a connection is still alive, or to measure latency.
	return nil
}

func (e *EmbeddedClient) Ping(ctx context.Context, addr, message string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// RoutingTable returns the latest version of the routing table.
func (e *EmbeddedClient) RoutingTable(ctx context.Context) (RoutingTable, error) {
	_ = "STUB: not implemented"
	return *new(RoutingTable), nil
}

// Members returns a thread-safe list of cluster members.
func (e *EmbeddedClient) Members(_ context.Context) ([]Member, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewPubSub returns a new PubSub client with the given options.
func (e *EmbeddedClient) NewPubSub(options ...PubSubOption) (*PubSub, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewEmbeddedClient creates and returns a new EmbeddedClient instance.
func (db *Olric) NewEmbeddedClient() *EmbeddedClient { _ = "STUB: not implemented"; return nil }

var (
	_ Client = (*EmbeddedClient)(nil)
	_ DMap   = (*EmbeddedDMap)(nil)
)
