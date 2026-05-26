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
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/tidwall/redcon"
)

type Put struct {
	DMap  string
	Key   string
	Value []byte
	EX    float64
	PX    int64
	EXAT  float64
	PXAT  int64
	NX    bool
	XX    bool
}

func NewPut(dmap, key string, value []byte) *Put { _ = "STUB: not implemented"; return nil }

func (p *Put) SetEX(ex float64) *Put { _ = "STUB: not implemented"; return nil }

func (p *Put) SetPX(px int64) *Put { _ = "STUB: not implemented"; return nil }

func (p *Put) SetEXAT(exat float64) *Put { _ = "STUB: not implemented"; return nil }

func (p *Put) SetPXAT(pxat int64) *Put { _ = "STUB: not implemented"; return nil }

func (p *Put) SetNX() *Put { _ = "STUB: not implemented"; return nil }

func (p *Put) SetXX() *Put { _ = "STUB: not implemented"; return nil }

func (p *Put) Command(ctx context.Context) *redis.StatusCmd { _ = "STUB: not implemented"; return nil }

func ParsePutCommand(cmd redcon.Command) (*Put, error) { _ = "STUB: not implemented"; return nil, nil }

// DMap
// Key
// Value

type PutEntry struct {
	DMap  string
	Key   string
	Value []byte
}

func NewPutEntry(dmap, key string, value []byte) *PutEntry { _ = "STUB: not implemented"; return nil }

