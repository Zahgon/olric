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
	"github.com/olric-data/olric/pkg/storage"
	"github.com/tidwall/redcon"
)

type fragmentPack struct {
	PartID  uint64
	Kind    partitions.Kind
	Name    string
	Payload []byte
}

func (dm *DMap) fragmentMergeFunction(f *fragment, hkey uint64, entry storage.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// No need to insert the winner

func (dm *DMap) mergeFragments(part *partitions.Partition, fp *fragmentPack) error {
	_ = "STUB: not implemented"
	return nil
}

// Acquire fragment's lock. No one should work on it.

func (s *Service) checkOwnership(part *partitions.Partition) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *Service) validateFragmentPack(fp *fragmentPack) error {
	_ = "STUB: not implemented"
	return nil
}

// Check ownership before merging. This is useful to prevent data corruption in network partitioning case.

func (s *Service) moveFragmentCommandHandler(conn redcon.Conn, cmd redcon.Command) {
	_ = "STUB: not implemented"
	return
}
