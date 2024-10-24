package metrics

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kubeservice-stack/echo/pkg/routers"
	"github.com/stretchr/testify/assert"
)

func TestMetricsRoute(t *testing.T) {
	assert := assert.New(t)
	r := gin.Default()
	router.Router(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/metrics", nil)
	req.Host = "127.0.0.1:9445"
	r.ServeHTTP(w, req)

	assert.Equal(200, w.Code)
	assert.Greater(len(w.Body.String()), 1000)
}
