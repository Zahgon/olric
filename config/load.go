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

package config

import (
	"github.com/hashicorp/memberlist"
	"github.com/olric-data/olric/config/internal/loader"
)

// mapYamlToConfig maps a parsed YAML to related configuration struct.
func mapYamlToConfig(rawDst, rawSrc interface{}) error { _ = "STUB: not implemented"; return nil }

// Special cases

func loadDMapConfig(c *loader.Loader) (*DMaps, error) { _ = "STUB: not implemented"; return nil, nil }

// loadMemberlistConfig creates a new *memberlist.Config by parsing olric.yaml
func loadMemberlistConfig(c *loader.Loader, mc *memberlist.Config) (*memberlist.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Load reads and loads Olric configuration.
func Load(filename string) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }
