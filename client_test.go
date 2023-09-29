package bunnynet

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCheckRespWithEmptyUnsuccessfulResp(t *testing.T) {
	req, err := http.NewRequest("get", "http://testbunny.com", nil)
	require.NoError(t, err)

	resp := http.Response{
		StatusCode: 400,
		Body:       io.NopCloser(strings.NewReader("")),
	}

	clt := NewClient("")

	err = clt.checkResp(req, &resp)
	require.Error(t, err)
	require.IsType(t, &HTTPError{}, err)

	httpErr := err.(*HTTPError)
	assert.Empty(t, httpErr.Errors)
}

func TestCheckRespWithJSONBody(t *testing.T) {
	apiErr := APIError{
		ErrorKey: "err",
		Field:    "id",
		Message:  "something br0ke",
	}

	buf, err := json.Marshal(&apiErr)
	require.NoError(t, err)

	const reqURL = "http://test.com"
	req, err := http.NewRequest("get", reqURL, nil)
	require.NoError(t, err)

	hdr := http.Header{}
	hdr.Add("content-type", "application/json; charset=utf-8")

	resp := http.Response{
		Header:     hdr,
		StatusCode: 400,
		Body:       io.NopCloser(bytes.NewReader(buf)),
	}

	clt := NewClient("")

	err = clt.checkResp(req, &resp)
	require.Error(t, err)
	require.IsType(t, &APIError{}, err, "error: "+err.Error())

	retAPIErr := err.(*APIError)
	assert.Equal(t, apiErr.ErrorKey, retAPIErr.ErrorKey, "unexpected errorKey value")
	assert.Equal(t, apiErr.Field, retAPIErr.Field, "unexpected field value")
	assert.Equal(t, apiErr.Message, retAPIErr.Message, "unexpected message value")

	assert.Equal(t, reqURL, retAPIErr.RequestURL, "unexpected RequestURL")
	assert.Equal(t, resp.StatusCode, retAPIErr.StatusCode, "unexpected status code")
	assert.Equal(t, buf, retAPIErr.RespBody)
}

func TestCheckRespWithJSONBodyAndMissingContentType(t *testing.T) {
	buf, err := json.Marshal(&APIError{Message: "something br0ke"})
	require.NoError(t, err)

	req, err := http.NewRequest("get", "", nil)
	require.NoError(t, err)

	resp := http.Response{
		StatusCode: 400,
		Body:       io.NopCloser(bytes.NewReader(buf)),
	}

	clt := NewClient("")

	err = clt.checkResp(req, &resp)
	require.Error(t, err)
	require.IsType(t, &HTTPError{}, err, "error: "+err.Error())

	retErr := err.(*HTTPError)
	assert.Equal(t, buf, retErr.RespBody)

	assert.EqualError(t, retErr.Errors[0], "processing response failed: content-type header is missing or empty")
}
