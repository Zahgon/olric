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
	"github.com/olric-data/olric/internal/discovery"
)

func (r *RoutingTable) processLeftOverDataReports(reports map[discovery.Member]*leftOverDataReport) {
	_ = "STUB: not implemented"
	return
}

// This section is protected by routingMtx against parallel writers.
//
// Copy owners and append the member to head

// Prepend

// data structures in this function is guarded by routingMtx
