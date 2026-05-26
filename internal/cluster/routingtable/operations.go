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
	"github.com/tidwall/redcon"
)

func (r *RoutingTable) lengthOfPartCommandHandler(conn redcon.Conn, cmd redcon.Command) {
	_ = "STUB: not implemented"
	// The command handlers of the routing table service should wait for the cluster join event.
	return
}

func (r *RoutingTable) verifyRoutingTable(id uint64, table map[uint64]*route) error {
	_ = "STUB: not implemented"
	// Check the coordinator
	return nil
}

// Compare partition counts to catch a possible inconsistencies in configuration

func (r *RoutingTable) updateRoutingCommandHandler(conn redcon.Conn, cmd redcon.Command) {
	_ = "STUB: not implemented"
	// The command handlers of the routing table service should wait for the cluster join event.
	return
}

// Log this event

// owners(atomic.value) is guarded by routingUpdateMtx against parallel writers.
// Calculate routing signature. This is useful to control balancing tasks.

// Set partition(primary copies) owners

// Set backup owners

// Used by the LRU implementation.

// Bootstrapped by the coordinator.

// Collect report

// Call balancer to distribute load evenly
