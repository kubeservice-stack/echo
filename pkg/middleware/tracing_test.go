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
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/propagation"
	oteltrace "go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/embedded"
)

func TestTracing(t *testing.T) {
	assert := assert.New(t)
	router := gin.New()
	router.Use(TracingFunc("echo"))
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "dongjiang")
	})

	router.GET("/tracing", func(c *gin.Context) {
		c.String(http.StatusOK, "tracing")
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(http.StatusOK, w.Code)
	assert.Equal(
		http.Header{
			"Content-Type": []string{"text/plain; charset=utf-8"},
		}, w.Header())
	assert.Equal("dongjiang", w.Body.String())

	req1 := httptest.NewRequest(http.MethodGet, "/tracing", nil)
	w1 := httptest.NewRecorder()
	router.ServeHTTP(w1, req1)
	t.Log(w1)
	fmt.Println(w1.Header())

}

type propagators struct {
	embedded.TracerProvider
}

func (p *propagators) Tracer(instrumentationName string, opts ...oteltrace.TracerOption) oteltrace.Tracer {
	return &tracer{}
}

type tracer struct {
	embedded.Tracer
}

func (t *tracer) Start(ctx context.Context, spanName string, opts ...oteltrace.SpanStartOption) (context.Context, oteltrace.Span) {
	return ctx, nil
}

type tracerProvider struct {
}

func (t *tracerProvider) Inject(ctx context.Context, carrier propagation.TextMapCarrier) {

}

func (t *tracerProvider) Extract(ctx context.Context, carrier propagation.TextMapCarrier) context.Context {
	return ctx
}

func (t *tracerProvider) Fields() []string {
	return []string{}
}

func TestWithPropagators(t *testing.T) {
	cfg := &traceConfig{}
	opt := WithPropagators(&tracerProvider{})
	opt(cfg)
}

func TestWithTracerProvider(t *testing.T) {
	cfg := &traceConfig{}
	opt := WithTracerProvider(&propagators{})
	opt(cfg)
}
