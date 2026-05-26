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
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/tidwall/redcon"
)

type Publish struct {
	Channel string
	Message string
}

func NewPublish(channel, message string) *Publish { _ = "STUB: not implemented"; return nil }

func (p *Publish) Command(ctx context.Context) *redis.IntCmd { _ = "STUB: not implemented"; return nil }

func ParsePublishCommand(cmd redcon.Command) (*Publish, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Channel
// Message

type PublishInternal struct {
	Channel string
	Message string
}

func NewPublishInternal(channel, message string) *PublishInternal {
	_ = "STUB: not implemented"
	return nil
}

func (p *PublishInternal) Command(ctx context.Context) *redis.IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func ParsePublishInternalCommand(cmd redcon.Command) (*PublishInternal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Channel
// Message

type Subscribe struct {
	Channels []string
}

func NewSubscribe(channels ...string) *Subscribe { _ = "STUB: not implemented"; return nil }

func (s *Subscribe) Command(ctx context.Context) *redis.SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func ParseSubscribeCommand(cmd redcon.Command) (*Subscribe, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type PSubscribe struct {
	Patterns []string
}

func NewPSubscribe(patterns ...string) *PSubscribe { _ = "STUB: not implemented"; return nil }

func (s *PSubscribe) Command(ctx context.Context) *redis.SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func ParsePSubscribeCommand(cmd redcon.Command) (*PSubscribe, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type PubSubChannels struct {
	Pattern string
}

func NewPubSubChannels() *PubSubChannels { _ = "STUB: not implemented"; return nil }

func (ps *PubSubChannels) SetPattern(pattern string) *PubSubChannels {
	_ = "STUB: not implemented"
	return nil
}

func (ps *PubSubChannels) Command(ctx context.Context) *redis.SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func ParsePubSubChannelsCommand(cmd redcon.Command) (*PubSubChannels, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type PubSubNumpat struct{}

func NewPubSubNumpat() *PubSubNumpat { _ = "STUB: not implemented"; return nil }

func (ps *PubSubNumpat) Command(ctx context.Context) *redis.IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func ParsePubSubNumpatCommand(cmd redcon.Command) (*PubSubNumpat, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type PubSubNumsub struct {
	Channels []string
}

func NewPubSubNumsub(channels ...string) *PubSubNumsub { _ = "STUB: not implemented"; return nil }

func (ps *PubSubNumsub) Command(ctx context.Context) *redis.SliceCmd {
	_ = "STUB: not implemented"
	return nil
}

func ParsePubSubNumsubCommand(cmd redcon.Command) (*PubSubNumsub, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
