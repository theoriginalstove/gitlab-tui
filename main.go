package main

import (
	"os"

	"github.com/levenlabs/go-llog"
)

func main() {
}

type glab struct {
	token string
}

// Configued returns a pointer to a glab instance with a token set
// from the user's local environment
func Configured() *glab {
	g := &glab{}
	g.setGitlabToken()
	return g
}

func (g *glab) setGitlabToken() {
	token := os.Getenv("GITLAB_TUI_TOKEN")
	if token == "" {
		llog.Fatal("GITLAB_TUI_TOKEN is not set")
	}
}
