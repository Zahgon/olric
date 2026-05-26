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
	"sync/atomic"
	"time"

	"github.com/olric-data/olric/config"
	"github.com/olric-data/olric/hasher"
	"github.com/olric-data/olric/internal/bufpool"
	"github.com/olric-data/olric/internal/dmap"
	"github.com/olric-data/olric/internal/protocol"
	"github.com/olric-data/olric/internal/server"
	"github.com/olric-data/olric/pkg/storage"
	"github.com/olric-data/olric/stats"
	"github.com/redis/go-redis/v9"
)

var pool = bufpool.New()

// DefaultRoutingTableFetchInterval is the default value of RoutingTableFetchInterval. ClusterClient implementation
// fetches the routing table from the cluster to route requests to the right partition.
const DefaultRoutingTableFetchInterval = time.Minute

type ClusterLockContext struct {
	key   string
	token string
	dm    *ClusterDMap
}

// ClusterDMap implements a client for DMaps.
type ClusterDMap struct {
	name          string
	newEntry      func() storage.Entry
	config        *dmapConfig
	client        *server.Client
	clusterClient *ClusterClient
}

// Name exposes name of the DMap.
func (dm *ClusterDMap) Name() string {
	_ = "STUB: not implemented"

	// processProtocolError processes protocol-related errors and translates them into defined application-level errors.
	return ""
}

func processProtocolError(err error) error { _ = "STUB: not implemented"; return nil }

// writePutCommand constructs and returns a new protocol.Put command based on the provided key, value, and configuration options.
func (dm *ClusterDMap) writePutCommand(c *dmap.PutConfig, key string, value []byte) *protocol.Put {
	_ = "STUB: not implemented"
	return nil
}

