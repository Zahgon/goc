/*
 Copyright 2020 Qiniu Cloud (qiniu.com)

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package cover

import (
	"errors"
	"sync"
)

var ErrServiceAlreadyRegistered = errors.New("service already registered")

// Store persistents the registered service information
type Store interface {
	// Add adds the given service to store
	Add(s ServiceUnderTest) error

	// Get returns the registered service information with the given service's name
	Get(name string) []string

	// Get returns all the registered service information as a map
	GetAll() map[string][]string

	// Init cleanup all the registered service information
	Init() error

	// Set stores the services information into internal state
	Set(services map[string][]string) error

	// Remove the service from the store by address
	Remove(addr string) error
}

// fileStore holds the registered services into memory and persistent to a local file
type fileStore struct {
	mu             sync.RWMutex
	persistentFile string

	memoryStore Store
}

// NewFileStore creates a store using local file
func NewFileStore(persistenceFile string) (store Store, err error) {
	_ = "STUB: not implemented"
	return *new(Store), nil
}

// Add adds the given service to file Store
func (l *fileStore) Add(s ServiceUnderTest) error { _ = "STUB: not implemented"; return nil }

// persistent to local store

// Get returns the registered service information with the given name
func (l *fileStore) Get(name string) []string { _ = "STUB: not implemented"; return nil }

// Get returns all the registered service information
func (l *fileStore) GetAll() map[string][]string { _ = "STUB: not implemented"; return nil }

// Remove the service from the memory store and the file store
func (l *fileStore) Remove(addr string) error { _ = "STUB: not implemented"; return nil }

// Init cleanup all the registered service information
// and the local persistent file
func (l *fileStore) Init() error { _ = "STUB: not implemented"; return nil }

// load all registered service from file to memory
func (l *fileStore) load() error { _ = "STUB: not implemented"; return nil }

// TODO: use regex

// set information to memory

func (l *fileStore) Set(services map[string][]string) error { _ = "STUB: not implemented"; return nil }

// no error will return from memorystore.set

func (l *fileStore) appendToFile(s ServiceUnderTest) error { _ = "STUB: not implemented"; return nil }

func format(s ServiceUnderTest) string { _ = "STUB: not implemented"; return "" }

func split(r rune) bool { _ = "STUB: not implemented"; return false }

func syncToFile(persistentFile string, services map[string][]string) error {
	_ = "STUB: not implemented"
	return nil
}

// memoryStore holds the registered services only into memory
type memoryStore struct {
	mu          sync.RWMutex
	servicesMap map[string][]string
}

// NewMemoryStore creates a memory store
func NewMemoryStore() Store { _ = "STUB: not implemented"; return *new(Store) }

// Add adds the given service to MemoryStore
func (l *memoryStore) Add(s ServiceUnderTest) error { _ = "STUB: not implemented"; return nil }

// load to memory

// Get returns the registered service information with the given name
func (l *memoryStore) Get(name string) []string { _ = "STUB: not implemented"; return nil }

// Get returns all the registered service information
func (l *memoryStore) GetAll() map[string][]string { _ = "STUB: not implemented"; return nil }

// Init cleanup all the registered service information
// and the local persistent file
func (l *memoryStore) Init() error { _ = "STUB: not implemented"; return nil }

func (l *memoryStore) Set(services map[string][]string) error {
	_ = "STUB: not implemented"
	return nil
}

// Remove one service from the memory store
// if service is not fount, return "no service found" error
func (l *memoryStore) Remove(removeAddr string) error { _ = "STUB: not implemented"; return nil }

// if no services left, remove by name
