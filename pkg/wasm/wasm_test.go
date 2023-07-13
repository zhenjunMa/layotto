/*
 * Copyright 2021 Layotto Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package wasm

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"mosn.io/layotto/pkg/filter/stream/common/http"
)

type MockEndpoint struct {
}

func (m *MockEndpoint) Handle(ctx context.Context, params http.ParamsScanner) (map[string]interface{}, error) {
	return nil, nil
}

// TestWasm test AddEndpoint and GetEndpoint.
func TestWasm(t *testing.T) {
	// get singleton Wasm
	wasm := GetDefault()
	// reset before test
	wasm.AddEndpoint("install", nil)

	endpoint, ok := wasm.GetEndpoint("install")
	assert.True(t, ok)
	assert.Nil(t, endpoint)

	wasm.AddEndpoint("", nil)
	endpoint, ok = wasm.GetEndpoint("")
	assert.True(t, ok)
	assert.Nil(t, endpoint)

	ep := &MockEndpoint{}
	wasm.AddEndpoint("install", ep)
	// reset
	defer func() {
		wasm.AddEndpoint("install", nil)
	}()
	endpoint, ok = wasm.GetEndpoint("install")
	assert.True(t, ok)
	assert.Equal(t, endpoint, ep)
}
