//go:build !assert

package assert

import (
	"net/http"
	"net/url"
)

// httpCode is a helper that returns HTTP code of the response. It returns -1 and
// an error if building a new request fails.
func httpCode(handler http.HandlerFunc, method, url string, values url.Values) {}

// HTTPSuccess asserts that a specified handler returns a success status code.
//
//	assert.HTTPSuccess(myHandler, "POST", "http://www.google.com", nil)
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPSuccess(handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...interface{}) {
}

// HTTPRedirect asserts that a specified handler returns a redirect status code.
//
//	assert.HTTPRedirect(myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPRedirect(handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...interface{}) {
}

// HTTPError asserts that a specified handler returns an error status code.
//
//	assert.HTTPError(myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPError(handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...interface{}) {
}

// HTTPStatusCode asserts that a specified handler returns a specified status code.
//
//	assert.HTTPStatusCode(myHandler, "GET", "/notImplemented", nil, 501)
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPStatusCode(handler http.HandlerFunc, method, url string, values url.Values, statuscode int, msgAndArgs ...interface{}) {
}

// HTTPBody is a helper that returns HTTP body of the response. It returns
// empty string if building a new request fails.
func HTTPBody(handler http.HandlerFunc, method, url string, values url.Values) {}

// HTTPBodyContains asserts that a specified handler returns a
// body that contains a string.
//
//	assert.HTTPBodyContains(myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPBodyContains(handler http.HandlerFunc, method, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) {
}

// HTTPBodyNotContains asserts that a specified handler returns a
// body that does not contain a string.
//
//	assert.HTTPBodyNotContains(myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPBodyNotContains(handler http.HandlerFunc, method, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) {
}