func (p *PutEntry) Command(ctx context.Context) *redis.StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func ParsePutEntryCommand(cmd redcon.Command) (*PutEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Get struct {
	DMap string
	Key  string
	Raw  bool
}

func NewGet(dmap, key string) *Get { _ = "STUB: not implemented"; return nil }

func (g *Get) SetRaw() *Get { _ = "STUB: not implemented"; return nil }

func (g *Get) Command(ctx context.Context) *redis.StringCmd { _ = "STUB: not implemented"; return nil }

func ParseGetCommand(cmd redcon.Command) (*Get, error) { _ = "STUB: not implemented"; return nil, nil }

type GetEntry struct {
	DMap    string
	Key     string
	Replica bool
}

func NewGetEntry(dmap, key string) *GetEntry { _ = "STUB: not implemented"; return nil }

func (g *GetEntry) SetReplica() *GetEntry { _ = "STUB: not implemented"; return nil }

func (g *GetEntry) Command(ctx context.Context) *redis.StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func ParseGetEntryCommand(cmd redcon.Command) (*GetEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DMap
// Key

type Del struct {
	DMap string
	Keys []string
}

func NewDel(dmap string, keys ...string) *Del { _ = "STUB: not implemented"; return nil }

func (d *Del) Command(ctx context.Context) *redis.IntCmd { _ = "STUB: not implemented"; return nil }

func ParseDelCommand(cmd redcon.Command) (*Del, error) { _ = "STUB: not implemented"; return nil, nil }

type DelEntry struct {
	Del     *Del
	Replica bool
}

func NewDelEntry(dmap, key string) *DelEntry { _ = "STUB: not implemented"; return nil }

func (d *DelEntry) SetReplica() *DelEntry { _ = "STUB: not implemented"; return nil }

func (d *DelEntry) Command(ctx context.Context) *redis.IntCmd {
	_ = "STUB: not implemented"
	return nil
}

func ParseDelEntryCommand(cmd redcon.Command) (*DelEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type PExpire struct {
	DMap         string
	Key          string
	Milliseconds time.Duration
}

func NewPExpire(dmap, key string, milliseconds time.Duration) *PExpire {
	_ = "STUB: not implemented"
	return nil
}

func (p *PExpire) Command(ctx context.Context) *redis.StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func ParsePExpireCommand(cmd redcon.Command) (*PExpire, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DMap
// Key

type Expire struct {
	DMap    string
	Key     string
	Seconds time.Duration
}

func NewExpire(dmap, key string, seconds time.Duration) *Expire {
	_ = "STUB: not implemented"
	return nil
}

func (e *Expire) Command(ctx context.Context) *redis.StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func ParseExpireCommand(cmd redcon.Command) (*Expire, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DMap
// Key

type Destroy struct {
	DMap  string
	Local bool
}

func NewDestroy(dmap string) *Destroy { _ = "STUB: not implemented"; return nil }

func (d *Destroy) SetLocal() *Destroy { _ = "STUB: not implemented"; return nil }

func (d *Destroy) Command(ctx context.Context) *redis.StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func ParseDestroyCommand(cmd redcon.Command) (*Destroy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Scan struct {
	PartID  uint64
	DMap    string
	Cursor  uint64
	Count   int
	Match   string
	Replica bool
}

func NewScan(partID uint64, dmap string, cursor uint64) *Scan {
	_ = "STUB: not implemented"
	return nil
}

func (s *Scan) SetMatch(match string) *Scan { _ = "STUB: not implemented"; return nil }

func (s *Scan) SetCount(count int) *Scan { _ = "STUB: not implemented"; return nil }

func (s *Scan) SetReplica() *Scan { _ = "STUB: not implemented"; return nil }

func (s *Scan) Command(ctx context.Context) *redis.ScanCmd { _ = "STUB: not implemented"; return nil }

const DefaultScanCount = 10

func ParseScanCommand(cmd redcon.Command) (*Scan, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DMap

type Incr struct {
	DMap  string
	Key   string
	Delta int
}

func NewIncr(dmap, key string, delta int) *Incr { _ = "STUB: not implemented"; return nil }

func (i *Incr) Command(ctx context.Context) *redis.IntCmd { _ = "STUB: not implemented"; return nil }

func ParseIncrCommand(cmd redcon.Command) (*Incr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Decr struct {
	*Incr
}

func NewDecr(dmap, key string, delta int) *Decr { _ = "STUB: not implemented"; return nil }

func (d *Decr) Command(ctx context.Context) *redis.IntCmd { _ = "STUB: not implemented"; return nil }

func ParseDecrCommand(cmd redcon.Command) (*Decr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type GetPut struct {
	DMap  string
	Key   string
	Value []byte
	Raw   bool
}

func NewGetPut(dmap, key string, value []byte) *GetPut { _ = "STUB: not implemented"; return nil }

func (g *GetPut) SetRaw() *GetPut { _ = "STUB: not implemented"; return nil }

func (g *GetPut) Command(ctx context.Context) *redis.StringCmd {
	_ = "STUB: not implemented"
	return nil
}

func ParseGetPutCommand(cmd redcon.Command) (*GetPut, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DMap
// Key
// Value

type IncrByFloat struct {
	DMap  string
	Key   string
	Delta float64
}

func NewIncrByFloat(dmap, key string, delta float64) *IncrByFloat {
	_ = "STUB: not implemented"
	return nil
}

func (i *IncrByFloat) Command(ctx context.Context) *redis.FloatCmd {
	_ = "STUB: not implemented"
	return nil
}

func ParseIncrByFloatCommand(cmd redcon.Command) (*IncrByFloat, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Lock struct {
	DMap     string
	Key      string
	Deadline float64
	EX       float64
	PX       int64
}

func NewLock(dmap, key string, deadline float64) *Lock { _ = "STUB: not implemented"; return nil }

func (l *Lock) SetEX(ex float64) *Lock { _ = "STUB: not implemented"; return nil }

func (l *Lock) SetPX(px int64) *Lock { _ = "STUB: not implemented"; return nil }

func (l *Lock) Command(ctx context.Context) *redis.StringCmd { _ = "STUB: not implemented"; return nil }

// Options

func ParseLockCommand(cmd redcon.Command) (*Lock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DMap
// Key
// Deadline

// EX or PX are optional.

type Unlock struct {
	DMap  string
	Key   string
	Token string
}

func NewUnlock(dmap, key, token string) *Unlock { _ = "STUB: not implemented"; return nil }

func (u *Unlock) Command(ctx context.Context) *redis.StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func ParseUnlockCommand(cmd redcon.Command) (*Unlock, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DMap
// Key
// Token

type LockLease struct {
	DMap    string
	Key     string
	Token   string
	Timeout float64
}

func NewLockLease(dmap, key, token string, timeout float64) *LockLease {
	_ = "STUB: not implemented"
	return nil
}

func (l *LockLease) Command(ctx context.Context) *redis.StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func ParseLockLeaseCommand(cmd redcon.Command) (*LockLease, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DMap
// Key
// Token
// Timeout

type PLockLease struct {
	DMap    string
	Key     string
	Token   string
	Timeout int64
}

func NewPLockLease(dmap, key, token string, timeout int64) *PLockLease {
	_ = "STUB: not implemented"
	return nil
}

func (p *PLockLease) Command(ctx context.Context) *redis.StatusCmd {
	_ = "STUB: not implemented"
	return nil
}

func ParsePLockLeaseCommand(cmd redcon.Command) (*PLockLease, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DMap
// Key
// Token
// Timeout
