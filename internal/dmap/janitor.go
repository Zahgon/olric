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
	"github.com/olric-data/olric/internal/cluster/partitions"
)

func wipeOutFragment(part *partitions.Partition, name string, f *fragment) error {
	_ = "STUB: not implemented"
	// Stop background services if there is any.
	return nil
}

// Destroy data on-disk or in-memory.

// Delete the fragment from partition.

func (s *Service) janitor(part *partitions.Partition) { _ = "STUB: not implemented"; return }

// This fragment belongs to a different data structure.

// It's not empty. Continue scanning.

// continue scanning

func (s *Service) deleteEmptyFragments() { _ = "STUB: not implemented"; return }

// Clean stale DMap fragments on partition table

// Clean stale DMap fragments on backup partition table

func (s *Service) janitorWorker() { _ = "STUB: not implemented"; return }
