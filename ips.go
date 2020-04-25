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

// New create a new ips
func New() *IPS {
	return &IPS{
		IPList:    make([]net.IP, 0),
		IPNetList: make([]*net.IPNet, 0),
	}
}

// Contains contains the ip
func (ips *IPS) Contains(ip string) bool {
	if ips.Mutex != nil {
		ips.Mutex.RLock()
		defer ips.Mutex.RUnlock()
	}
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
	return false
}

// Add add ip to list
func (ips *IPS) Add(ip string) (err error) {
	if ips.Mutex != nil {
		ips.Mutex.Lock()
		defer ips.Mutex.Unlock()
	}
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
	return
}
