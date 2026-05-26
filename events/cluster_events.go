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

package events

import (
	"reflect"
)

const (
	ClusterEventsChannel       = "cluster.events"
	KindNodeJoinEvent          = "node-join-event"
	KindNodeLeftEvent          = "node-left-event"
	KindFragmentMigrationEvent = "fragment-migration-event"
	KindFragmentReceivedEvent  = "fragment-received-event"
)

type Event interface {
	Encode() (string, error)
}

// encodeEvents encodes given interface to its JSON representation and preserves the order in fields slice.
func encodeEvent(data interface{}, fields []string, valueExtractor func(r reflect.Value, field string) (interface{}, error)) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// marshal key

// marshal value

type NodeJoinEvent struct {
	Kind      string `json:"kind"`
	Source    string `json:"source"`
	NodeJoin  string `json:"node_join"`
	Timestamp int64  `json:"timestamp"`
}

func (n *NodeJoinEvent) Encode() (string, error) { _ = "STUB: not implemented"; return "", nil }

type NodeLeftEvent struct {
	Kind      string `json:"kind"`
	Source    string `json:"source"`
	NodeLeft  string `json:"node_left"`
	Timestamp int64  `json:"timestamp"`
}

func (n *NodeLeftEvent) Encode() (string, error) { _ = "STUB: not implemented"; return "", nil }

type FragmentMigrationEvent struct {
	Kind          string `json:"kind"`
	Source        string `json:"source"`
	Target        string `json:"target"`
	Identifier    string `json:"identifier"`
	PartitionID   uint64 `json:"partition_id"`
	DataStructure string `json:"data_structure"`
	Length        int    `json:"length"`
	IsBackup      bool   `json:"is_backup"`
	Timestamp     int64  `json:"timestamp"`
}

func (f *FragmentMigrationEvent) Encode() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type FragmentReceivedEvent struct {
	Kind          string `json:"kind"`
	Source        string `json:"source"`
	Identifier    string `json:"identifier"`
	PartitionID   uint64 `json:"partition_id"`
	DataStructure string `json:"data_structure"`
	Length        int    `json:"length"`
	IsBackup      bool   `json:"is_backup"`
	Timestamp     int64  `json:"timestamp"`
}

func (f *FragmentReceivedEvent) Encode() (string, error) { _ = "STUB: not implemented"; return "", nil }
