package httpmock

import (
	"fmt"
	"strings"
)

// ErrNoResponderFound is a custom error type used when no responders were
// found.
type ErrNoResponderFound struct {
	errs []error
}

// Error ensures our ErrNoResponderFound type implements the error interface
func (e *ErrNoResponderFound) Error() string {
	// TODO: is there a better way of giving a rich error message than this?

	if len(e.errs) == 0 {
		return "No registered stubs"
	}

	msg := `

Errors
----------------------------
%s
`
	errMsgs := []string{}
	for _, e := range e.errs {
		errMsgs = append(errMsgs, e.Error())
	}

	return fmt.Sprintf(msg, strings.Join(errMsgs, "\n"))
}

// NewErrNoResponderFound returns a new ErrNoResponderFound error
func NewErrNoResponderFound(errs []error) *ErrNoResponderFound {
	return &ErrNoResponderFound{
		errs: errs,
	}
}

// ErrStubsNotCalled is a type implementing the error interface we return when
// not all registered stubs were called
type ErrStubsNotCalled struct {
	uncalledStubs []*StubRequest
}

// Error ensures our ErrStubsNotCalled type implements the error interface
func (e *ErrStubsNotCalled) Error() string {
	// TODO: is there a better way of giving a rich error message than this?

	msg := `

Uncalled stubs
----------------------------
%s
`
	uncalled := []string{}
	for _, s := range e.uncalledStubs {
		uncalled = append(uncalled, s.String())
	}

	return fmt.Sprintf(msg, strings.Join(uncalled, "\n"))
}

// NewErrStubsNotCalled returns a new StubsNotCalled error
func NewErrStubsNotCalled(uncalledStubs []*StubRequest) *ErrStubsNotCalled {
	return &ErrStubsNotCalled{
		uncalledStubs: uncalledStubs,
	}
}
