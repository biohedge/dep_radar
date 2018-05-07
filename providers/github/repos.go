package github

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/stamm/dep_radar"
)

type reposResponse struct {
	FullName string `json:"full_name"`
}

// GetAllOrgRepos get repos
func (g *Provider) GetAllOrgRepos(ctx context.Context, org string) ([]dep_radar.Pkg, error) {
	var (
		resultRepos []dep_radar.Pkg
	)

	repos, err := g.getRepos(ctx, org)
	if err != nil {
		return resultRepos, err
	}
	for _, repo := range repos {
		resultRepos = append(resultRepos, dep_radar.Pkg(Prefix+"/"+repo.FullName))
	}
	return resultRepos, nil
}

func (g *Provider) getRepos(ctx context.Context, org string) ([]reposResponse, error) {
	var repos []reposResponse
	url := g.getOrgReposURL(org)
	reposResponse, err := g.client.Get(ctx, url)
	if err != nil {
		return repos, err
	}
	err = json.Unmarshal(reposResponse, &repos)

	return repos, err
}

func (g *Provider) getOrgReposURL(org string) string {
	return fmt.Sprintf("https://api.github.com/orgs/%s/repos", org)
}