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

package protocol

import (
	"errors"
	"sync"

	"github.com/tidwall/redcon"
)

var ErrInvalidArgument = errors.New("invalid argument")

var GenericError = "ERR"

var errorWithPrefix = struct {
	mtx    sync.RWMutex
	prefix map[string]error
	err    map[error]string
}{
	prefix: make(map[string]error),
	err:    make(map[error]string),
}

func init() {
	SetError("INVALIDARGUMENT", ErrInvalidArgument)
}

func SetError(prefix string, err error) { _ = "STUB: not implemented"; return }

func GetError(prefix string) error { _ = "STUB: not implemented"; return nil }

func getPrefix(err error) string { _ = "STUB: not implemented"; return "" }

func GetPrefix(err error) string { _ = "STUB: not implemented"; return "" }

func ConvertError(err error) error { _ = "STUB: not implemented"; return nil }

func WriteError(conn redcon.Conn, err error) { _ = "STUB: not implemented"; return }

func errWrongNumber(args [][]byte) error { _ = "STUB: not implemented"; return nil }
