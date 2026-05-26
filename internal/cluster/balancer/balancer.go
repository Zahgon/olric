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

package balancer

import (
	"context"
	"sync"

	"github.com/olric-data/olric/config"
	"github.com/olric-data/olric/internal/cluster/partitions"
	"github.com/olric-data/olric/internal/cluster/routingtable"
	"github.com/olric-data/olric/internal/discovery"
	"github.com/olric-data/olric/internal/environment"
	"github.com/olric-data/olric/internal/service"
	"github.com/olric-data/olric/pkg/flog"
)

type Balancer struct {
	sync.Mutex

	log     *flog.Logger
	config  *config.Config
	primary *partitions.Partitions
	backup  *partitions.Partitions
	rt      *routingtable.RoutingTable
	wg      sync.WaitGroup
	ctx     context.Context
	cancel  context.CancelFunc
}

func New(e *environment.Environment) *Balancer { _ = "STUB: not implemented"; return nil }

func (b *Balancer) isAlive() bool { _ = "STUB: not implemented"; return false }

// The node is gone.

func (b *Balancer) scanPartition(sign uint64, part *partitions.Partition, owners ...discovery.Member) {
	_ = "STUB: not implemented"
	return
}

// if this returns true, the iteration continues

func (b *Balancer) primaryCopies() { _ = "STUB: not implemented"; return }

// Empty partition. Skip it.

// Here we don't use CompareByID function because the routing table is an
// eventually consistent data structure and a node can try to move data
// to previous instance(the same name but a different birthdate)
// of itself. So just check the name.

// Already belongs to me.

// This is a previous owner. Move the keys.

func (b *Balancer) breakLoop(sign uint64) bool { _ = "STUB: not implemented"; return false }

// Routing table is updated. Just quit. Another balancer goroutine
// will work on the new table immediately.

func (b *Balancer) backupCopies() { _ = "STUB: not implemented"; return }

// Here we don't use CompareById function because the routing table
// is an eventually consistent data structure and a node can try to
// move data to previous instance(the same name but a different birthdate)
// of itself. So just check the name.

// Already belongs to me.

func (b *Balancer) triggerBalancer() { _ = "STUB: not implemented"; return }

func (b *Balancer) BalanceEagerly() { _ = "STUB: not implemented"; return }

func (b *Balancer) balance() { _ = "STUB: not implemented"; return }

func (b *Balancer) Start() error { _ = "STUB: not implemented"; return nil }

func (b *Balancer) RegisterHandlers() { _ = "STUB: not implemented"; return }

func (b *Balancer) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// already closed

var _ service.Service = (*Balancer)(nil)
