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

	"github.com/olric-data/olric/pkg/storage"
)

func (dm *DMap) loadCurrentAtomicInt(e *env) (int, int64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func (dm *DMap) atomicIncrDecr(cmd string, e *env, delta int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Incr atomically increments key by delta. The return value is the new value after being incremented or an error.
func (dm *DMap) Incr(ctx context.Context, key string, delta int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Decr atomically decrements key by delta. The return value is the new value after being decremented or an error.
func (dm *DMap) Decr(ctx context.Context, key string, delta int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (dm *DMap) getPut(e *env) (storage.Entry, error) {
	_ = "STUB: not implemented"
	return *new(storage.Entry), nil
}

// The value is nil.

// GetPut atomically sets key to value and returns the old value stored at key.
func (dm *DMap) GetPut(ctx context.Context, key string, value interface{}) (storage.Entry, error) {
	_ = "STUB: not implemented"
	return *new(storage.Entry), nil
}

func (dm *DMap) atomicIncrByFloat(e *env, delta float64) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// IncrByFloat atomically increments key by delta. The return value is the new value after being incremented or an error.
func (dm *DMap) IncrByFloat(ctx context.Context, key string, delta float64) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
