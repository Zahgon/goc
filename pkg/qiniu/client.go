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

package qiniu

import (
	"context"
	"regexp"
	"time"

	"github.com/qiniu/api.v7/v7/storage"
)

// Config store the credentials to connect with qiniu cloud
type Config struct {
	Bucket    string `json:"bucket"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`

	// domain used to download files from qiniu cloud
	Domain string `json:"domain"`
}

// Client is the interface contains the operation with qiniu cloud
type Client interface {
	QiniuObjectHandle(key string) ObjectHandle
	ReadObject(key string) ([]byte, error)
	ListAll(ctx context.Context, prefix string, delimiter string) ([]string, error)
	GetAccessURL(key string, timeout time.Duration) string
	GetArtifactDetails(key string) (*LogHistoryTemplate, error)
	ListSubDirs(prefix string) ([]string, error)
}

// QnClient for the operation with qiniu cloud
type QnClient struct {
	cfg           *Config
	BucketManager *storage.BucketManager
}

// NewClient creates a new QnClient to work with qiniu cloud
func NewClient(cfg *Config) *QnClient { _ = "STUB: not implemented"; return nil }

// QiniuObjectHandle construct a object hanle to access file in qiniu
func (q *QnClient) QiniuObjectHandle(key string) ObjectHandle {
	_ = "STUB: not implemented"
	return *new(ObjectHandle)
}

// ReadObject to read all the content of key
func (q *QnClient) ReadObject(key string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListAll to list all the files with contains the expected prefix
func (q *QnClient) ListAll(ctx context.Context, prefix string, delimiter string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// listEntries to list all the entries with contains the expected prefix
func (q *QnClient) listEntries(prefix string, delimiter string) ([]storage.ListItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAccessURL return a url which can access artifact directly in qiniu
func (q *QnClient) GetAccessURL(key string, timeout time.Duration) string {
	_ = "STUB: not implemented"
	return ""
}

// LogHistoryTemplate is the template of the log history
type LogHistoryTemplate struct {
	BucketName string
	KeyPath    string
	Items      []logHistoryItem
}

// logHistoryItem represents a log history item
type logHistoryItem struct {
	Name string
	Size string
	Time string
	Url  string
}

// GetArtifactDetails lists all artifacts available for the given job source
func (q *QnClient) GetArtifactDetails(key string) (*LogHistoryTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func splitKey(item, key string) string { _ = "STUB: not implemented"; return "" }

func size(fsize int64) string { _ = "STUB: not implemented"; return "" }

func timeConv(ptime int64) string { _ = "STUB: not implemented"; return "" }

// ListSubDirs list all the sub directions of the prefix string in qiniu client
func (q *QnClient) ListSubDirs(prefix string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// use rsf list v2 interface to get the sub folder based on the delimiter

// entry.Dir should be like "logs/kodo-periodics-integration-test/1181915661132107776/"
// the sub folder is 1181915661132107776, also known as prowjob buildid.

var nonPRLogsBuildIdSubffixRe = regexp.MustCompile("([0-9]+)/$")

// extract the build number from dir path
// expect the dir as the following formats:
// 1. logs/kodo-periodics-integration-test/1181915661132107776/
func getBuildId(dir string) string { _ = "STUB: not implemented"; return "" }
