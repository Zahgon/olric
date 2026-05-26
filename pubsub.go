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
	"context"

	"github.com/olric-data/olric/internal/server"
	"github.com/redis/go-redis/v9"
)

type PubSub struct {
	config *pubsubConfig
	rc     *redis.Client
	client *server.Client
}

func newPubSub(client *server.Client, options ...PubSubOption) (*PubSub, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ps *PubSub) Subscribe(ctx context.Context, channels ...string) *redis.PubSub {
	_ = "STUB: not implemented"
	return nil
}

func (ps *PubSub) PSubscribe(ctx context.Context, channels ...string) *redis.PubSub {
	_ = "STUB: not implemented"
	return nil
}

func (ps *PubSub) Publish(ctx context.Context, channel string, message interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (ps *PubSub) PubSubChannels(ctx context.Context, pattern string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ps *PubSub) PubSubNumSub(ctx context.Context, channels ...string) (map[string]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ps *PubSub) PubSubNumPat(ctx context.Context) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
