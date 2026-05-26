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

package protocol

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/tidwall/redcon"
)

type ClusterRoutingTable struct{}

func NewClusterRoutingTable() *ClusterRoutingTable { _ = "STUB: not implemented"; return nil }

func (c *ClusterRoutingTable) Command(ctx context.Context) *redis.Cmd {
	_ = "STUB: not implemented"
	return nil
}

func ParseClusterRoutingTable(cmd redcon.Command) (*ClusterRoutingTable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ClusterMembers struct{}

func NewClusterMembers() *ClusterMembers { _ = "STUB: not implemented"; return nil }

func (c *ClusterMembers) Command(ctx context.Context) *redis.Cmd {
	_ = "STUB: not implemented"
	return nil
}

func ParseClusterMembers(cmd redcon.Command) (*ClusterMembers, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
