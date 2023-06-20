package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/getMany?size=20&page=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusOK)

	var data []map[string]any
	if err = json.Unmarshal(resp.Body.Bytes(), &data); err != nil {
		t.Fatal(err)
	}

	if len(data) != 20 {
		t.Error("length not equal 20!")
	}
}
