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
)

func (dm *DMap) destroyOnCluster(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Don't block routing table to destroy a DMap on the cluster.
// Just get a copy of members and run Destroy.

// Destroy flushes the given DMap on the cluster. You should know that there
// is no global lock on DMaps. So if you call Put, Put with EX and Destroy methods
// concurrently on the cluster, Put and Put with EX calls may set new values to the DMap.
func (dm *DMap) Destroy(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
