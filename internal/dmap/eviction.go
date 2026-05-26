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

// isKeyIdleOnFragment is not a thread-safe function. It accesses underlying fragment for the given hkey.
func (dm *DMap) isKeyIdleOnFragment(hkey uint64, f *fragment) bool {
	_ = "STUB: not implemented"
	return false
}

// Maximum time in seconds for each entry to stay idle in the map.
// It limits the lifetime of the entries relative to the time of the last
// read or write access performed on them. The entries whose idle period
// exceeds this limit are expired and evicted automatically.

//TODO: Handle other errors.

func (dm *DMap) isKeyIdle(hkey uint64) bool { _ = "STUB: not implemented"; return false }

// it's no possible to know whether the key is idle or not.

// This could be a programming error and should never be happened on production systems.

func (s *Service) evictKeysAtBackground() { _ = "STUB: not implemented"; return }

// Good for developing tests.

func (s *Service) evictKeys() { _ = "STUB: not implemented"; return }

// this breaks the loop, we only scan one dmap instance per call

func (s *Service) scanFragmentForEviction(partID uint64, name string, f *fragment) {
	_ = "STUB: not implemented"
	/*
		From Redis Docs:
			1- Test 20 random keys from the set of keys with an associated expire.
			2- Delete all the keys found expired.
			3- If more than 25% of keys were expired, start again from step 1.
	*/return
}

// We need limits to prevent CPU starvation. deleteOnCluster does some network operation
// to delete keys from the backup nodes and the previous owners.

// Release the lock. Eviction will be triggered again.

// this means 'break'.

// continue

// continue

// It will be tried again.

// number of valid items removed from cache to free memory for new items.

// the fragment is closed.

// The server has gone.

// Call janitorWorker again until it returns false.

type lruItem struct {
	HKey       uint64
	LastAccess int64
}

func (dm *DMap) evictKeyWithLRU(e *env) error { _ = "STUB: not implemented"; return nil }

// Warning: fragment is already locked by DMap.Put. Be sure about that before editing this function.

// Pick random items from the distributed map and sort them by accessedAt.

// Pick the first item to delete. It's the least recently used item in the sample.

// Here we have a key/value pair to evict for making room for a new pair.

// number of valid items removed from cache to free memory for new items.
