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
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kubeservice-stack/common/pkg/errno"
	"github.com/kubeservice-stack/common/pkg/logger"
	"github.com/kubeservice-stack/echo/pkg/middleware"
	"github.com/kubeservice-stack/echo/pkg/response"
)

var log = logger.GetLogger("pkg/router", "router")

// handername - hander func
var handlerAdapter = make(map[string]gin.HandlerFunc)

func Register(name string, h gin.HandlerFunc) {
	if handlerAdapter == nil {
		panic("gin.Handler: Register adapter is nil")
	}
	if _, ok := handlerAdapter[name]; ok {
		panic("gin.Handler: Register called twice for adapter :" + name)
	}
	handlerAdapter[name] = h
}

func UnRegister(name string) {
	delete(handlerAdapter, name)
}

func FullRegisters() map[string]gin.HandlerFunc {
	return handlerAdapter
}

// 配置信息
type HandlerService struct {
	HandleName string   // handle name
	Group      string   // default group "/"
	Path       string   // domain path
	Method     string   // http.Method
	Host       []string // Host
}

// Router 路由规则
func Router(r *gin.Engine) {
	for _, mid := range middleware.AllMiddlewarePlugins() {
		log.Info("use gin middleware", logger.String("name", mid.Name))
		r.Use(mid.F())
	}

	v := r.Group("/")
	{
		v.Handle(http.MethodGet, "/*anypath", RootHander)
		v.Handle(http.MethodHead, "/*anypath", RootHander)
		v.Handle(http.MethodPut, "/*anypath", RootHander)
		v.Handle(http.MethodOptions, "/*anypath", RootHander)
	}
}

func RootHander(ctx *gin.Context) {
	path := ctx.Request.URL.Path
	if len(ctx.Request.URL.RawPath) > 0 {
		path = ctx.Request.URL.RawPath
	}

	if _, ok := handlerAdapter[strings.Trim(path, "/")]; ok {
		handlerAdapter[strings.Trim(path, "/")](ctx)
		return
	}

	response.JSON(ctx, errno.NotFound, nil)
}
