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

// delegate is a struct which implements memberlist.Delegate interface.
type delegate struct {
	meta []byte
}

// newDelegate returns a new delegate instance.
func (d *Discovery) newDelegate() (delegate, error) {
	_ = "STUB: not implemented"
	return *new(delegate), nil
}

// NodeMeta is used to retrieve meta-data about the current node
// when broadcasting an alive message. It's length is limited to
// the given byte size. This metadata is available in the Node structure.
func (d delegate) NodeMeta(limit int) []byte {
	_ = "STUB: not implemented"

	// NotifyMsg is called when a user-data message is received.
	return nil
}

func (d delegate) NotifyMsg(data []byte) {
	_ = "STUB: not implemented"

	// GetBroadcasts is called when user data messages can be broadcast.
	return
}

func (d delegate) GetBroadcasts(overhead, limit int) [][]byte {
	_ = "STUB: not implemented"

	// LocalState is used for a TCP Push/Pull.
	return nil
}

func (d delegate) LocalState(join bool) []byte {
	_ = "STUB: not implemented"

	// MergeRemoteState is invoked after a TCP Push/Pull.
	return nil
}

func (d delegate) MergeRemoteState(buf []byte, join bool) { _ = "STUB: not implemented"; return }
