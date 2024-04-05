//go:build assert

package assert

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
)

// httpCode is a helper that returns HTTP code of the response. It returns -1 and
// an error if building a new request fails.
func httpCode(handler http.HandlerFunc, method, url string, values url.Values) (int, error) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, http.NoBody)
	if err != nil {
		return -1, err
	}
	req.URL.RawQuery = values.Encode()
	handler(w, req)
	return w.Code, nil
}

// HTTPSuccess asserts that a specified handler returns a success status code.
//
//	assert.HTTPSuccess(myHandler, "POST", "http://www.google.com", nil)
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPSuccess(handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...interface{}) bool {
	code, err := httpCode(handler, method, url, values)
	if err != nil {
		Fail(fmt.Sprintf("Failed to build test request, got error: %s", err), msgAndArgs...)
	}

	isSuccessCode := code >= http.StatusOK && code <= http.StatusPartialContent
	if !isSuccessCode {
		Fail(fmt.Sprintf("Expected HTTP success status code for %q but received %d", url+"?"+values.Encode(), code), msgAndArgs...)
	}

	return isSuccessCode
}

// HTTPRedirect asserts that a specified handler returns a redirect status code.
//
//	assert.HTTPRedirect(myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPRedirect(handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...interface{}) bool {
	code, err := httpCode(handler, method, url, values)
	if err != nil {
		Fail(fmt.Sprintf("Failed to build test request, got error: %s", err), msgAndArgs...)
	}

	isRedirectCode := code >= http.StatusMultipleChoices && code <= http.StatusTemporaryRedirect
	if !isRedirectCode {
		Fail(fmt.Sprintf("Expected HTTP redirect status code for %q but received %d", url+"?"+values.Encode(), code), msgAndArgs...)
	}

	return isRedirectCode
}

// HTTPError asserts that a specified handler returns an error status code.
//
//	assert.HTTPError(myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPError(handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...interface{}) bool {
	code, err := httpCode(handler, method, url, values)
	if err != nil {
		Fail(fmt.Sprintf("Failed to build test request, got error: %s", err), msgAndArgs...)
	}

	isErrorCode := code >= http.StatusBadRequest
	if !isErrorCode {
		Fail(fmt.Sprintf("Expected HTTP error status code for %q but received %d", url+"?"+values.Encode(), code), msgAndArgs...)
	}

	return isErrorCode
}

// HTTPStatusCode asserts that a specified handler returns a specified status code.
//
//	assert.HTTPStatusCode(myHandler, "GET", "/notImplemented", nil, 501)
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPStatusCode(handler http.HandlerFunc, method, url string, values url.Values, statuscode int, msgAndArgs ...interface{}) bool {
	code, err := httpCode(handler, method, url, values)
	if err != nil {
		Fail(fmt.Sprintf("Failed to build test request, got error: %s", err), msgAndArgs...)
	}

	successful := code == statuscode
	if !successful {
		Fail(fmt.Sprintf("Expected HTTP status code %d for %q but received %d", statuscode, url+"?"+values.Encode(), code), msgAndArgs...)
	}

	return successful
}

// HTTPBody is a helper that returns HTTP body of the response. It returns
// empty string if building a new request fails.
func HTTPBody(handler http.HandlerFunc, method, url string, values url.Values) string {
	w := httptest.NewRecorder()
	if len(values) > 0 {
		url += "?" + values.Encode()
	}
	req, err := http.NewRequest(method, url, http.NoBody)
	if err != nil {
		return ""
	}
	handler(w, req)
	return w.Body.String()
}

// HTTPBodyContains asserts that a specified handler returns a
// body that contains a string.
//
//	assert.HTTPBodyContains(myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPBodyContains(handler http.HandlerFunc, method, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) bool {
	body := HTTPBody(handler, method, url, values)

	contains := strings.Contains(body, fmt.Sprint(str))
	if !contains {
		Fail(fmt.Sprintf("Expected response body for \"%s\" to contain \"%s\" but found \"%s\"", url+"?"+values.Encode(), str, body), msgAndArgs...)
	}

	return contains
}

// HTTPBodyNotContains asserts that a specified handler returns a
// body that does not contain a string.
//
//	assert.HTTPBodyNotContains(myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPBodyNotContains(handler http.HandlerFunc, method, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) bool {
	body := HTTPBody(handler, method, url, values)

	contains := strings.Contains(body, fmt.Sprint(str))
	if contains {
		Fail(fmt.Sprintf("Expected response body for \"%s\" to NOT contain \"%s\" but found \"%s\"", url+"?"+values.Encode(), str, body), msgAndArgs...)
	}

	return !contains
}
