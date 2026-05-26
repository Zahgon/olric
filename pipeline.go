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
	"errors"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	// ErrNotReady denotes that the Future instance you hold is not ready to read the response yet.
	ErrNotReady = errors.New("not ready yet")

	// ErrPipelineClosed denotes that the underlying pipeline is closed, and it's impossible to operate.
	ErrPipelineClosed = errors.New("pipeline is closed")

	// ErrPipelineExecuted denotes that Exec was already called on the underlying pipeline.
	ErrPipelineExecuted = errors.New("pipeline already executed")
)

// DMapPipeline implements a pipeline for the following methods of the DMap API:
//
// * Put
// * Get
// * Delete
// * Incr
// * Decr
// * GetPut
// * IncrByFloat
//
// DMapPipeline enables batch operations on DMap data.
type DMapPipeline struct {
	mtx          sync.Mutex
	dm           *ClusterDMap
	commands     map[uint64][]redis.Cmder
	result       map[uint64][]redis.Cmder
	ctx          context.Context
	cancel       context.CancelFunc
	closedCtx    context.Context // used to detect if the pipeline is closed / discarded
	closedCancel context.CancelFunc

	concurrency int // defaults to runtime.NumCPU()
}

func (dp *DMapPipeline) addCommand(key string, cmd redis.Cmder) (uint64, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// if there are no existing commands, get a new slice from the pool

// FuturePut is used to read the result of a pipelined Put command.
type FuturePut struct {
	dp        *DMapPipeline
	partID    uint64
	index     int
	ctx       context.Context
	closedCtx context.Context
}

// Result returns a response for the pipelined Put command.
func (f *FuturePut) Result() error {
	_ = "STUB: not implemented"
	// this select is separate from the one below on purpose, since select is non-deterministic if multiple
	// cases are available, and we need to guarantee this check first.
	return nil
}

// Put queues a Put command. The parameters are identical to the DMap.Put,
// but it returns FuturePut to read the batched response.
func (dp *DMapPipeline) Put(ctx context.Context, key string, value interface{}, options ...PutOption) (*FuturePut, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FutureGet is used to read result of a pipelined Get command.
type FutureGet struct {
	dp        *DMapPipeline
	partID    uint64
	index     int
	ctx       context.Context
	closedCtx context.Context
}

// Result returns a response for the pipelined Get command.
func (f *FutureGet) Result() (*GetResponse, error) {
	_ = "STUB: not implemented"
	// this select is separate from the one below on purpose, since select is non-deterministic if multiple
	// cases are available, and we need to guarantee this check first.
	return nil, nil
}

// Get queues a Get command. The parameters are identical to the DMap.Get,
// but it returns FutureGet to read the batched response.
func (dp *DMapPipeline) Get(ctx context.Context, key string) *FutureGet {
	_ = "STUB: not implemented"
	return nil
}

// FutureDelete is used to read the result of a pipelined Delete command.
type FutureDelete struct {
	dp        *DMapPipeline
	partID    uint64
	index     int
	ctx       context.Context
	closedCtx context.Context
}

// Result returns a response for the pipelined Delete command.
func (f *FutureDelete) Result() (int, error) {
	_ = "STUB: not implemented"
	// this select is separate from the one below on purpose, since select is non-deterministic if multiple
	// cases are available, and we need to guarantee this check first.
	return 0, nil
}

// Delete queues a Delete command. The parameters are identical to the DMap.Delete,
// but it returns FutureDelete to read the batched response.
func (dp *DMapPipeline) Delete(ctx context.Context, key string) *FutureDelete {
	_ = "STUB: not implemented"
	return nil
}

// FutureExpire is used to read the result of a pipelined Expire command.
type FutureExpire struct {
	dp        *DMapPipeline
	partID    uint64
	index     int
	ctx       context.Context
	closedCtx context.Context
}

// Result returns a response for the pipelined Expire command.
func (f *FutureExpire) Result() error {
	_ = "STUB: not implemented"
	// this select is separate from the one below on purpose, since select is non-deterministic if multiple
	// cases are available, and we need to guarantee this check first.
	return nil
}

// Expire queues an Expire command. The parameters are identical to the DMap.Expire,
// but it returns FutureExpire to read the batched response.
func (dp *DMapPipeline) Expire(ctx context.Context, key string, timeout time.Duration) (*FutureExpire, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FutureIncr is used to read the result of a pipelined Incr command.
type FutureIncr struct {
	dp        *DMapPipeline
	partID    uint64
	index     int
	ctx       context.Context
	closedCtx context.Context
}

// Result returns a response for the pipelined Incr command.
func (f *FutureIncr) Result() (int, error) {
	_ = "STUB: not implemented"
	// this select is separate from the one below on purpose, since select is non-deterministic if multiple
	// cases are available, and we need to guarantee this check first.
	return 0, nil
}

// Incr queues an Incr command. The parameters are identical to the DMap.Incr,
// but it returns FutureIncr to read the batched response.
func (dp *DMapPipeline) Incr(ctx context.Context, key string, delta int) (*FutureIncr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FutureDecr is used to read the result of a pipelined Decr command.
type FutureDecr struct {
	dp        *DMapPipeline
	partID    uint64
	index     int
	ctx       context.Context
	closedCtx context.Context
}

// Result returns a response for the pipelined Decr command.
func (f *FutureDecr) Result() (int, error) {
	_ = "STUB: not implemented"
	// this select is separate from the one below on purpose, since select is non-deterministic if multiple
	// cases are available, and we need to guarantee this check first.
	return 0, nil
}

// Decr queues a Decr command. The parameters are identical to the DMap.Decr,
// but it returns FutureDecr to read the batched response.
func (dp *DMapPipeline) Decr(ctx context.Context, key string, delta int) (*FutureDecr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FutureGetPut is used to read the result of a pipelined GetPut command.
type FutureGetPut struct {
	dp        *DMapPipeline
	partID    uint64
	index     int
	ctx       context.Context
	closedCtx context.Context
}

// Result returns a response for the pipelined GetPut command.
func (f *FutureGetPut) Result() (*GetResponse, error) {
	_ = "STUB: not implemented"
	// this select is separate from the one below on purpose, since select is non-deterministic if multiple
	// cases are available, and we need to guarantee this check first.
	return nil, nil
}

// This should be the first run.

// GetPut queues a GetPut command. The parameters are identical to the DMap.GetPut,
// but it returns FutureGetPut to read the batched response.
func (dp *DMapPipeline) GetPut(ctx context.Context, key string, value interface{}) (*FutureGetPut, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FutureIncrByFloat is used to read the result of a pipelined IncrByFloat command.
type FutureIncrByFloat struct {
	dp        *DMapPipeline
	partID    uint64
	index     int
	ctx       context.Context
	closedCtx context.Context
}

// Result returns a response for the pipelined IncrByFloat command.
func (f *FutureIncrByFloat) Result() (float64, error) {
	_ = "STUB: not implemented"
	// this select is separate from the one below on purpose, since select is non-deterministic if multiple
	// cases are available, and we need to guarantee this check first.
	return 0, nil
}

// IncrByFloat queues an IncrByFloat command. The parameters are identical to the DMap.IncrByFloat,
// but it returns FutureIncrByFloat to read the batched response.
func (dp *DMapPipeline) IncrByFloat(ctx context.Context, key string, delta float64) (*FutureIncrByFloat, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dp *DMapPipeline) execOnPartition(ctx context.Context, partID uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// There is no need to protect dp.commands map and its content.
// It's already filled before running Exec, and it's now a read-only
// data structure

// Exec executes all previously queued commands using one
// client-server roundtrip.
//
// Exec always returns list of commands and error of the first failed
// command if any.

// Exec executes all queued commands using one client-server roundtrip per partition.
func (dp *DMapPipeline) Exec(ctx context.Context) error {
	_ = "STUB: not implemented"
	// this select is separate from the one below on purpose, since select is non-deterministic if multiple
	// cases are available, and we need to guarantee this check first.
	return nil
}

// this checks to see if Exec has already run. While Exec should only be called once, it is possible that
// the user could call Exec multiple times. If we stored the result of errGr.Wait on the pipeline, we could
// return that error and make Exec idempotent.

// If execOnPartition returns an error, it will eventually stop
// all flush operation.

// Discard discards the pipelined commands and resets all internal states.
// A pipeline can be reused after calling Discard.
func (dp *DMapPipeline) Discard() error { _ = "STUB: not implemented"; return nil }

// return all command slices to the pool

// the deletes below are purposefully not combined with the loops above, as these are recognized and optimized
// by the compiler. https://go-review.googlesource.com/c/go/+/110055

// Close closes the pipeline and frees the allocated resources. You shouldn't try to
// reuse a closed pipeline.
func (dp *DMapPipeline) Close() {
	_ = "STUB: not implemented"

	// Pipeline is a mechanism to realise Redis Pipeline technique.
	//
	// Pipelining is a technique to extremely speed up processing by packing
	// operations to batches, send them at once to Redis and read a replies in a
	// singe step.
	// See https://redis.io/topics/pipelining
	//
	// Pay attention, that Pipeline is not a transaction, so you can get unexpected
	// results in case of big pipelines and small read/write timeouts.
	// Redis client has retransmission logic in case of timeouts, pipeline
	// can be retransmitted and commands can be executed more than once.
	return
}

func (dm *ClusterDMap) Pipeline(opts ...PipelineOption) (*DMapPipeline, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// initContexts sets up chained contexts for the pipeline. The base is closedCtx, which is closed either in
// Close or Discard. ctx is a child of closedCtx, as we want to cancel the pipeline if it is closed. It is
// canceled in Exec, and used to block FutureXXX.Result() calls until Exec has completed.
func (dp *DMapPipeline) initContexts() { _ = "STUB: not implemented"; return }

// This stores a slice of commands for each partition. There is a possibility that a single
// large slice could be allocated with an unusually large number of commands in a single pipeline that
// are very unbalanced across partitions, but that is unlikely to be a problem in practice.
//
// It does not store a pointer to the slice as recommended by staticcheck because that is harder to reason
// about, and a single allocation is not a big deal compared to the slices we're able to reuse.
// https://staticcheck.io/docs/checks#SA6002
// https://github.com/dominikh/go-tools/issues/1336#issuecomment-1331206290
var pipelineCmdPool = sync.Pool{
	New: func() interface{} {
		return make([]redis.Cmder, 0)
	},
}

func getPipelineCmdsFromPool() []redis.Cmder { _ = "STUB: not implemented"; return nil }

func putPipelineCmdsIntoPool(cmds []redis.Cmder) {
	_ = "STUB: not implemented"
	// remove references to underlying commands so they can be GCed
	return
}
