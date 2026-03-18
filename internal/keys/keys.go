package keys

import (
	"log"
	"vx/config"
	"vx/internal/keyrings"
	"vx/pkg/gutils"
)

func KeySrc() config.RepoMode {
	if config.RepoModeActive() {
		repoCfg, err := config.LoadRepoConfig()
		if err == nil {
			return repoCfg
		}
		log.Println("Error loading repo config:", err)
	}

	owner, _, _ := gutils.GetRepo()
	if owner != "" {
		if repoCfg, ok := resolveContextByOwner(owner); ok {
			return repoCfg
		}
	}

	if repoCfg, ok := resolveActiveContext(); ok {
		return repoCfg
	}

	return config.RepoMode{}
}

func resolveContextByOwner(owner string) (config.RepoMode, bool) {
	ctxCfg, err := config.LoadContextConfig()
	if err != nil {
		return config.RepoMode{}, false
	}
	for name, ctx := range ctxCfg.Contexts {
		if ctx.GithubUser == owner {
			return keysFromContext(name)
		}
	}
	return config.RepoMode{}, false
}

func resolveActiveContext() (config.RepoMode, bool) {
	ctxCfg, err := config.LoadContextConfig()
	if err != nil || ctxCfg.Active == "" {
		return config.RepoMode{}, false
	}
	return keysFromContext(ctxCfg.Active)
}

func keysFromContext(name string) (config.RepoMode, bool) {
	githubKey, err := keyrings.GetSecret(name, "github_key")
	if err != nil || githubKey == "" {
		return config.RepoMode{}, false
	}
	openaiKey, _ := keyrings.GetSecret(name, "openai_key")
	_, repo, repoURL := gutils.GetRepo()
	return config.RepoMode{
		Github_key: githubKey,
		Openai_key: openaiKey,
		Repo:       repo,
		RepoURL:    repoURL,
	}, true
}
