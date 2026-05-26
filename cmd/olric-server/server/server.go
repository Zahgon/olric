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

/*Package server provides a standalone server implementation for Olric*/
package server

import (
	"context"
	"log"

	"github.com/olric-data/olric"
	"github.com/olric-data/olric/config"
	"golang.org/x/sync/errgroup"
)

// OlricServer represents an instance of the Olric distributed in-memory data structure store.
// It encapsulates logging, configuration, the Olric database instance, and an error group for
// concurrency management.
type OlricServer struct {
	log    *log.Logger
	config *config.Config
	db     *olric.Olric
	errGr  errgroup.Group
}

// New initializes a new OlricServer instance using the provided configuration and returns it or an error.
func New(c *config.Config) (*OlricServer, error) { _ = "STUB: not implemented"; return nil, nil }

// waitForInterrupt waits for termination signals (SIGTERM, SIGINT) to gracefully shut down the Olric server instance.
func (s *OlricServer) waitForInterrupt() { _ = "STUB: not implemented"; return }

// Awaits for shutdown

// This is not a goroutine leak. The process will quit.

// Start launches the Olric server instance and begins listening for incoming requests and termination signals.
func (s *OlricServer) Start() error { _ = "STUB: not implemented"; return nil }

// Wait for SIGTERM or SIGINT

// Shutdown gracefully stops the Olric server instance, releasing resources and ensuring a clean termination.
func (s *OlricServer) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
