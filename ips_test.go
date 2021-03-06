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
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPS(t *testing.T) {
	assert := assert.New(t)
	ips := New()
	ips.Mutex = new(sync.RWMutex)
	_ = ips.Add("12.12.12.12")
	_ = ips.Add("12.12.12.13", "192.168.1.1/24")
	assert.True(ips.Contains("12.12.12.12"))
	assert.True(ips.Contains("12.12.12.13"))
	assert.False(ips.Contains("12.12.12.14"))

	assert.True(ips.Contains("192.168.1.1"))
	assert.True(ips.Contains("192.168.1.2"))
	assert.False(ips.Contains("192.168.2.1"))

	assert.Equal([]string{
		"12.12.12.12",
		"12.12.12.13",
		"192.168.1.0/24",
	}, ips.Strings())

	assert.Nil(ips.Replace("1.1.1.1"))
	assert.Equal([]string{
		"1.1.1.1",
	}, ips.Strings())

	ips.Reset()
	assert.Empty(ips.Strings())
}
