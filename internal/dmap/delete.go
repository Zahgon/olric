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

	"github.com/olric-data/olric/internal/cluster/partitions"
	"github.com/olric-data/olric/internal/discovery"
	"github.com/olric-data/olric/internal/stats"
)

var (
	// DeleteHits is the number of deletion requests resulting in an item being removed.
	DeleteHits = stats.NewInt64Counter()

	// DeleteMisses is the number of deletion requests for missing keys.
	DeleteMisses = stats.NewInt64Counter()
)

func (dm *DMap) deleteFromFragment(key string, kind partitions.Kind) error {
	_ = "STUB: not implemented"
	return nil
}

// key doesn't exist

func (dm *DMap) deleteFromPreviousOwners(key string, owners []discovery.Member) error {
	_ = "STUB: not implemented"
	// Traverse in reverse order. Except from the latest host, this one.
	return nil
}

func (dm *DMap) deleteBackupOnCluster(hkey uint64, key string) error {
	_ = "STUB: not implemented"
	return nil
}

// deleteOnCluster is not a thread-safe function
func (dm *DMap) deleteOnCluster(hkey uint64, key string, f *fragment) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteHits is the number of deletion reqs resulting in an item being removed.

func (dm *DMap) deleteKey(key string) error { _ = "STUB: not implemented"; return nil }

// Check the HKey before trying to delete it.

// DeleteMisses is the number of deletions reqs for missing keys

func (dm *DMap) deleteKeys(ctx context.Context, keys ...string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Delete deletes the value for the given key. Delete will not return error if key doesn't exist. It's thread-safe.
// It is safe to modify the contents of the argument after Delete returns.
func (dm *DMap) Delete(ctx context.Context, keys ...string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
