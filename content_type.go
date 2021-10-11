package grequests

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"reflect"
)

// DetectContentType method is used to figure out `Request.Body` content type for request header
func DetectContentType(body interface{}) string {
	contentType := plainTextType
	kind := reflect.Indirect(reflect.ValueOf(body)).Kind()
	switch kind {
	case reflect.Struct, reflect.Slice, reflect.Map:
		contentType = jsonContentType
	case reflect.String:
		contentType = plainTextType
	default:
	}
	return contentType
}

func handleRequestBody(opt *Options, body interface{}) (*bytes.Buffer, error) {
	if body == nil {
		return new(bytes.Buffer), nil
	}

	contentType := opt.Header.Get(HeaderContentType)
	if IsStringEmpty(contentType) {
		contentType = DetectContentType(body)
		opt.Header.Set(HeaderContentType, contentType)
	}

	bodyBytes, err := []byte{}, (error)(nil)
	kind := reflect.Indirect(reflect.ValueOf(body)).Kind()
	if reader, ok := body.(io.Reader); ok {
		bodyBytes, err = ioutil.ReadAll(reader)
	} else if b, ok := body.([]byte); ok {
		bodyBytes = b
	} else if s, ok := body.(string); ok {
		bodyBytes = []byte(s)
	} else if IsJSONContentType(contentType) && (kind == reflect.Struct || kind == reflect.Map || kind == reflect.Slice) {
		bodyBytes, err = json.Marshal(body)
	}
	return bytes.NewBuffer(bodyBytes), err
}
