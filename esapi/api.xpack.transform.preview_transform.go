// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated from specification version 8.0.0: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"net/http"
	"strings"
)

func newTransformPreviewTransformFunc(t Transport) TransformPreviewTransform {
	return func(body io.Reader, o ...func(*TransformPreviewTransformRequest)) (*Response, error) {
		var r = TransformPreviewTransformRequest{Body: body}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// TransformPreviewTransform -
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/preview-transform.html.
//
type TransformPreviewTransform func(body io.Reader, o ...func(*TransformPreviewTransformRequest)) (*Response, error)

// TransformPreviewTransformRequest configures the Transform Preview Transform API request.
//
type TransformPreviewTransformRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r TransformPreviewTransformRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(len("/_data_frame/transforms/_preview"))
	path.WriteString("/_data_frame/transforms/_preview")

	params = make(map[string]string)

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, err := newRequest(method, path.String(), r.Body)
	if err != nil {
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if r.Body != nil {
		req.Header[headerContentType] = headerContentTypeJSON
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f TransformPreviewTransform) WithContext(v context.Context) func(*TransformPreviewTransformRequest) {
	return func(r *TransformPreviewTransformRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f TransformPreviewTransform) WithPretty() func(*TransformPreviewTransformRequest) {
	return func(r *TransformPreviewTransformRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f TransformPreviewTransform) WithHuman() func(*TransformPreviewTransformRequest) {
	return func(r *TransformPreviewTransformRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f TransformPreviewTransform) WithErrorTrace() func(*TransformPreviewTransformRequest) {
	return func(r *TransformPreviewTransformRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f TransformPreviewTransform) WithFilterPath(v ...string) func(*TransformPreviewTransformRequest) {
	return func(r *TransformPreviewTransformRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f TransformPreviewTransform) WithHeader(h map[string]string) func(*TransformPreviewTransformRequest) {
	return func(r *TransformPreviewTransformRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}
