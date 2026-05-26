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

package entry

import (
	"github.com/olric-data/olric/pkg/storage"
)

// In-memory layout for an entry:
//
// KEY-LENGTH(uint8) | KEY(bytes) | TTL(uint64) | | Timestamp(uint64) | VALUE-LENGTH(uint32) | VALUE(bytes)

// Entry represents a value with its metadata.
type Entry struct {
	key        string
	ttl        int64
	timestamp  int64
	lastAccess int64
	value      []byte
}

var _ storage.Entry = (*Entry)(nil)

func New() *Entry { _ = "STUB: not implemented"; return nil }

func (e *Entry) SetKey(key string) { _ = "STUB: not implemented"; return }

func (e *Entry) Key() string { _ = "STUB: not implemented"; return "" }

func (e *Entry) SetValue(value []byte) { _ = "STUB: not implemented"; return }

func (e *Entry) Value() []byte { _ = "STUB: not implemented"; return nil }

func (e *Entry) SetTTL(ttl int64) { _ = "STUB: not implemented"; return }

func (e *Entry) TTL() int64 { _ = "STUB: not implemented"; return 0 }

func (e *Entry) SetTimestamp(timestamp int64) { _ = "STUB: not implemented"; return }

func (e *Entry) Timestamp() int64 { _ = "STUB: not implemented"; return 0 }

func (e *Entry) SetLastAccess(lastAccess int64) { _ = "STUB: not implemented"; return }

func (e *Entry) LastAccess() int64 { _ = "STUB: not implemented"; return 0 }

func (e *Entry) Encode() []byte { _ = "STUB: not implemented"; return nil }

// Set key length. It's 1 byte.

// Set the key.

// Set the TTL. It's 8 bytes.

// Set the Timestamp. It's 8 bytes.

// Set the LastAccess. It's 8 bytes.

// Set the value length. It's 4 bytes.

// Set the value.

func (e *Entry) Decode(buf []byte) { _ = "STUB: not implemented"; return }
