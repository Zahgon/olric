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

type Ping struct {
	Message string
}

func NewPing() *Ping { _ = "STUB: not implemented"; return nil }

func (p *Ping) SetMessage(m string) *Ping { _ = "STUB: not implemented"; return nil }

func (p *Ping) Command(ctx context.Context) *redis.StringCmd { _ = "STUB: not implemented"; return nil }

func ParsePingCommand(cmd redcon.Command) (*Ping, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type MoveFragment struct {
	Payload []byte
}

func NewMoveFragment(payload []byte) *MoveFragment { _ = "STUB: not implemented"; return nil }

func (m *MoveFragment) Command(ctx context.Context) *redis.StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func ParseMoveFragmentCommand(cmd redcon.Command) (*MoveFragment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type UpdateRouting struct {
	Payload       []byte
	CoordinatorID uint64
}

func NewUpdateRouting(payload []byte, coordinatorID uint64) *UpdateRouting {
	_ = "STUB: not implemented"
	return nil
}

func (u *UpdateRouting) Command(ctx context.Context) *redis.StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func ParseUpdateRoutingCommand(cmd redcon.Command) (*UpdateRouting, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type LengthOfPart struct {
	PartID  uint64
	Replica bool
}

func NewLengthOfPart(partID uint64) *LengthOfPart { _ = "STUB: not implemented"; return nil }

func (l *LengthOfPart) SetReplica() *LengthOfPart { _ = "STUB: not implemented"; return nil }

func (l *LengthOfPart) Command(ctx context.Context) *redis.IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func ParseLengthOfPartCommand(cmd redcon.Command) (*LengthOfPart, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Stats struct {
	CollectRuntime bool
}

func NewStats() *Stats { _ = "STUB: not implemented"; return nil }

func (s *Stats) SetCollectRuntime() *Stats { _ = "STUB: not implemented"; return nil }

func (s *Stats) Command(ctx context.Context) *redis.StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func ParseStatsCommand(cmd redcon.Command) (*Stats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Auth represents a structure for authentication containing a password.
type Auth struct {
	Password string
}

// NewAuth creates and returns a new Auth instance initialized with the given password.
func NewAuth(password string) *Auth { _ = "STUB: not implemented"; return nil }

// Command constructs a Redis AUTH command using the provided authentication password from the Auth instance.
func (a *Auth) Command(ctx context.Context) *redis.StatusCmd { _ = "STUB: not implemented"; return nil }

// ParseAuthCommand parses a redcon.Command to create an Auth instance and validates command arguments.
func ParseAuthCommand(cmd redcon.Command) (*Auth, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
