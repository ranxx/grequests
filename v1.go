package grequests

import (
	"context"
	"io/ioutil"
	"net/http"
)

func respHeader(opt *Options, h http.Header, e *error) {
	if e == nil || *e == nil || opt == nil || opt.ResponseHeader == nil {
		return
	}
	opt.ResponseHeader(h)
}

func request(ctx context.Context, method, url string, body, out interface{}, opts ...Option) (e error) {
	// 默认 header
	opt := MergeOption(opts...)

	url = opt.ParamURL(url)

	bbb, err := handleRequestBody(opt, body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, url, bbb)
	if err != nil {
		return err
	}

	req.Header = opt.HeaderFunc(req.Header, opt.Header)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if out == nil {
		return nil
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	defer respHeader(opt, resp.Header, &e)

	return opt.Decoder(bodyBytes, out)
}

// Post post ...
func Post(ctx context.Context, url string, body, out interface{}, opts ...Option) error {
	return request(ctx, http.MethodPost, url, body, out, opts...)
}

// Put put ...
func Put(ctx context.Context, url string, body, out interface{}, opts ...Option) error {
	return request(ctx, http.MethodPut, url, body, out, opts...)
}

// Get get ...
func Get(ctx context.Context, url string, body, out interface{}, opts ...Option) error {
	return request(ctx, http.MethodGet, url, body, out, opts...)
}

// Delete delete ...
func Delete(ctx context.Context, url string, body, out interface{}, opts ...Option) error {
	return request(ctx, http.MethodDelete, url, body, out, opts...)
}
