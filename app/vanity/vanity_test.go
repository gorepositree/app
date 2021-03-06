// Copyright Jeevanandam M. (https://github.com/jeevatkm, jeeva@myjeeva.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vanity

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"thumbai/app/models"

	"github.com/stretchr/testify/assert"
)

func TestTreeAddAndLookup(t *testing.T) {
	Thumbai = &vanities{RWMutex: sync.RWMutex{}, Hosts: make(map[string]*vanityHost)}

	for _, importPath := range []string{
		"/aah",
		"/cli",
		"/cache/provider/redis",
	} {
		err := Add2Tree(&models.VanityPackage{Host: "aahframe.work", Path: importPath, Repo: "https://github.com/go-aah/aah.git"})
		assert.Nil(t, err)
	}

	// Lookup
	for _, importPath := range []string{
		"/cache/provider/redis",
		"/",
		"/cli",
		"/cli/aah",
		"/aah/vfs",
		"/aah/cache",
		"/unknown",
	} {
		vh := Thumbai.Lookup("aahframe.work")
		assert.NotNil(t, vh)
		r := vh.Lookup(importPath)
		if r != nil {
			assert.True(t, strings.HasPrefix(importPath, r.Path))
		}
	}
}

func testdataBaseDir() string {
	wd, _ := os.Getwd()
	if idx := strings.Index(wd, ".testdata"); idx > 0 {
		wd = wd[:idx]
	}
	return filepath.Join(wd, ".testdata")
}
