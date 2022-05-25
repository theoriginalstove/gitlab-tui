package main

import (
	"net/http"
	"os"

	"github.com/levenlabs/go-llog"
)

type GlabTui struct {
	token   string
	handler *http.Client
}

// Configued returns a pointer to a GlabTui instance with a token set
// from the user's local environment
func Configured() *GlabTui {
	g := &GlabTui{}
	g.setGitlabToken()
	return g
}

func (g *GlabTui) setGitlabToken() {
	token := os.Getenv("GITLAB_TUI_TOKEN")
	if token == "" {
		llog.Fatal("GITLAB_TUI_TOKEN is not set")
	}
	g.token = token
}

///////////////////////////////////////////////////////////////////////////////

type gitlabMergeRequestParams struct {
}

type gitlabMergeRequest struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Description string `json:"description"`
	Visibility  string
}

type gitlabMergeRequestRes struct {
}

// TODO: Create query parameter/creator to parse from

func (g *glabTui) getMergeRequests() {
	// TODO: Get the group info
}
