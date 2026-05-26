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
	"sync"

	"github.com/olric-data/olric/internal/discovery"
)

type Members struct {
	sync.RWMutex
	m map[uint64]discovery.Member
}

func newMembers() *Members { _ = "STUB: not implemented"; return nil }

func (m *Members) Add(member discovery.Member) { _ = "STUB: not implemented"; return }

func (m *Members) Get(id uint64) (discovery.Member, error) {
	_ = "STUB: not implemented"
	return *new(discovery.Member), nil
}

func (m *Members) Delete(id uint64) { _ = "STUB: not implemented"; return }

func (m *Members) DeleteByName(other discovery.Member) { _ = "STUB: not implemented"; return }

func (m *Members) Length() int { _ = "STUB: not implemented"; return 0 }

func (m *Members) Range(f func(id uint64, member discovery.Member) bool) {
	_ = "STUB: not implemented"
	return
}
