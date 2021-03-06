// Copyright 2019 Altinity Ltd and/or its affiliates. All rights reserved.
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

package models

import (
	"encoding/hex"
	"io"
	"fmt"
	"math/rand"
	"time"
)

// randomString generates random string
func randomString() string {
	b := make([]byte, 3)
	rand.New(rand.NewSource(time.Now().UnixNano())).Read(b)
	return hex.EncodeToString(b)
}

// includeNonEmpty inserts (and overwrites) data into map object using specified key, if not empty value provided
func includeNonEmpty(dst map[string]string, key, src string) {
	// Do not include empty value
	if src == "" {
		return
	}

	// Include (and overwrite) value by specified key
	dst[key] = src

	return
}

// fprintf suppresses warning for unused returns of fmt.Fprintf()
func fprintf(w io.Writer, format string, a ...interface{}) {
	_, _ = fmt.Fprintf(w, format, a...)
}
