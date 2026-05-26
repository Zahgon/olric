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
Package olric provides a distributed cache and in-memory key/value data store.
It can be used both as an embedded Go library and as a language-independent
service.

With Olric, you can instantly create a fast, scalable, shared pool of RAM across
a cluster of computers.

Olric is designed to be a distributed cache. But it also provides Publish/Subscribe,
data replication, failure detection and simple anti-entropy services.
So it can be used as an ordinary key/value data store to scale your cloud
application.
*/
package olric

import (
	"context"
	"sync"

	"github.com/olric-data/olric/config"
	"github.com/olric-data/olric/hasher"
	"github.com/olric-data/olric/internal/cluster/balancer"
	"github.com/olric-data/olric/internal/cluster/partitions"
	"github.com/olric-data/olric/internal/cluster/routingtable"
	"github.com/olric-data/olric/internal/dmap"
	"github.com/olric-data/olric/internal/environment"
	"github.com/olric-data/olric/internal/pubsub"
	"github.com/olric-data/olric/internal/server"
	"github.com/olric-data/olric/pkg/flog"
	"github.com/pkg/errors"
	"github.com/tidwall/redcon"
)

// ReleaseVersion is the current stable version of Olric
const ReleaseVersion string = "0.7.3"

var (
	// ErrOperationTimeout is returned when an operation times out.
	ErrOperationTimeout = errors.New("operation timeout")

	// ErrServerGone means that a cluster member is closed unexpectedly.
	ErrServerGone = errors.New("server is gone")

	// ErrKeyNotFound means that returned when a key could not be found.
	ErrKeyNotFound = errors.New("key not found")

	// ErrKeyFound means that the requested key found in the cluster.
	ErrKeyFound = errors.New("key found")

	// ErrWriteQuorum means that write quorum cannot be reached to operate.
	ErrWriteQuorum = errors.New("write quorum cannot be reached")

	// ErrReadQuorum means that read quorum cannot be reached to operate.
	ErrReadQuorum = errors.New("read quorum cannot be reached")

	// ErrLockNotAcquired is returned when the requested lock could not be acquired
	ErrLockNotAcquired = errors.New("lock not acquired")

	// ErrNoSuchLock is returned when the requested lock does not exist
	ErrNoSuchLock = errors.New("no such lock")

	// ErrClusterQuorum means that the cluster could not reach a healthy numbers of members to operate.
	ErrClusterQuorum = errors.New("failed to find enough peers to create quorum")

	// ErrKeyTooLarge means that the given key is too large to process.
	// The maximum length of a key is 256 bytes.
	ErrKeyTooLarge = errors.New("key too large")

	// ErrEntryTooLarge returned if the required space for an entry is bigger than table size.
	ErrEntryTooLarge = errors.New("entry too large for the configured table size")

	// ErrConnRefused returned if the target node refused a connection request.
	// It is good to call RefreshMetadata to update the underlying data structures.
	ErrConnRefused = errors.New("connection refused")

	// ErrWrongPass indicates that the provided password is incorrect during authentication.
	ErrWrongPass = errors.New("wrong password")
)

// Olric implements a distributed cache and in-memory key/value data store.
// It can be used both as an embedded Go library and as a language-independent
// service.
type Olric struct {
	// name is BindAddr:BindPort. It defines servers unique name in the cluster.
	name     string
	env      *environment.Environment
	config   *config.Config
	log      *flog.Logger
	hashFunc hasher.Hasher

	// Logical units to store data
	primary *partitions.Partitions
	backup  *partitions.Partitions

	// RESP server and clients.
	server *server.Server
	client *server.Client

	rt       *routingtable.RoutingTable
	balancer *balancer.Balancer

	pubsub *pubsub.Service
	dmap   *dmap.Service

	// Structures for flow control
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup

	// Callback function. Olric calls this after
	// the server is ready to accept new connections.
	started func()
}

func prepareConfig(c *config.Config) (*config.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func initializeServices(db *Olric) error { _ = "STUB: not implemented"; return nil }

// Add Services

// New creates a new Olric instance, otherwise returns an error.
func New(c *config.Config) (*Olric, error) { _ = "STUB: not implemented"; return nil, nil }

// Set the hash function. Olric distributes keys over partitions by hashing.

// Create a Redcon server instance

func (db *Olric) preconditionFunc(conn redcon.Conn, _ redcon.Command) bool {
	_ = "STUB: not implemented"
	return false
}

func (db *Olric) registerCommandHandlers() { _ = "STUB: not implemented"; return }

// callStartedCallback checks passed checkpoint count and calls the callback
// function.
func (db *Olric) callStartedCallback() { _ = "STUB: not implemented"; return }

func convertClusterError(err error) error { _ = "STUB: not implemented"; return nil }

// isOperable controls bootstrapping status and cluster quorum to prevent split-brain syndrome.
func (db *Olric) isOperable() error { _ = "STUB: not implemented"; return nil }

// An Olric node has to be bootstrapped to function properly.

// Start starts background servers and joins the cluster. You still must call Shutdown
// method if Start function returns an early error.
func (db *Olric) Start() error { _ = "STUB: not implemented"; return nil }

// This error group is responsible to run the TCP server at background and report errors.

// TCP server has been started

// TCP server could not be started due to an error. There is no need to run
// Olric.Shutdown here because we could not start anything.

// Balancer works periodically to balance partition data across the cluster.

// First, we need to join the cluster. Then, the routing table has been started.

// Start routing table service and member discovery subsystem.

// Start publish-subscribe service

// Start distributed map service

// Warn the user about his/her choice of configuration

// Wait for the TCP server.

// Shutdown stops background servers and leaves the cluster.
func (db *Olric) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Shutdown only once.

// Shutdown Redcon server

// db.name will be shown as empty string, if the program is killed before
// bootstrapping.

func convertDMapError(err error) error { _ = "STUB: not implemented"; return nil }

// registerErrors registers application-specific errors with their corresponding prefixes in the error management system.
func registerErrors() { _ = "STUB: not implemented"; return }
