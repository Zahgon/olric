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
	"net"
)

// The following functions are mostly extracted from Serf. See setupAgent function in cmd/serf/command/agent/command.go
// Thanks for the extraordinary software.
//
// Source: https://github.com/hashicorp/serf/blob/master/cmd/serf/command/agent/command.go#L204

func addrParts(address string) (string, int, error) {
	_ = "STUB: not implemented"
	// Get the address
	return "", 0, nil
}

func getBindIPFromNetworkInterface(addrs []net.Addr) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Waiting for https://github.com/golang/go/issues/5395 to use IPNet only

// Skip self-assigned IPs

func getBindIP(ifname, address string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Check if we have an interface

// If there is no bind IP, pick an address

// If there is a bind IP, ensure it is available

// if we're not bound to a specific IP, let's use a suitable private IP address.

// if we could not find a private address, we need to expand our search to a public
// ip address

// SetupNetworkConfig tries to find an appropriate bindIP to bind and propagate.
func (c *Config) SetupNetworkConfig() (err error) { _ = "STUB: not implemented"; return nil }
