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
	"errors"
	"time"
)

var (
	// ErrLockNotAcquired is returned when the requested lock could not be acquired
	ErrLockNotAcquired = errors.New("lock not acquired")

	// ErrNoSuchLock is returned when the requested lock does not exist
	ErrNoSuchLock = errors.New("no such lock")
)

// unlockKey tries to unlock the lock by verifying the lock with token.
func (dm *DMap) unlockKey(ctx context.Context, key string, token []byte) error {
	_ = "STUB: not implemented"
	return nil

	// Only one unlockKey should work for a given key.
}

// get the key to check its value

// the lock is released by the node(timeout) or the user

// release it.

// Unlock takes key and token and tries to unlock the key.
// It redirects the request to the partition owner, if required.
func (dm *DMap) Unlock(ctx context.Context, key string, token []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// tryLock takes a deadline and env and sets a key-value pair by using
// Put with NX and PX commands. It tries to acquire the lock 100 times per second
// if the lock is already acquired. It returns ErrLockNotAcquired if the deadline exceeds.
func (dm *DMap) tryLock(e *env, deadline time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// If it returns ErrKeyFound, the lock is already acquired.

// something went wrong

// Try to acquire lock.

// not released by the other process/goroutine. try again.

// something went wrong.

// Acquired! Quit without error.

// Deadline exceeded. Quit with an error.

// Lock prepares a token and env, then calls tryLock
func (dm *DMap) Lock(ctx context.Context, key string, timeout, deadline time.Duration) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// leaseKey tries to update the expiry of the key by verifying token.
func (dm *DMap) leaseKey(ctx context.Context, key string, token []byte, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil

	// Only one unlockKey should work for a given key.
}

// get the key to check its value

// the lock is released by the node(timeout) or the user

// already expired

// update

// Lease takes key and token and tries to update the expiry with duration.
// It redirects the request to the partition owner, if required.
func (dm *DMap) Lease(ctx context.Context, key string, token []byte, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}
