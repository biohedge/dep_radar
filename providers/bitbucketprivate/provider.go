package bitbucketprivate

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/stamm/dep_radar"
)

var (
	_ dep_radar.ITagGetter = &Provider{}
	_ dep_radar.IProvider  = &Provider{}
)

// TagsResponse response from bitbucket for tags
type TagsResponse struct {
	IsLastPage    bool        `json:"isLastPage"`
	Values        []TagsValue `json:"values"`
	NextPageStart int         `json:"nextPageStart"`
}

// TagsValue sub response from bibucket for tags
type TagsValue struct {
	ID           string `json:"id"`
	LatestCommit string `json:"latestCommit"`
}

// Repo struct for repo
type Repo struct {
	Project     string
	Name        string
	Slug        string
	HashVersion string
	Version     string
	Deps        map[string]Repo
	Tags        []string
	Branch      string
}

// FileOpts options for getting file
type FileOpts struct {
	Project    string
	RepoName   string
	Filename   string
	BranchName string
}

// Provider for bitbucket
type Provider struct {
	project    string
	httpClient dep_radar.IWebClient
	gitDomain  string
	goGetURL   string
	apiURL     string
	muMap      sync.RWMutex
	mapProject map[dep_radar.Pkg]string
}

// Options for bitbucket
type Options struct {
	URL      string
	User     string
	Password string
}

// New creates new instance of provider
func New(httpClient dep_radar.IWebClient, gitDomain, goGetURL, apiURL string) *Provider {
	return &Provider{
		httpClient: httpClient,
		goGetURL:   goGetURL,
		gitDomain:  gitDomain,
		apiURL:     apiURL,
		mapProject: make(map[dep_radar.Pkg]string),
	}
}

// File gets file
func (p *Provider) File(ctx context.Context, pkg dep_radar.Pkg, branch, name string) ([]byte, error) {
	project, err := p.getProject(ctx, pkg)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/projects/%s/repos/%s/raw/%s?at=refs%%2Fheads%%2F%s", p.apiURL, project, p.repoName(pkg), name, branch)
	return p.httpClient.Get(ctx, url)
}

// GoGetURL gets url for go get
func (p *Provider) GoGetURL() string {
	return p.goGetURL
}

// todo cache result in map
func (p *Provider) getProject(ctx context.Context, pkg dep_radar.Pkg) (string, error) {
	p.muMap.RLock()
	if project, ok := p.mapProject[pkg]; ok {
		p.muMap.RUnlock()
		return project, nil
	}
	p.muMap.RUnlock()
	// if p.project != "" {
	// 	return nil
	// }
	project, err := GetProject(ctx, p.httpClient, pkg, p.gitDomain)
	if err != nil {
		return "", err
	}
	p.muMap.Lock()
	p.mapProject[pkg] = project
	p.muMap.Unlock()
	// fmt.Printf("project = %+v\n", project)
	// p.project = project
	return project, nil
}

// Tags get tags from bitbucket
func (p *Provider) Tags(ctx context.Context, pkg dep_radar.Pkg) ([]dep_radar.Tag, error) {
	project, err := p.getProject(ctx, pkg)
	if err != nil {
		return nil, err
	}
	var (
		tagsResult []dep_radar.Tag
		start      int
		isLastPage bool
	)
	for !isLastPage {
		tags, err := p.tags(ctx, pkg, project, start)
		if err != nil {
			return tagsResult, fmt.Errorf("Error on getting tags: %s", err)
		}
		for _, tag := range tags.Values {
			tagVersion := strings.Replace(tag.ID, "refs/tags/", "", 1)
			tagsResult = append(tagsResult, dep_radar.Tag{Version: tagVersion, Hash: dep_radar.Hash(tag.LatestCommit)})
		}
		isLastPage = tags.IsLastPage
		start = tags.NextPageStart
	}
	return tagsResult, nil
}

func (p *Provider) tags(ctx context.Context, pkg dep_radar.Pkg, project string, start int) (TagsResponse, error) {
	var tags TagsResponse
	url := fmt.Sprintf("%s/rest/api/1.0/projects/%s/repos/%s/tags?start=%d", p.apiURL, project, p.repoName(pkg), start)
	reposResponse, err := p.httpClient.Get(ctx, url)
	if err != nil {
		return tags, err
	}
	err = json.Unmarshal(reposResponse, &tags)

	return tags, err
}

func (p *Provider) repoName(pkg dep_radar.Pkg) string {
	strpkg := string(pkg)
	pos := strings.Index(strpkg, "/")
	if pos == -1 {
		return strpkg
	}
	return strpkg[pos+1:]
}
