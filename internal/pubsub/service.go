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

package pubsub

import (
	"context"
	"sync"

	"github.com/olric-data/olric/internal/cluster/routingtable"
	"github.com/olric-data/olric/internal/environment"
	"github.com/olric-data/olric/internal/server"
	"github.com/olric-data/olric/internal/service"
	"github.com/olric-data/olric/internal/stats"
	"github.com/olric-data/olric/pkg/flog"
)

var (
	// PublishedTotal is the total number of published messages during the life of this instance.
	PublishedTotal = stats.NewInt64Counter()

	// CurrentSubscribers is the current number of listeners of Pub/Sub.
	CurrentSubscribers = stats.NewInt64Gauge()

	// SubscribersTotal is the total number of registered listeners during the life of this instance.
	SubscribersTotal = stats.NewInt64Counter()

	CurrentPSubscribers = stats.NewInt64Gauge()
	PSubscribersTotal   = stats.NewInt64Counter()
)

type Service struct {
	sync.RWMutex

	log    *flog.Logger
	pubsub *PubSub
	rt     *routingtable.RoutingTable
	server *server.Server
	client *server.Client
	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

func (s *Service) RegisterHandlers() { _ = "STUB: not implemented"; return }

func NewService(e *environment.Environment) (service.Service, error) {
	_ = "STUB: not implemented"
	return *new(service.Service), nil
}

func (s *Service) Start() error {
	_ = "STUB: not implemented"
	// dummy implementation
	return nil
}

func (s *Service) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

var _ service.Service = (*Service)(nil)
