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

package ramblock

import (
	"github.com/olric-data/olric/internal/ramblock/table"
)

func (rb *RamBlock) evictTable(t *table.Table) error { _ = "STUB: not implemented"; return nil }

// try again

// log this error and continue

func (rb *RamBlock) isTableExpired(recycledAt int64) bool { _ = "STUB: not implemented"; return false }

// That would be impossible

func (rb *RamBlock) isCompactionOK(t *table.Table) bool { _ = "STUB: not implemented"; return false }

func (rb *RamBlock) Compaction() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Continue scanning
