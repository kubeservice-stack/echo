/*
Copyright 2024 The KubeService-Stack Authors.

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

package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionInfo(t *testing.T) {
	assert := assert.New(t)
	info := Info()
	assert.Equal("(version=, branch=, revision=unknown)", info)
}

func TestVersionBuildContext(t *testing.T) {
	assert := assert.New(t)
	info := BuildContext()
	assert.NotEmpty(info)
}

func TestVersionPrint(t *testing.T) {
	assert := assert.New(t)
	info := Print("echo")
	assert.NotEmpty(info)
}
