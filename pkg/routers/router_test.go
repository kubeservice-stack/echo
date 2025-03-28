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

package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_Register(t *testing.T) {
	assert := assert.New(t)
	router := gin.New()
	aaa := func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	}
	Register("aaa", "/", "aaa", http.MethodGet, aaa)
	Router(router)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/bbb", nil)
	router.ServeHTTP(w, req)

	assert.Equal(404, w.Code)
	assert.Len(w.Body.String(), 27)
}
