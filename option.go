package grequests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Options opt ...
type Options struct {
	Header         http.Header
	HeaderFunc     func(dst http.Header, src http.Header) http.Header // default DefaultHeaderAppend
	ResponseHeader func(http.Header)
	PathParams     map[string]string
	QueryParams    map[string]string
	Decoder        func(data []byte, out interface{}) error // default json.Unmarshal
	// Coder          func(data interface{}) ([]byte, error)
}

// ParamURL url
func (o *Options) ParamURL(urlPath string) string {
	// path
	for key, value := range o.PathParams {
		urlPath = strings.ReplaceAll(urlPath, ":"+url.PathEscape(key), url.PathEscape(value))
	}

	// query
	urlPath += "?"
	for key, value := range o.QueryParams {
		urlPath = fmt.Sprintf("%s%s=%s&", urlPath, url.QueryEscape(key), url.QueryEscape(value))
	}

	return strings.TrimSuffix(strings.TrimSuffix(urlPath, "&"), "?")
}

// Option 筛选器
type Option func(*Options)

// MergeOption merge
func MergeOption(opts ...Option) *Options {
	opt := Options{
		Header:         make(http.Header),
		HeaderFunc:     DefaultHeaderAppend,
		PathParams:     make(map[string]string),
		QueryParams:    make(map[string]string),
		Decoder:        json.Unmarshal,
		ResponseHeader: func(h http.Header) {},
	}
	for _, v := range opts {
		v(&opt)
	}
	return &opt
}

// AddHeader add header
func AddHeader(key string, values ...string) Option {
	return func(o *Options) {
		for _, value := range values {
			o.Header.Add(key, value)
		}
	}
}

// DelHeader del header
func DelHeader(keys ...string) Option {
	return func(o *Options) {
		for _, key := range keys {
			o.Header.Del(key)
		}
	}
}

// SetHeader set header
func SetHeader(header http.Header) Option {
	return func(o *Options) {
		o.Header = header
	}
}

// AddPathParams add path params ep: /api/v/:bar/:foo, key=name
func AddPathParams(key, value string) Option {
	return func(o *Options) {
		o.PathParams[key] = value
	}
}

// SetPathParams set path params ep: /api/v/:bar/:foo, key=name
func SetPathParams(params map[string]string) Option {
	return func(o *Options) {
		o.PathParams = params
	}
}

// DelPathParams add path params ep: /api/v/:bar/:foo, key=name
func DelPathParams(key string) Option {
	return func(o *Options) {
		delete(o.PathParams, key)
	}
}

// AddQueryParams add query params ep: /api/v/books?bar=baz&foo=quux
func AddQueryParams(key, value string) Option {
	return func(o *Options) {
		o.QueryParams[key] = value
	}
}

// SetQueryParams set query params ep: /api/v/books?bar=baz&foo=quux
func SetQueryParams(params map[string]string) Option {
	return func(o *Options) {
		o.QueryParams = params
	}
}

// DelQueryParams del query params ep: /api/v/books?bar=baz&foo=quux
func DelQueryParams(key string) Option {
	return func(o *Options) {
		delete(o.QueryParams, key)
	}
}

// SetHeaderFunc hander func
func SetHeaderFunc(fn func(http.Header, http.Header) http.Header) Option {
	return func(o *Options) {
		o.HeaderFunc = fn
	}
}

// DefaultHeaderAppend 默认追加方式
func DefaultHeaderAppend(dst, src http.Header) http.Header {
	for k, v := range src {
		dst[k] = append(dst[k], v...)
	}
	return dst
}

// DefaultHeaderCover 覆盖的方式
func DefaultHeaderCover(dst, src http.Header) http.Header {
	return src
}

// SetDecoder 设置 decode
func SetDecoder(fn func(data []byte, out interface{}) error) Option {
	return func(o *Options) {
		o.Decoder = fn
	}
}
