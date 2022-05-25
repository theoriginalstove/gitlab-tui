package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/levenlabs/go-llog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetGitlabToken(t *testing.T) {
	// os.setenv of a token
	err := os.Setenv("GITLAB_TUI_TOKEN", "supersecrettoken")
	require.NoError(t, err, llog.ErrKV(err))
	// create instance of glab
	testInstance := &GlabTui{}
	testInstance.setGitlabToken()
	assert.Equal(t, "supersecrettoken", testInstance.token)

}

func TestGetMergeRequests(t *testing.T) {
	// mock http server and results
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "api/v4/merge_requests" {
			t.Errorf("Expected to request '/api/v4/merge_requests', got %s", r.URL.Path)
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected to Accept 'applcation/json', got %s", r.Header.Get("Accept"))
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[
        ]`))
	}))
	defer srv.Close()
}
