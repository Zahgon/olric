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
	"github.com/olric-data/olric/internal/discovery"
)

type Kind int

func (k Kind) String() string { _ = "STUB: not implemented"; return "" }

const (
	PRIMARY = Kind(iota + 1)
	BACKUP
)

type Partitions struct {
	count uint64
	kind  Kind
	m     map[uint64]*Partition
}

func New(count uint64, kind Kind) *Partitions { _ = "STUB: not implemented"; return nil }

// PartitionByID returns the partition for the given HKey
func (ps *Partitions) PartitionByID(partID uint64) *Partition {
	_ = "STUB: not implemented"
	return nil

	// PartitionIDByHKey returns partition ID for a given HKey.
}

func (ps *Partitions) PartitionIDByHKey(hkey uint64) uint64 { _ = "STUB: not implemented"; return 0 }

// PartitionByHKey returns the partition for the given HKey
func (ps *Partitions) PartitionByHKey(hkey uint64) *Partition {
	_ = "STUB: not implemented"
	return nil
}

// PartitionOwnersByHKey loads the partition owners list for a given hkey.
func (ps *Partitions) PartitionOwnersByHKey(hkey uint64) []discovery.Member {
	_ = "STUB: not implemented"
	return nil
}

// PartitionOwnersByID loads the partition owners list for a given hkey.
func (ps *Partitions) PartitionOwnersByID(partID uint64) []discovery.Member {
	_ = "STUB: not implemented"
	return nil
}
