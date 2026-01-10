package endpoints

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	router "github.com/bsv-blockchain/block-headers-service/transports/http/endpoints/routes"
)

// mockMiddleware implements router.APIMiddleware for testing
type mockMiddleware struct {
	called bool
}

func (m *mockMiddleware) ApplyToAPI(c *gin.Context) {
	m.called = true
}

func TestToHandlers_EmptySlice(t *testing.T) {
	result := toHandlers()

	assert.NotNil(t, result)
	assert.Empty(t, result)
}

func TestToHandlers_SingleMiddleware(t *testing.T) {
	middleware := &mockMiddleware{}

	result := toHandlers(middleware)

	assert.Len(t, result, 1)
	assert.NotNil(t, result[0])
}

func TestToHandlers_MultipleMiddlewares(t *testing.T) {
	middlewares := []router.APIMiddleware{
		&mockMiddleware{},
		&mockMiddleware{},
		&mockMiddleware{},
	}

	result := toHandlers(middlewares[0], middlewares[1], middlewares[2])

	assert.Len(t, result, 3)
	for i, handler := range result {
		assert.NotNil(t, handler, "handler %d should not be nil", i)
	}
}

func TestToHandlers_HandlerCallsMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	middleware := &mockMiddleware{}
	handlers := toHandlers(middleware)

	assert.Len(t, handlers, 1)

	// Create a test context and call the handler
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	handlers[0](c)

	assert.True(t, middleware.called, "middleware should have been called")
}

func TestToHandlers_PreservesOrder(t *testing.T) {
	gin.SetMode(gin.TestMode)

	var callOrder []int
	middleware1 := router.APIMiddlewareFunc(func(c *gin.Context) {
		callOrder = append(callOrder, 1)
	})
	middleware2 := router.APIMiddlewareFunc(func(c *gin.Context) {
		callOrder = append(callOrder, 2)
	})
	middleware3 := router.APIMiddlewareFunc(func(c *gin.Context) {
		callOrder = append(callOrder, 3)
	})

	handlers := toHandlers(middleware1, middleware2, middleware3)

	assert.Len(t, handlers, 3)

	// Call each handler in order
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	for _, h := range handlers {
		h(c)
	}

	assert.Equal(t, []int{1, 2, 3}, callOrder, "middlewares should be called in order")
}
