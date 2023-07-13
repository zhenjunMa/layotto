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

package client

// Tool for `Recover go func()`, copied from `mosn.io/pkg/utils/goroutine.go`

import (
	"fmt"
	"io"
	"os"
	"runtime/debug"
)

var recoverLogger = defaultRecoverLogger

// RegisterRecoverLogger replace the log handler when go with recover catch a panic
// notice the replaced handler should be simple.
// if the handler panic, the recover handler will be failed.
func RegisterRecoverLogger(f func(w io.Writer, r interface{})) {
	recoverLogger = f
}

func defaultRecoverLogger(w io.Writer, r interface{}) {
	fmt.Fprintf(w, "goroutine panic: %v\n%s\n", r, string(debug.Stack()))
}

// GoWithRecover wraps a `go func()` with recover()
func GoWithRecover(handler func(), recoverHandler func(r interface{})) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				recoverLogger(os.Stderr, r)
				if recoverHandler != nil {
					go func() {
						defer func() {
							if p := recover(); p != nil {
								recoverLogger(os.Stderr, p)
							}
						}()
						recoverHandler(r)
					}()
				}
			}
		}()
		handler()
	}()
}
