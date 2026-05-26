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
	"github.com/buraksezer/consistent"
	"github.com/olric-data/olric/internal/discovery"
)

func (r *RoutingTable) distributePrimaryCopies(partID uint64) []discovery.Member {
	_ = "STUB: not implemented"
	// First you need to create a copy of the owners list. Don't modify the current list.
	return nil
}

// Find the new partition owner.

// First run.

// Prune dead nodes

// Prune empty nodes

// Pass it. If the node is down, memberlist package will send a leave event.

// Pass it. If the node is down, memberlist package will send a leave event.

// Empty partition. Delete it from ownership list.

// Here add the new partition newOwner.

// Remove it from the current position

// Append it again to head

func (r *RoutingTable) getReplicaOwners(partID uint64) ([]consistent.Member, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Fail early

func isOwner(member discovery.Member, owners []consistent.Member) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *RoutingTable) distributeBackups(partID uint64) []discovery.Member {
	_ = "STUB: not implemented"
	return nil
}

// Remove the primary owner

// First run

// Prune dead nodes

// Delete it.

// Delete it.

// Prune empty nodes

// Pass it. If the node is down, memberlist package will send a leave event.

// Pass it. If the node is down, memberlist package will send a leave event.

// About this scenario:
//
// * ReplicaCount = 3
// * Create three nodes and insert some keys
// * Kill one of the nodes
// * Now we have replicas that it's impossible to transfer its ownership
// * Since we cannot drop a healthy replica, we prefer to keep it until
//   a new node joined. Then, we transfer the ownership safely.
// * During this incident, a node owns a primary and backup replicas at the same time.

// Empty node, delete it.

// Here add the new backup owners.

// Remove it from the current position

// Append it again to head
