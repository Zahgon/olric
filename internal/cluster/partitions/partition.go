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

package partitions

import (
	"sync"
	"sync/atomic"

	"github.com/olric-data/olric/internal/discovery"
)

// Partition is a basic, logical storage unit in Olric and stores DMaps in a sync.Map.
type Partition struct {
	sync.RWMutex

	id     uint64
	kind   Kind
	m      *sync.Map
	owners atomic.Value
}

func (p *Partition) Kind() Kind { _ = "STUB: not implemented"; return *new(Kind) }

func (p *Partition) ID() uint64 { _ = "STUB: not implemented"; return 0 }

func (p *Partition) Map() *sync.Map {
	_ = "STUB: not implemented"

	// Owner returns partition Owner. It's not thread-safe.
	return nil
}

func (p *Partition) Owner() discovery.Member {
	_ = "STUB: not implemented"
	return *

	// programming error. it cannot occur at production!
	new(discovery.Member)
}

// OwnerCount returns the current Owner count of a partition.
func (p *Partition) OwnerCount() int { _ = "STUB: not implemented"; return 0 }

// Owners loads the partition owners from atomic.Value and returns.
func (p *Partition) Owners() []discovery.Member { _ = "STUB: not implemented"; return nil }

func (p *Partition) SetOwners(owners []discovery.Member) { _ = "STUB: not implemented"; return }

func (p *Partition) Length() int { _ = "STUB: not implemented"; return 0 }

// Continue scanning.
