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

// Exists check the ip exists
func (iPS *IPS) Exists(ip string) (exists bool) {
	if iPS.Mutex != nil {
		iPS.Mutex.RLock()
		defer iPS.Mutex.RUnlock()
	}
	currentIP := net.ParseIP(ip)
	for _, value := range iPS.IPList {
		if currentIP.Equal(value) {
			return true
		}
	}
	for _, ipNet := range iPS.IPNetList {
		if ipNet.Contains(currentIP) {
			return true
		}
	}
	return
}

// Add add ip to list
func (iPS *IPS) Add(ip string) (err error) {
	if iPS.Mutex != nil {
		iPS.Mutex.Lock()
		defer iPS.Mutex.Unlock()
	}
	// IPNet
	if strings.Index(ip, "/") != -1 {
		_, ipNet, err := net.ParseCIDR(ip)
		if err != nil {
			return err
		}
		iPS.IPNetList = append(iPS.IPNetList, ipNet)
	}
	value := net.ParseIP(ip)
	if value != nil {
		iPS.IPList = append(iPS.IPList, value)
	}
	return
}
