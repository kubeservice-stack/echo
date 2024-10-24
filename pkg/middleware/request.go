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

package middleware

import (
	"context"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

const (
	REQUESTINFO       = "REQUESTINFO"
	REQUESTINFOWEIGHT = 120
)

func RequestInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqCxt := c.Request.Context()
		//来源请求ID
		forwardRequestId := c.Request.Header.Get("uniqid")
		reqCxt = context.WithValue(reqCxt, "forwardRequestId", forwardRequestId)
		//请求ID
		requestId := c.Request.Header.Get("requestId")

		if requestId == "" {
			requestId = uuid.NewV4().String()
		}

		reqCxt = context.WithValue(reqCxt, "requestId", requestId)
		reqCxt = context.WithValue(reqCxt, "clientAddress", c.Request.RemoteAddr)
		if http.LocalAddrContextKey != nil && reqCxt.Value(http.LocalAddrContextKey) != nil {
			reqCxt = context.WithValue(reqCxt, "serverAddress", reqCxt.Value(http.LocalAddrContextKey).(*net.TCPAddr).String())
		}
		c.Request = c.Request.WithContext(reqCxt)

		// 处理请求
		c.Next()
	}
}

func init() {
	Register(&Instance{F: RequestInfo, Weight: REQUESTINFOWEIGHT, Name: REQUESTINFO})
}
