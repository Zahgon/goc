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
	"io"
	"net/http"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/client"
	"github.com/qiniu/api.v7/v7/storage"
)

// ObjectHandle is the interface contains the operations on an object in a qiniu cloud bucket
type ObjectHandle interface {
	NewReader(ctx context.Context) (io.ReadCloser, error)
	NewRangeReader(ctx context.Context, offset, length int64) (io.ReadCloser, error)
}

// QnObjectHandle provides operations on an object in a qiniu cloud bucket
type QnObjectHandle struct {
	key    string
	cfg    *Config
	bm     *storage.BucketManager
	mac    *qbox.Mac
	client *client.Client
}

// NewReader creates a reader to read the contents of the object.
// ErrObjectNotExist will be returned if the object is not found.
// The caller must call Close on the returned Reader when done reading.
func (o *QnObjectHandle) NewReader(ctx context.Context) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// NewRangeReader reads parts of an object, reading at most length bytes starting
// from the given offset. If length is negative, the object is read until the end.
func (o *QnObjectHandle) NewRangeReader(ctx context.Context, offset, length int64) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// The end character isn't affected by how many bytes we have seen.

//TODO enhance

func runWithRetry(maxTry int, f func() (bool, error)) error { _ = "STUB: not implemented"; return nil }

// fix -  needRetry, err := f(), err hides the outside error

func shouldRetry(res *http.Response) bool {
	_ = "STUB: not implemented"

	// 571 and 573 mean the request was limited by cloud storage because of concurrency count exceed
	// so it's better to retry after a while
	return false
}
