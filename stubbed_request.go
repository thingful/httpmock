package httpmock

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	// ErrIncorrectMethod is the error returned when Matches detects an incorrect
	// HTTP method
	ErrIncorrectMethod = errors.New("Incorrect request method attempted")

	// ErrIncorrectURL is the error returned when Matches detects an incorrect
	// normalized URL
	ErrIncorrectURL = errors.New("Incorrect URL used")

	// ErrIncorrectHeaders is the error returned when Matches detects incorrect
	// headers sent
	ErrIncorrectHeaders = errors.New("Incorrect HTTP headers sent")

	// ErrIncorrectRequestBody = is the error returned when Matches detects an
	// incorrect request body was sent
	ErrIncorrectRequestBody = errors.New("Incorrect request body sent")
)

// Option is a functional configurator allowing non-standard configuration to be
// added to the StubRequest without the previous chained function call approach.
type Option func(*StubRequest)

// NewStubRequest is a constructor function that returns a StubRequest for the
// given method and url. We also supply a responder which actually generates
// the response should the stubbed request match the request.
func NewStubRequest(method, url string, responder Responder, options ...Option) *StubRequest {
	r := &StubRequest{
		Method:    method,
		URL:       url,
		Responder: responder,
	}

	for _, option := range options {
		option(r)
	}

	return r
}

// StubRequest is used to capture data about a new stubbed request. It wraps up
// the Method and URL along with optional http.Header struct, holds the
// Responder used to generate a response, and also has a flag indicating
// whether or not this stubbed request has actually been called.
type StubRequest struct {
	Method    string
	URL       string
	Header    *http.Header
	Body      io.Reader
	Responder Responder
	Called    bool
}

// WithHeader is a function used to add http headers onto a stubbed request.
//
// Deprecated: use the functional WithHeader configuration function instead
func (r *StubRequest) WithHeader(header *http.Header) *StubRequest {
	r.Header = header
	return r
}

// WithBody is a function used to add a body to a stubbed request
//
// Deprecated: use the functional WithBody configuration function instead
func (r *StubRequest) WithBody(body io.Reader) *StubRequest {
	r.Body = body
	return r
}

// WithHeader is a functional configuration option used to add http headers onto
// a stubbed request
func WithHeader(header *http.Header) Option {
	return func(r *StubRequest) {
		r.Header = header
	}
}

// WithBody is a functional configuration option used to add a body to a stubbed
// request
func WithBody(body io.Reader) Option {
	return func(r *StubRequest) {
		r.Body = body
	}
}

// Matches is a test function that returns true if an incoming request is
// matched by this fetcher. Should an incoming request URL cause an error when
// normalized, we return false.
func (r *StubRequest) Matches(req *http.Request) error {
	if !strings.EqualFold(req.Method, r.Method) {
		return ErrIncorrectMethod
	}

	normalizedURL, err := normalizeURL(r.URL)
	if err != nil {
		return err
	}

	normalizedReqURL, err := normalizeURL(req.URL.String())
	if err != nil {
		return err
	}

	if normalizedURL != normalizedReqURL {
		return ErrIncorrectURL
	}

	// only check headers if the stubbed request has set headers to some not nil
	// value
	if r.Header != nil {

		// for each header defined on the stub, iterate through all the values and
		// make sure they are present in the corresponding header on the request
		for header, stubValues := range map[string][]string(*r.Header) {
			// get the values for this header on the request
			reqValues := req.Header[http.CanonicalHeaderKey(header)]
			for _, v := range stubValues {
				if !contains(reqValues, v) {
					return ErrIncorrectHeaders
				}
			}
		}
	}

	// if our stub includes a body, then it should equal the actual request body
	// to match
	if r.Body != nil {
		stubBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}

		requestBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return err
		}

		if bytes.Compare(stubBody, requestBody) != 0 {
			return ErrIncorrectRequestBody
		}
	}

	return nil
}

func (r *StubRequest) String() string {
	str := fmt.Sprintf("%s %s", r.Method, r.URL)

	if r.Header != nil {
		str = str + fmt.Sprintf(" with headers %v", r.Header)
	}

	return str
}
