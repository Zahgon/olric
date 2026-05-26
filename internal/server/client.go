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

package server

import (
	"context"
	"sync"

	"github.com/olric-data/olric/config"
	"github.com/olric-data/olric/internal/roundrobin"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	mu sync.RWMutex

	config     *config.Client
	clients    map[string]*redis.Client
	roundRobin *roundrobin.RoundRobin
}

func NewClient(c *config.Client) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) Addresses() map[string]struct{} { _ = "STUB: not implemented"; return nil }

func (c *Client) Get(addr string) *redis.Client { _ = "STUB: not implemented"; return nil }

// Need the lock for writing, we modify c.clients map and the round-robin
// implementation updates its internal state.

// Need to check again, because another goroutine may have updated clients
// between our calls to RUnlock and Lock.

func (c *Client) pickNodeRoundRobin() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (c *Client) Pick() (*redis.Client, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Client) Close(addr string) error { _ = "STUB: not implemented"; return nil }

func (c *Client) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
