/*
 Copyright 2020 Qiniu Cloud (qiniu.com)

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package github

import (
	"context"

	"github.com/google/go-github/github"

	"github.com/qiniu/goc/pkg/cover"
)

// CommentsPrefix is the prefix when commenting on Github Pull Requests
// It is also the flag when checking whether the target comment exists or not to avoid duplicate
const CommentsPrefix = "The following is the coverage report on the affected files."

// PrComment is the interface of the entry which is able to comment on Github Pull Requests
type PrComment interface {
	CreateGithubComment(commentPrefix string, diffCovList cover.DeltaCovList) (err error)
	PostComment(content, commentPrefix string) error
	EraseHistoryComment(commentPrefix string) error
	GetPrChangedFiles() (files []string, err error)
	GetCommentFlag() string
}

// GitPrComment is the entry which is able to comment on Github Pull Requests
type GitPrComment struct {
	RobotUserName string
	RepoOwner     string
	RepoName      string
	CommentFlag   string
	PrNumber      int
	Ctx           context.Context
	opt           *github.ListOptions
	GithubClient  *github.Client
}

// NewPrClient creates an Client which be able to comment on Github Pull Request
func NewPrClient(githubTokenPath, repoOwner, repoName, prNumStr, botUserName, commentFlag string) *GitPrComment {
	_ = "STUB: not implemented"
	return nil

	// performs automatic retries when connection error occurs or a 500-range response code received (except 501)
}

// CreateGithubComment post github comment of diff coverage
func (c *GitPrComment) CreateGithubComment(commentPrefix string, diffCovList cover.DeltaCovList) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// PostComment post comment on github. It erased the old one if existed to avoid duplicate
func (c *GitPrComment) PostComment(content, commentPrefix string) error {
	_ = "STUB: not implemented"
	//step1: erase history similar comment to avoid too many comment for same job
	return nil
}

//step2: post comment with new result

// EraseHistoryComment erase history similar comment before post again
func (c *GitPrComment) EraseHistoryComment(commentPrefix string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetPrChangedFiles get github pull request changes file list
func (c *GitPrComment) GetPrChangedFiles() (files []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetCommentFlag get CommentFlag from the GitPrComment
func (c *GitPrComment) GetCommentFlag() string { _ = "STUB: not implemented"; return "" }

// GenCommentContent generate github comment content based on diff coverage and commentFlag
func GenCommentContent(commentPrefix string, delta cover.DeltaCovList) string {
	_ = "STUB: not implemented"
	return ""
}
