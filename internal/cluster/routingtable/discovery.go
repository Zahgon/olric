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

package routingtable

import (
	"context"
	"errors"
	"time"
)

var (
	ErrServerGone   = errors.New("server is gone")
	ErrNotJoinedYet = errors.New("not joined yet")
	ErrClusterJoin  = errors.New("cannot join the cluster")
	// ErrOperationTimeout is returned when an operation times out.
	ErrOperationTimeout = errors.New("operation timeout")
)

// bootstrapCoordinator prepares the very first routing table and bootstraps the coordinator node.
func (r *RoutingTable) bootstrapCoordinator() error { _ = "STUB: not implemented"; return nil }

// The coordinator bootstraps itself.

func (r *RoutingTable) attemptToJoin() error { _ = "STUB: not implemented"; return nil }

// The node is gone.

func (r *RoutingTable) tryWithInterval(ctx context.Context, interval time.Duration, f func() error) error {
	_ = "STUB: not implemented"
	return nil
}

// Done. No need to try with interval

// context is done
