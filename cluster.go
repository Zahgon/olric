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

	"github.com/tidwall/redcon"
)

type Route struct {
	PrimaryOwners []string
	ReplicaOwners []string
}

type RoutingTable map[uint64]Route

func mapToRoutingTable(slice []interface{}) (RoutingTable, error) {
	_ = "STUB: not implemented"
	return *new(RoutingTable), nil
}

func (db *Olric) clusterRoutingTableCommandHandler(conn redcon.Conn, cmd redcon.Command) {
	_ = "STUB: not implemented"
	return
}

// Redirect to the cluster coordinator

func (db *Olric) fillRoutingTable() RoutingTable {
	_ = "STUB: not implemented"
	return *new(RoutingTable)
}

func (db *Olric) routingTable(ctx context.Context) (RoutingTable, error) {
	_ = "STUB: not implemented"
	return *new(RoutingTable), nil
}

func (db *Olric) clusterMembersCommandHandler(conn redcon.Conn, cmd redcon.Command) {
	_ = "STUB: not implemented"
	return
}

// go-redis/redis package cannot handle uint64. At the time of this writing,
// there is no solution for this, and I don't want to use a soft fork to repair it.
//conn.WriteUint64(member.ID)
