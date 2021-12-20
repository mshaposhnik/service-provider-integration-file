package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"regexp"
)

var gitHubRegexp = regexp.MustCompile(`(?m)^https://github.com/.git`) //example
var gitLabRegexp = regexp.MustCompile(``)                             //example

var github Github

var gitlab Gitlab

type Github struct {
	template *regexp.Regexp
}

type Gitlab struct {
	template *regexp.Regexp
}

func init() {
	github = Github{template: gitHubRegexp}
	gitlab = Gitlab{template: gitLabRegexp}
}

func getFileContents(ctx context.Context, repoUrl, filepath, ref string, callback func(url string)) (io.ReadCloser, error) {
	// TODO: how to iterate over all implementations ?
	if github.accept(repoUrl) {
		return github.fetch(ctx, repoUrl, filepath, ref, callback)
	}
	return nil, errors.New(fmt.Sprintf("Unknown scm provider fpr %s", repoUrl))
}

func (gh Github) accept(repo string) bool {
	return gh.template.MatchString(repo)
}

func (gh Github) fetch(ctx context.Context, repoUrl, filepath, ref string, callback func(url string)) (io.ReadCloser, error) {
	//TODO implement me
	return nil, errors.New("not implemented")
}

func (gl Gitlab) accept(repo string) bool {
	return gl.template.MatchString(repo)
}

func (gl Gitlab) fetch(ctx context.Context, repoUrl, filepath, ref string, callback func(url string)) (io.ReadCloser, error) {
	//TODO implement me
	return nil, errors.New("not implemented")
}
