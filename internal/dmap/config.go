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

package dmap

import (
	"time"

	"github.com/olric-data/olric/config"
)

// dmapConfig keeps DMap config control parameters and access-log for keys in a dmap.
type dmapConfig struct {
	engine          *config.Engine
	maxIdleDuration time.Duration
	ttlDuration     time.Duration
	maxKeys         int
	maxInuse        int
	lruSamples      int
	evictionPolicy  config.EvictionPolicy
}

func (c *dmapConfig) load(dc *config.DMaps, name string) error {
	_ = "STUB: not implemented"
	// Try to set config configuration for this dmap.
	return nil
}

// config.DMap struct can be used for fine-grained control.

//TODO: Create a new function to verify config.

// set the default value.
