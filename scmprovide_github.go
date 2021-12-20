package main

import (
	"regexp"
	"strings"
)

var GithubURLRegexp = regexp.MustCompile(`(?Um)^(?:http)(?:s)?(?:\:\/\/)github.com/(?P<repoUser>[^/]+)/(?P<repoName>[^/]+)((/)|(?:/tree/(?P<branchName>[^/]+)(?:/(?P<subFolder>.*))?)|(/pull/(?P<pullRequestId>[^/]+)))?$`)
var GithubURLRegexpNames = GithubURLRegexp.SubexpNames()

// GitHubScmProvider implements Detector to detect GitHub URLs.
type GitHubScmProvider struct {
}

func (d *GitHubScmProvider) Detect(repoUrl, filepath, ref string) (bool, string, error) {
	if len(repoUrl) == 0 {
		return false, "", nil
	}

	GithubURLRegexp.FindStringSubmatch(repoUrl)
	if GithubURLRegexp.MatchString(repoUrl) {
		result := GithubURLRegexp.FindAllStringSubmatch(repoUrl, -1)
		m := map[string]string{}
		for i, n := range result[0] {
			m[GithubURLRegexpNames[i]] = n
		}
		str := []string{"https://raw.githubusercontent.com", m["repoUser"], m["repoName"], "HEAD", filepath}
		return true, strings.Join(str, "/"), nil
	}

	return false, "", nil
}
