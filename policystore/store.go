// Copyright (c) 2018 Tigera, Inc. All rights reserved.

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

package policystore

import (
	"sync"

	"github.com/projectcalico/app-policy/proto"

	"github.com/prometheus/common/log"
)

// PolicyStore is a data store that holds Calico policy information.
type PolicyStore struct {
	// The RWMutex protects the entire contents of the PolicyStore. No one should read from or write to the PolicyStore
	// without acquiring the corresponding lock.
	// Helper methods Write() and Read() encapsulate the correct locking logic.
	RWMutex sync.RWMutex

	PolicyByID  map[proto.PolicyID]*proto.Policy
	ProfileByID map[proto.ProfileID]*proto.Profile
	IPSetByID   map[string]IPSet
	Endpoint    *proto.WorkloadEndpoint
}

func NewPolicyStore() *PolicyStore {
	return &PolicyStore{
		RWMutex:     sync.RWMutex{},
		IPSetByID:   make(map[string]IPSet),
		ProfileByID: make(map[proto.ProfileID]*proto.Profile),
		PolicyByID:  make(map[proto.PolicyID]*proto.Policy)}
}

// Write to/update the PolicyStore, handling locking logic.
// writeFn is the logic that actually does the update.
func (s *PolicyStore) Write(writeFn func(store *PolicyStore)) {
	// TODO (spikecurtis) create a correlator that can be tracked for logging.
	log.Debug("About to write lock PolicyStore")
	s.RWMutex.Lock()
	log.Debug("PolicyStore write locked")
	writeFn(s)
	s.RWMutex.Unlock()
	log.Debug("PolicyStore write unlocked")
}

// Read the PolicyStore, handling locking logic.
// readFn is the logic that actually does the reading.
func (s *PolicyStore) Read(readFn func(store *PolicyStore)) {
	log.Debug("About to read lock PolicyStore")
	s.RWMutex.RLock()
	log.Debug("PolicyStore read locked")
	readFn(s)
	s.RWMutex.RUnlock()
	log.Debug("PolicyStore read unlocked")
}
