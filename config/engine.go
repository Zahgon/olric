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
	"github.com/olric-data/olric/pkg/storage"
)

// Engine contains storage engine configuration and their implementations.
// If you don't have a custom storage engine implementation or configuration for
// the default one, just call NewStorageEngine() function to use it with sane defaults.
type Engine struct {
	Name string

	Implementation storage.Engine

	// Config is a map that contains configuration of the storage engines, for
	// both plugins and imported ones. If you want to use a storage engine other
	// than the default one, you must set configuration for it.
	Config map[string]interface{}
}

// NewEngine initializes Engine with sane defaults.
// Olric will set its own storage engine implementation and related configuration,
// if there is no other engine.
func NewEngine() *Engine { _ = "STUB: not implemented"; return nil }

// Validate finds errors in the current configuration.
func (s *Engine) Validate() error { _ = "STUB: not implemented"; return nil }

// Sanitize sets default values to empty configuration variables, if it's possible.
func (s *Engine) Sanitize() error { _ = "STUB: not implemented"; return nil }

// Backward compatibility: accept the old name "kvstore"

// Interface guard
var _ IConfig = (*Engine)(nil)
