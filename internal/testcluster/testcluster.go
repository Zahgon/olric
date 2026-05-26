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

package testcluster

import (
	"context"
	"sync"

	"github.com/olric-data/olric/config"
	"github.com/olric-data/olric/internal/environment"
	"github.com/olric-data/olric/internal/service"
	"golang.org/x/sync/errgroup"
)

type TestCluster struct {
	mu sync.Mutex

	environments []*environment.Environment
	memberPorts  []int
	constructor  func(e *environment.Environment) (service.Service, error)
	errGr        errgroup.Group
	ctx          context.Context
	cancel       context.CancelFunc
}

func NewEnvironment(c *config.Config) *environment.Environment {
	_ = "STUB: not implemented"
	return nil
}

func (t *TestCluster) newService(e *environment.Environment) service.Service {
	_ = "STUB: not implemented"
	return *new(service.Service)
}

func New(constructor func(e *environment.Environment) (service.Service, error)) *TestCluster {
	_ = "STUB: not implemented"
	return nil
}

func (t *TestCluster) syncCluster() {
	_ = "STUB: not implemented"
	// Update routing table on the cluster before running balancer
	return
}

// The coordinator pushes the routing table immediately.
// Normally, this is triggered by every cluster event but we don't want to
// do this asynchronously to avoid randomness in tests.

// Normally, balancer is triggered by routing table after a successful update, but we don't want to
// balance the test cluster asynchronously. So we balance the partitions here explicitly.

func (t *TestCluster) AddMember(e *environment.Environment) service.Service {
	_ = "STUB: not implemented"
	return *new(service.Service)
}

func (t *TestCluster) Shutdown() { _ = "STUB: not implemented"; return }
