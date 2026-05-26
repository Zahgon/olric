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
	"github.com/olric-data/olric/pkg/storage"
)

type transferIterator struct {
	storage *RamBlock
}

func (t *transferIterator) Next() bool { _ = "STUB: not implemented"; return false }

func (t *transferIterator) Drop(index int) error { _ = "STUB: not implemented"; return nil }

func (t *transferIterator) Export() ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (rb *RamBlock) Import(data []byte, f func(uint64, storage.Entry) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (rb *RamBlock) TransferIterator() storage.TransferIterator {
	_ = "STUB: not implemented"
	return *new(storage.TransferIterator)
}
