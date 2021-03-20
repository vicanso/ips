// Copyright 2019 tree xie
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

package ips

import (
	"net"
	"strings"
	"sync"
)

type (
	// IPS ip list
	IPS struct {
		Mutex     *sync.RWMutex
		IPList    []net.IP
		IPNetList []*net.IPNet
	}
)

// New return a new ips with mutex
func New() *IPS {
	ips := NewWithoutMutex()
	ips.Mutex = &sync.RWMutex{}
	return ips
}

// NewWithoutMutex return a new ips without mutex
func NewWithoutMutex() *IPS {
	ips := &IPS{}
	ips.reset()
	return ips
}

// Contains returns true if contains any of ip
func (ips *IPS) Contains(ipList ...string) bool {
	if ips.Mutex != nil {
		ips.Mutex.RLock()
		defer ips.Mutex.RUnlock()
	}
	for _, ip := range ipList {
		currentIP := net.ParseIP(ip)
		for _, value := range ips.IPList {
			if currentIP.Equal(value) {
				return true
			}
		}
		for _, ipNet := range ips.IPNetList {
			if ipNet.Contains(currentIP) {
				return true
			}
		}
	}
	return false
}

func (ips *IPS) reset() {
	ips.IPList = make([]net.IP, 0)
	ips.IPNetList = make([]*net.IPNet, 0)
}

// Reset cleans all ip data
func (ips *IPS) Reset() {
	if ips.Mutex != nil {
		ips.Mutex.Lock()
		defer ips.Mutex.Unlock()
	}
	ips.reset()
}

// Replace cleans all ip data and add new ip data
func (ips *IPS) Replace(ipList ...string) (err error) {
	if ips.Mutex != nil {
		ips.Mutex.Lock()
		defer ips.Mutex.Unlock()
	}
	ips.reset()
	return ips.add(ipList)
}

func (ips *IPS) add(ipList []string) (err error) {
	for _, ip := range ipList {
		// IPNet
		if strings.Contains(ip, "/") {
			_, ipNet, err := net.ParseCIDR(ip)
			if err != nil {
				return err
			}
			ips.IPNetList = append(ips.IPNetList, ipNet)
		}
		value := net.ParseIP(ip)
		if value != nil {
			ips.IPList = append(ips.IPList, value)
		}
	}
	return
}

// Add adds ip to list
func (ips *IPS) Add(ipList ...string) (err error) {
	if ips.Mutex != nil {
		ips.Mutex.Lock()
		defer ips.Mutex.Unlock()
	}
	return ips.add(ipList)
}

// Strings retruns string slice of ips
func (ips *IPS) Strings() []string {
	if ips.Mutex != nil {
		ips.Mutex.RLock()
		defer ips.Mutex.RUnlock()
	}
	arr := make([]string, len(ips.IPList)+len(ips.IPNetList))
	index := 0
	for _, value := range ips.IPList {
		arr[index] = value.String()
		index++
	}
	for _, ipNet := range ips.IPNetList {
		arr[index] = ipNet.String()
		index++
	}
	return arr
}
