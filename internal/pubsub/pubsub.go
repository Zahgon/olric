// The MIT License (MIT)
//
// Copyright (c) 2016 Josh Baker
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package pubsub

import (
	"sync"

	"github.com/tidwall/btree"
	"github.com/tidwall/redcon"
)

// PubSub is a Redis compatible pub/sub server
type PubSub struct {
	mu     sync.RWMutex
	nextid uint64
	initd  bool
	chans  *btree.BTree
	conns  map[redcon.Conn]*pubSubConn

	// callbacks
	unsubscribeCallback  func()
	punsubscribeCallback func()
}

// Subscribe a connection to PubSub
func (ps *PubSub) Subscribe(conn redcon.Conn, channel string) { _ = "STUB: not implemented"; return }

// Psubscribe a connection to PubSub
func (ps *PubSub) Psubscribe(conn redcon.Conn, channel string) { _ = "STUB: not implemented"; return }

// Publish a message to subscribers
func (ps *PubSub) Publish(channel, message string) int { _ = "STUB: not implemented"; return 0 }

// write messages to all clients that are subscribed on the channel

// match on and write all psubscribe clients

type pubSubConn struct {
	id      uint64
	mu      sync.Mutex
	conn    redcon.Conn
	dconn   redcon.DetachedConn
	entries map[*pubSubEntry]bool
}

type pubSubEntry struct {
	pattern bool
	sconn   *pubSubConn
	channel string
}

func (sconn *pubSubConn) writeMessage(pat bool, pchan, channel, msg string) {
	_ = "STUB: not implemented"
	return
}

// bgrunner runs in the background and reads incoming commands from the
// detached client.
func (sconn *pubSubConn) bgrunner(ps *PubSub) {
	_ = "STUB: not implemented"

	// client connection has ended, disconnect from the PubSub instances
	// and close the network connection.
	return
}

// byEntry is a "less" function that sorts the entries in a btree. The tree
// is sorted be (pattern, channel, conn.id). All pattern=true entries are at
// the end (right) of the tree.
func byEntry(a, b interface{}) bool { _ = "STUB: not implemented"; return false }

func (ps *PubSub) subscribe(conn redcon.Conn, pattern bool, channel string) {
	_ = "STUB: not implemented"
	return
}

// initialize the PubSub instance

// fetch the pubSubConn

// initialize a new pubSubConn, which runs on a detached connection,
// and attach it to the PubSub channels/conn btree

// add an entry to the pubsub btree

// send a message to the client

// start the background client operation

func (ps *PubSub) unsubscribe(conn redcon.Conn, pattern, all bool, channel string) {
	_ = "STUB: not implemented"
	return
}

// fetch the pubSubConn. This must exist

// unsubscribe from all (p)subscribe entries

// unsubscribe single channel from (p)subscribe.

func (ps *PubSub) Channels() []string { _ = "STUB: not implemented"; return nil }

func (ps *PubSub) ChannelsWithPatterns(pattern string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (ps *PubSub) Numpat() int { _ = "STUB: not implemented"; return 0 }

func (ps *PubSub) Numsub(channel string) int { _ = "STUB: not implemented"; return 0 }
