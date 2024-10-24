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
)

var defaultRouters = map[string]*HandlerService{
	"metrics": {
		HandleName: "metrics",
		Group:      "/",
		Path:       "/metrics",
		Method:     http.MethodGet,
		Host:       []string{"*"},
	},

	"healthz": {
		HandleName: "healthz",
		Group:      "/",
		Path:       "/healthz",
		Method:     http.MethodGet,
		Host:       []string{"*"},
	},

	"favicon.ico": {
		HandleName: "favicon.ico",
		Group:      "/",
		Path:       "/favicon.ico",
		Method:     http.MethodGet,
		Host:       []string{"*"},
	},
}
