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

package dmap

func (s *Service) callCompactionOnFragment(f *fragment) bool {
	_ = "STUB: not implemented"
	return false
}

// Continue

// Break

func (s *Service) doCompaction(partID uint64) { _ = "STUB: not implemented"; return }

// Continue. This fragment belongs to a different data structure.

func (s *Service) triggerCompaction() { _ = "STUB: not implemented"; return }

// NumCPU returns the number of logical CPUs usable by the current process.
//
// The set of available CPUs is checked by querying the operating system
// at process startup. Changes to operating system CPU allocation after
// process startup are not reflected.

func (s *Service) compactionWorker() { _ = "STUB: not implemented"; return }
