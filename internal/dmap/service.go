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

package dmap

import (
	"context"
	"errors"
	"sync"

	"github.com/olric-data/olric/config"
	"github.com/olric-data/olric/events"
	"github.com/olric-data/olric/internal/cluster/partitions"
	"github.com/olric-data/olric/internal/cluster/routingtable"
	"github.com/olric-data/olric/internal/environment"
	"github.com/olric-data/olric/internal/locker"
	"github.com/olric-data/olric/internal/server"
	"github.com/olric-data/olric/internal/service"
	"github.com/olric-data/olric/pkg/flog"
	"github.com/olric-data/olric/pkg/storage"
)

var errFragmentNotFound = errors.New("fragment not found")

type storageMap struct {
	engines map[string]storage.Engine
	configs map[string]map[string]interface{}
}

type Service struct {
	sync.RWMutex // protects dmaps map

	log     *flog.Logger
	config  *config.Config
	client  *server.Client
	server  *server.Server
	rt      *routingtable.RoutingTable
	primary *partitions.Partitions
	backup  *partitions.Partitions
	locker  *locker.Locker
	dmaps   map[string]*DMap
	storage *storageMap
	wg      sync.WaitGroup
	ctx     context.Context
	cancel  context.CancelFunc
}

func registerErrors() { _ = "STUB: not implemented"; return }

func NewService(e *environment.Environment) (service.Service, error) {
	_ = "STUB: not implemented"
	return *new(service.Service), nil
}

func (s *Service) isAlive() bool { _ = "STUB: not implemented"; return false }

// The node is gone.

func getType(data interface{}) string { _ = "STUB: not implemented"; return "" }

func (s *Service) publishEvent(e events.Event) { _ = "STUB: not implemented"; return }

// Start starts the distributed map service.
func (s *Service) Start() error { _ = "STUB: not implemented"; return nil }

func (s *Service) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

var _ service.Service = (*Service)(nil)
