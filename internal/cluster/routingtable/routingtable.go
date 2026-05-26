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

package routingtable

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/buraksezer/consistent"
	"github.com/olric-data/olric/config"
	"github.com/olric-data/olric/internal/cluster/partitions"
	"github.com/olric-data/olric/internal/discovery"
	"github.com/olric-data/olric/internal/environment"
	"github.com/olric-data/olric/internal/server"
	"github.com/olric-data/olric/internal/service"
	"github.com/olric-data/olric/pkg/flog"
)

// ErrClusterQuorum means that the cluster could not reach a healthy numbers of members to operate.
var ErrClusterQuorum = errors.New("cannot be reached cluster quorum to operate")

type route struct {
	Owners  []discovery.Member
	Backups []discovery.Member
}

type RoutingTable struct {
	sync.RWMutex // routingMtx

	// Currently owned partition count. Approximate LRU implementation
	// uses that.
	ownedPartitionCount uint64
	signature           uint64
	// numMembers is used to check cluster quorum.
	numMembers int32

	// These values is useful to control operation status.
	bootstrapped int32

	updateRoutingMtx sync.Mutex
	table            map[uint64]*route
	consistent       *consistent.Consistent
	this             discovery.Member
	members          *Members
	config           *config.Config
	log              *flog.Logger
	primary          *partitions.Partitions
	backup           *partitions.Partitions
	client           *server.Client
	server           *server.Server
	discovery        *discovery.Discovery
	callbacks        []func()
	callbackMtx      sync.Mutex
	pushPeriod       time.Duration
	// The command handlers of the routing table service should wait for the cluster join event.
	joined chan struct{}
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

func registerErrors() { _ = "STUB: not implemented"; return }

func New(e *environment.Environment) *RoutingTable {
	_ = "STUB: not implemented"
	// The routing table has to be started properly before accepting connections.
	return nil
}

// TODO: This also may be a configuration param.

func (r *RoutingTable) Discovery() *discovery.Discovery { _ = "STUB: not implemented"; return nil }

func (r *RoutingTable) This() discovery.Member {
	_ = "STUB: not implemented"

	// setNumMembers assigns the current number of members in the cluster to a variable.
	return *new(discovery.Member)
}

func (r *RoutingTable) setNumMembers() {
	_ = "STUB: not implemented"
	// Calling NumMembers in every request is quite expensive.
	// It's rarely updated. Just call this when the membership info changed.
	return
}

func (r *RoutingTable) SetNumMembersEagerly(nr int32) { _ = "STUB: not implemented"; return }

func (r *RoutingTable) NumMembers() int32 { _ = "STUB: not implemented"; return 0 }

func (r *RoutingTable) Members() *Members { _ = "STUB: not implemented"; return nil }

func (r *RoutingTable) setSignature(s uint64) { _ = "STUB: not implemented"; return }

func (r *RoutingTable) Signature() uint64 { _ = "STUB: not implemented"; return 0 }

func (r *RoutingTable) setOwnedPartitionCount() { _ = "STUB: not implemented"; return }

func (r *RoutingTable) OwnedPartitionCount() uint64 { _ = "STUB: not implemented"; return 0 }

func (r *RoutingTable) CheckMemberCountQuorum() error {
	_ = "STUB: not implemented"
	// This type of quorum function determines the presence of quorum based on the count of members in the cluster,
	// as observed by the local member’s cluster membership manager
	return nil
}

func (r *RoutingTable) markBootstrapped() {
	_ = "STUB: not implemented"
	// Bootstrapped by the coordinator.
	return
}

func (r *RoutingTable) IsBootstrapped() bool {
	_ = "STUB: not implemented"
	// Bootstrapped by the coordinator.
	return false
}

// CheckBootstrap is called for every request and checks whether the node is bootstrapped.
// It has to be very fast for a smooth operation.
func (r *RoutingTable) CheckBootstrap() error {
	_ = "STUB: not implemented"
	// Prevent creating expensive structures for every request,
	// Just check an integer value atomically.
	return nil
}

// Final error

func (r *RoutingTable) fillRoutingTable() { _ = "STUB: not implemented"; return }

func (r *RoutingTable) UpdateEagerly() { _ = "STUB: not implemented"; return }

func (r *RoutingTable) updateRouting() {
	_ = "STUB: not implemented"
	// This function is called by listenMemberlistEvents and updateRoutingPeriodically
	// So this lock prevents parallel execution.
	return
}

// This function is only run by the cluster coordinator.

// This type of quorum function determines the presence of quorum based on the count of members in the cluster,
// as observed by the local member’s cluster membership manager

func (r *RoutingTable) processClusterEvent(event *discovery.ClusterEvent) {
	_ = "STUB: not implemented"
	return
}

// Don't try to used closed sockets again.

// Node's birthdate may be changed. Close the pool and re-add to the hash ring.
// This takes linear time, but member count should be too small for a decent computer!

// Store the current number of members in the member list.
// We need this to implement a simple split-brain protection algorithm.

func (r *RoutingTable) listenClusterEvents(eventCh chan *discovery.ClusterEvent) {
	_ = "STUB: not implemented"
	return
}

func (r *RoutingTable) pushPeriodically() { _ = "STUB: not implemented"; return }

func (r *RoutingTable) Join() error { _ = "STUB: not implemented"; return nil }

func (r *RoutingTable) Start() error { _ = "STUB: not implemented"; return nil }

// It's time to start the routing table service. Otherwise, this method will return an error.

// Not yet, or the join process has failed

// Store the current number of members in the member list.
// We need this to implement a simple split-brain protection algorithm.

// 1 Hour

// Check member count quorum now. If there are not enough peers to work, wait forever.

func (r *RoutingTable) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// already closed

var _ service.Service = (*RoutingTable)(nil)
