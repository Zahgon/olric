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
	"errors"
	"time"

	"github.com/olric-data/olric/internal/cluster/partitions"
	"github.com/olric-data/olric/pkg/storage"
)

const nilTimeout = 0 * time.Second

var (
	// ErrKeyNotFound is returned when a key could not be found.
	ErrKeyNotFound  = errors.New("key not found")
	ErrDMapNotFound = errors.New("dmap not found")
	ErrServerGone   = errors.New("server is gone")
)

// DMap implements a single-hop distributed hash table.
type DMap struct {
	name         string
	fragmentName string
	s            *Service
	engine       storage.Engine
	config       *dmapConfig
}

// Name exposes name of the DMap.
func (dm *DMap) Name() string {
	_ = "STUB: not implemented"

	// getDMap returns an initialized DMap instance, otherwise it returns ErrDMapNotFound.
	return ""
}

func (s *Service) getDMap(name string) (*DMap, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Service) fragmentName(name string) string { _ = "STUB: not implemented"; return "" }

// NewDMap creates and returns a new DMap instance. It checks member count quorum
// and bootstrapping status before creating a new DMap.
func (s *Service) NewDMap(name string) (*DMap, error) {
	_ = "STUB: not implemented"
	// Check operation status first:
	//
	//   - Checks member count in the cluster, returns ErrClusterQuorum if
	//     the quorum value cannot be satisfied,
	//   - Checks bootstrapping status and awaits for a short period before
	//     returning ErrRequest timeout.
	return nil, nil
}

// An Olric node has to be bootstrapped to function properly.

// It's a shortcut.

// getOrCreate is a shortcut function to create a new DMap or get an already initialized DMap instance.
func (s *Service) getOrCreateDMap(name string) (*DMap, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dm *DMap) getPartitionByHKey(hkey uint64, kind partitions.Kind) *partitions.Partition {
	_ = "STUB: not implemented"
	return nil
}

func isKeyExpired(ttl int64) bool { _ = "STUB: not implemented"; return false }

// convert nanoseconds to milliseconds

// number of valid items removed from cache to free memory for new items.
