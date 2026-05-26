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

package olric

import (
	"errors"
	"time"

	"github.com/olric-data/olric/pkg/storage"
)

var ErrNilResponse = errors.New("storage entry is nil")

type GetResponse struct {
	entry storage.Entry
}

func (g *GetResponse) Scan(v interface{}) error { _ = "STUB: not implemented"; return nil }

func (g *GetResponse) Int() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (g *GetResponse) String() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (g *GetResponse) Int8() (int8, error) { _ = "STUB: not implemented"; return 0, nil }

func (g *GetResponse) Int16() (int16, error) { _ = "STUB: not implemented"; return 0, nil }

func (g *GetResponse) Int32() (int32, error) { _ = "STUB: not implemented"; return 0, nil }

func (g *GetResponse) Int64() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (g *GetResponse) Uint() (uint, error) { _ = "STUB: not implemented"; return 0, nil }

func (g *GetResponse) Uint8() (uint8, error) { _ = "STUB: not implemented"; return 0, nil }

func (g *GetResponse) Uint16() (uint16, error) { _ = "STUB: not implemented"; return 0, nil }

func (g *GetResponse) Uint32() (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func (g *GetResponse) Uint64() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (g *GetResponse) Float32() (float32, error) { _ = "STUB: not implemented"; return 0, nil }

func (g *GetResponse) Float64() (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func (g *GetResponse) Bool() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (g *GetResponse) Time() (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func (g *GetResponse) Duration() (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func (g *GetResponse) Byte() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (g *GetResponse) TTL() int64 { _ = "STUB: not implemented"; return 0 }

func (g *GetResponse) Timestamp() int64 { _ = "STUB: not implemented"; return 0 }
