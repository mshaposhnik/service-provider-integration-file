package main

import (
	"context"
	"io"
	"regexp"
)

var gitHubRegexp = "p([a-z]+)ch"
var gitLabRegexp = "p([a-z]+)ch"

var github Github

var gitlab Gitlab

type Github struct {
	template *regexp.Regexp
}

type Gitlab struct {
	template *regexp.Regexp
}

func init() {
	github = Github{template: regexp.MustCompile(gitHubRegexp)}
	gitlab = Gitlab{template: regexp.MustCompile(gitLabRegexp)}
}

func getFileContents(ctx context.Context, repoUrl, filepath, ref string, callback func(url string)) (io.ReadCloser, error) {
	// TODO: how to iterate over all implementations ?
	if github.accept(repoUrl) {
		return github.fetch(ctx, repoUrl, filepath, ref, callback)
	}
}

func (gh Github) accept(repo string) bool {
	return gh.template.MatchString(repo)
}

func (gh Github) fetch(ctx context.Context, repoUrl, filepath, ref string, callback func(url string)) (io.ReadCloser, error) {
	//TODO implement me
}

func (gl Gitlab) accept(repo string) bool {
	return gl.template.MatchString(repo)
}

func (gl Gitlab) fetch(ctx context.Context, repoUrl, filepath, ref string, callback func(url string)) (io.ReadCloser, error) {
	//TODO implement me
}
