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

package testutil

import (
	"testing"
	"time"

	"github.com/olric-data/olric/config"
	"github.com/olric-data/olric/internal/server"
	"github.com/olric-data/olric/pkg/flog"
)

func GetFreePort() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func NewFlogger(c *config.Config) *flog.Logger { _ = "STUB: not implemented"; return nil }

func NewEngineConfig(t *testing.T) *config.Engine { _ = "STUB: not implemented"; return nil }

func NewConfig() *config.Config { _ = "STUB: not implemented"; return nil }

func NewServer(c *config.Config) *server.Server { _ = "STUB: not implemented"; return nil }

func TryWithInterval(max int, interval time.Duration, f func() error) error {
	_ = "STUB: not implemented"
	return nil
}

// Done. No need to try with interval

func ToKey(i int) string { _ = "STUB: not implemented"; return "" }

func ToVal(i int) []byte { _ = "STUB: not implemented"; return nil }