func (cl *ClusterClient) clientByPartID(partID uint64) (*redis.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cl *ClusterClient) smartPick(dmap, key string) (*redis.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Put sets the value for the given key. It overwrites any previous value for
// that key, and it's thread-safe. The key has to be a string. value type is arbitrary.
// It is safe to modify the contents of the arguments after Put returns but not before.
func (dm *ClusterDMap) Put(ctx context.Context, key string, value interface{}, options ...PutOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (dm *ClusterDMap) makeGetResponse(cmd *redis.StringCmd) (*GetResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get gets the value for the given key. It returns ErrKeyNotFound if the DB
// does not contain the key. It's thread-safe. It is safe to modify the contents
// of the returned value. See GetResponse for the details.
func (dm *ClusterDMap) Get(ctx context.Context, key string) (*GetResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete deletes values for the given keys. Delete will not return an error if the key doesn't exist.
// It's thread-safe. It is safe to modify the contents of the argument after Delete returns.
func (dm *ClusterDMap) Delete(ctx context.Context, keys ...string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Incr atomically increments the key by delta. The return value is the new value
// after being incremented or an error.
func (dm *ClusterDMap) Incr(ctx context.Context, key string, delta int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Decr atomically decrements the key by delta. The return value is the new value
// after being decremented or an error.
func (dm *ClusterDMap) Decr(ctx context.Context, key string, delta int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetPut atomically sets the key to value and returns the old value stored at a key. It returns nil if there is no
// previous value.
func (dm *ClusterDMap) GetPut(ctx context.Context, key string, value interface{}) (*GetResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// First try to set a key/value with GetPut

// IncrByFloat atomically increments the key by delta. The return value is the new value
// after being incremented or an error.
func (dm *ClusterDMap) IncrByFloat(ctx context.Context, key string, delta float64) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Expire updates the expiry for the given key. It returns ErrKeyNotFound if
// the DB does not contain the key. It's thread-safe.
func (dm *ClusterDMap) Expire(ctx context.Context, key string, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Lock sets a lock for the given key. Acquired lock is only for the key in
// this dmap.
//
// It returns immediately if it acquires the lock for the given key. Otherwise,
// it waits until deadline.
//
// You should know that the locks are approximate and only to be used for
// non-critical purposes.
func (dm *ClusterDMap) Lock(ctx context.Context, key string, deadline time.Duration) (LockContext, error) {
	_ = "STUB: not implemented"
	return *new(LockContext), nil
}

// LockWithTimeout sets a lock for the given key. If the lock is still unreleased
// the end of a given period of time, it automatically releases the lock.
// Acquired lock is only for the key in this DMap.
//
// It returns immediately if it acquires the lock for the given key. Otherwise,
// it waits until deadline.
//
// You should know that the locks are approximate and only to be used for
// non-critical purposes.
func (dm *ClusterDMap) LockWithTimeout(ctx context.Context, key string, timeout, deadline time.Duration) (LockContext, error) {
	_ = "STUB: not implemented"
	return *new(LockContext), nil
}

// Close stops background routines and frees allocated resources.
func (dm *ClusterDMap) Close(_ context.Context) error {
	_ = "STUB: not implemented"

	// Unlock releases the distributed lock associated with the current context by using the provided context for execution.
	return nil
}

func (c *ClusterLockContext) Unlock(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Lease extends the lease of the distributed lock associated with the context for the specified duration.
func (c *ClusterLockContext) Lease(ctx context.Context, duration time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Scan returns an iterator to loop over the keys.
//
// Available scan options:
//
// * Count
// * Match
func (dm *ClusterDMap) Scan(ctx context.Context, options ...ScanOption) (Iterator, error) {
	_ = "STUB: not implemented"
	return *new(Iterator), nil
}

// Embedded iterator uses a slightly different scan function.

// Load the route for the first partition (0) to scan.

// Destroy flushes the given DMap on the cluster. You should know that there
// is no global lock on DMaps. So if you call Put/PutEx and Destroy methods
// concurrently on the cluster, Put call may set new values to the DMap.
func (dm *ClusterDMap) Destroy(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// ClusterClient is a client for managing and interacting with a distributed cluster of nodes.
type ClusterClient struct {
	client         *server.Client
	config         *clusterClientConfig
	logger         *log.Logger
	routingTable   atomic.Value
	partitionCount uint64
	wg             sync.WaitGroup
	ctx            context.Context
	cancel         context.CancelFunc
}

// Ping sends a ping message to an Olric node. Returns PONG if a message is empty,
// otherwise return a copy of the message as bulk. This command is often used to test
// if a connection is still alive or to measure latency.
func (cl *ClusterClient) Ping(ctx context.Context, addr, message string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// RoutingTable returns the latest version of the routing table.
func (cl *ClusterClient) RoutingTable(ctx context.Context) (RoutingTable, error) {
	_ = "STUB: not implemented"
	return *new(RoutingTable), nil
}

// Stats returns stats.Stats with the given options.
func (cl *ClusterClient) Stats(ctx context.Context, address string, options ...StatsOption) (stats.Stats, error) {
	_ = "STUB: not implemented"
	return *new(stats.Stats), nil
}

// Members returns a thread-safe list of cluster members.
func (cl *ClusterClient) Members(ctx context.Context) ([]Member, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// go-redis/redis package cannot handle uint64 type. At the time of this writing,
// there is no solution for this, and I don't want to use a soft fork to repair it.

// RefreshMetadata fetches a list of available members and the latest routing
// table version. It also closes stale clients if there are any.
func (cl *ClusterClient) RefreshMetadata(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Fetch a list of currently available cluster members.
	return nil
}

// Use a map for fast access.

// Clean stale client connections

// Gone

// Re-fetch the routing table, we should use the latest routing table version.

// Close stops background routines and frees allocated resources.
func (cl *ClusterClient) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Wait for the background workers:
// * fetchRoutingTablePeriodically

// Close the underlying TCP sockets gracefully.

// NewPubSub returns a new PubSub client with the given options.
func (cl *ClusterClient) NewPubSub(options ...PubSubOption) (*PubSub, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewDMap returns a new DMap client with the given options.
func (cl *ClusterClient) NewDMap(name string, options ...DMapOption) (DMap, error) {
	_ = "STUB: not implemented"
	return *new(DMap), nil
}

// ClusterClientOption is a functional option for configuring a clusterClientConfig instance.
type ClusterClientOption func(c *clusterClientConfig)

// clusterClientConfig holds the configuration required to initialize and manage a cluster client instance.
type clusterClientConfig struct {
	logger                    *log.Logger
	config                    *config.Client
	authentication            *config.Authentication
	hasher                    hasher.Hasher
	routingTableFetchInterval time.Duration
}

// WithHasher sets a custom hasher implementation to the cluster client configuration.
func WithHasher(h hasher.Hasher) ClusterClientOption {
	_ = "STUB: not implemented"
	return *new(ClusterClientOption)
}

// WithLogger sets a custom logger for the cluster client configuration.
func WithLogger(l *log.Logger) ClusterClientOption {
	_ = "STUB: not implemented"
	return *new(ClusterClientOption)
}

// WithConfig applies a specified config.Client to the clusterClientConfig.
func WithConfig(c *config.Client) ClusterClientOption {
	_ = "STUB: not implemented"
	return *new(ClusterClientOption)
}

// WithPassword configures a cluster client with the specified password for authentication.
func WithPassword(password string) ClusterClientOption {
	_ = "STUB: not implemented"
	return *new(ClusterClientOption)
}

// WithRoutingTableFetchInterval sets the interval for periodic fetching of the routing table in a cluster client configuration.
func WithRoutingTableFetchInterval(interval time.Duration) ClusterClientOption {
	_ = "STUB: not implemented"
	return *new(ClusterClientOption)
}

// fetchRoutingTable updates the cluster routing table by fetching the latest version from the cluster.
// It initializes the partition count if it's the first invocation. Returns an error if fetching fails.
func (cl *ClusterClient) fetchRoutingTable() error { _ = "STUB: not implemented"; return nil }

// First run. Partition count is a constant, actually. It has to be greater than zero.

// fetchRoutingTablePeriodically periodically updates the routing table by invoking fetchRoutingTable at configured intervals.
// It stops gracefully when the context is canceled or an error occurs.
func (cl *ClusterClient) fetchRoutingTablePeriodically() { _ = "STUB: not implemented"; return }

// NewClusterClient creates a new Client instance. It needs one node address at least to discover the whole cluster.
func NewClusterClient(addresses []string, options ...ClusterClientOption) (*ClusterClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initialize clients for the given cluster members.

// Discover all cluster members

// Hash function is required to target primary owners instead of random cluster members.

// Initial fetch. ClusterClient targets the primary owners for a smooth and quick operation.

// Refresh the routing table in every 15 seconds.

var (
	_ Client = (*ClusterClient)(nil)
	_ DMap   = (*ClusterDMap)(nil)
)
