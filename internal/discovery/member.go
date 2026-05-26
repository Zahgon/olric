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

package discovery

import (
	"github.com/olric-data/olric/config"
)

// Member represents a node in the cluster.
type Member struct {
	Name      string
	NameHash  uint64
	ID        uint64
	Birthdate int64
}

// CompareByID returns true if two members denote the same member in the cluster.
func (m Member) CompareByID(other Member) bool {
	_ = "STUB: not implemented"
	// ID variable is calculated by combining member's name and birthdate
	return false
}

// CompareByName returns true if the two members has the same name in the cluster.
// This function is intended to redirect the requests to the partition owner.
func (m Member) CompareByName(other Member) bool { _ = "STUB: not implemented"; return false }

func (m Member) String() string { _ = "STUB: not implemented"; return "" }

func (m Member) Encode() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewMemberFromMetadata(metadata []byte) (Member, error) {
	_ = "STUB: not implemented"
	return *new(Member), nil
}

func MemberID(name string, birthdate int64) uint64 {
	_ = "STUB: not implemented"
	// Calculate member's identity. It's useful to compare hosts.
	return 0
}

func NewMember(c *config.Config) Member { _ = "STUB: not implemented"; return *new(Member) }
