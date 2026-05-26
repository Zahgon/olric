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
	"sync"

	"github.com/olric-data/olric/internal/cluster/partitions"
	"github.com/olric-data/olric/internal/discovery"
	"github.com/olric-data/olric/pkg/storage"
)

type fragment struct {
	sync.RWMutex

	service *Service
	storage storage.Engine
	ctx     context.Context
	cancel  context.CancelFunc
}

func (f *fragment) Stats() storage.Stats { _ = "STUB: not implemented"; return *new(storage.Stats) }

func (f *fragment) Compaction() (bool, error) {
	_ = "STUB: not implemented"
	return false,

		// fragment is closed or destroyed
		nil
}

func (f *fragment) Destroy() error { _ = "STUB: not implemented"; return nil }

func (f *fragment) Close() error { _ = "STUB: not implemented"; return nil }

func (f *fragment) Name() string { _ = "STUB: not implemented"; return "" }

func (f *fragment) Move(part *partitions.Partition, name string, owners []discovery.Member) error {
	_ = "STUB: not implemented"
	return nil
}

func (dm *DMap) newFragment() (*fragment, error) { _ = "STUB: not implemented"; return nil, nil }

func (dm *DMap) loadOrCreateFragment(part *partitions.Partition) (*fragment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Critical section here. It should be protected by a lock.

// We already have the fragment.

func (dm *DMap) loadFragment(part *partitions.Partition) (*fragment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var _ partitions.Fragment = (*fragment)(nil)
