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

/*Package discovery provides a basic memberlist integration.*/
package discovery

import (
	"context"
	"errors"
	"net"
	"sync"

	"github.com/hashicorp/memberlist"
	"github.com/olric-data/olric/config"
	"github.com/olric-data/olric/internal/stats"
	"github.com/olric-data/olric/pkg/flog"
	"github.com/olric-data/olric/pkg/service_discovery"
)

const eventChanCapacity = 256

// UptimeSeconds is number of seconds since the server started.
var UptimeSeconds = stats.NewInt64Counter()

// ErrMemberNotFound indicates that the requested member could not be found in the member list.
var ErrMemberNotFound = errors.New("member not found")

// ClusterEvent is a single event related to node activity in the memberlist.
// The Node member of this struct must not be directly modified.
type ClusterEvent struct {
	Event    memberlist.NodeEventType
	NodeName string
	NodeAddr net.IP
	NodePort uint16
	NodeMeta []byte // Metadata from the delegate for this node.
}

func (c *ClusterEvent) MemberAddr() string { _ = "STUB: not implemented"; return "" }

// Discovery is a structure that encapsulates memberlist and
// provides useful functions to utilize it.
type Discovery struct {
	log        *flog.Logger
	member     *Member
	memberlist *memberlist.Memberlist
	config     *config.Config

	// To manage Join/Leave/Update events
	clusterEventsMtx sync.RWMutex
	ClusterEvents    chan *ClusterEvent

	// Try to reconnect dead members
	eventSubscribers []chan *ClusterEvent
	serviceDiscovery service_discovery.ServiceDiscovery

	// Flow control
	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

// New creates a new memberlist with a proper configuration and returns a new Discovery instance along with it.
func New(log *flog.Logger, c *config.Config) *Discovery { _ = "STUB: not implemented"; return nil }

func (d *Discovery) loadServiceDiscoveryPlugin() error { _ = "STUB: not implemented"; return nil }

// increaseUptimeSeconds calls UptimeSeconds.Increase function every second.
func (d *Discovery) increaseUptimeSeconds() { _ = "STUB: not implemented"; return }

func (d *Discovery) Start() error { _ = "STUB: not implemented"; return nil }

// ClusterEvents chan is consumed by the Olric package to maintain a consistent hash ring.

// Initialize a new memberlist

// Join is used to take an existing Memberlist and attempt to Join a cluster
// by contacting all the given hosts and performing a state sync. Initially,
// the Memberlist only contains our own state, so doing this will cause remote
// nodes to become aware of the existence of this node, effectively joining the cluster.
func (d *Discovery) Join() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (d *Discovery) Rejoin(peers []string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// GetMembers returns a full list of known alive nodes.
func (d *Discovery) GetMembers() []Member { _ = "STUB: not implemented"; return nil }

// sort members by birthdate

func (d *Discovery) NumMembers() int { _ = "STUB: not implemented"; return 0 }

// FindMemberByName finds and returns an alive member.
func (d *Discovery) FindMemberByName(name string) (Member, error) {
	_ = "STUB: not implemented"
	return *new(Member), nil
}

// FindMemberByID finds and returns an alive member.
func (d *Discovery) FindMemberByID(id uint64) (Member, error) {
	_ = "STUB: not implemented"
	return *new(Member), nil
}

// GetCoordinator returns the oldest node in the memberlist.
func (d *Discovery) GetCoordinator() Member { _ = "STUB: not implemented"; return *new(Member) }

// IsCoordinator returns true if the caller is the coordinator node.
func (d *Discovery) IsCoordinator() bool { _ = "STUB: not implemented"; return false }

// LocalNode is used to return the local Node
func (d *Discovery) LocalNode() *memberlist.Node { _ = "STUB: not implemented"; return nil }

// Shutdown will stop any background maintenance of network activity
// for this memberlist, causing it to appear "dead". A leave message
// will not be broadcasted prior, so the cluster being left will have
// to detect this node's Shutdown using probing. If you wish to more
// gracefully exit the cluster, call Leave prior to shutting down.
//
// This method is safe to call multiple times.
func (d *Discovery) Shutdown() error { _ = "STUB: not implemented"; return nil }

// We don't do that in a goroutine with a timeout mechanism
// because this mechanism may cause goroutine leak.

// Leave will broadcast a leave message but will not shutdown the background
// listeners, meaning the node will continue participating in gossip and state
// updates.
