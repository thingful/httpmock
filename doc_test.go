package httpmock_test

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/thingful/httpmock"
)

func ExampleRegisterStubRequests() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterStubRequests(
		httpmock.NewStubRequest(
			"GET",
			"http://example.com/",
			httpmock.NewStringResponder(200, "ok"),
		),
		httpmock.NewStubRequest(
			"GET",
			"http://another.com/",
			httpmock.NewStringResponder(200, "also ok"),
		),
	)

	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error properly in real code
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error properly in real code
		panic(err)
	}

	fmt.Println(string(body))

	resp, err = http.Get("http://another.com/")
	if err != nil {
		// handle error properly in real code
		panic(err)
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error properly in real code
		panic(err)
	}

	fmt.Println(string(body))

	if err = httpmock.AllStubsCalled(); err != nil {
		// handle error properly in real code
		panic(err)
	}

	// Output:
	// ok
	// also ok
}

func ExampleRegisterStubRequest_WithHeader() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterStubRequests(
		httpmock.NewStubRequest(
			"GET",
			"http://example.com/",
			httpmock.NewStringResponder(200, "ok"),
			httpmock.WithHeader(
				&http.Header{
					"Authorization": []string{"Bearer api-key"},
				},
			),
		),
	)

	_, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Println("Error without header")
	}

	req, err := http.NewRequest("GET", "http://example.com/", nil)
	if err != nil {
		// handle error properly
		panic(err)
	}

	req.Header.Add("Authorization", "Bearer api-key")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		// handle error properly in real code
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error properly in real code
		panic(err)
	}

	fmt.Println(string(body))

	if err = httpmock.AllStubsCalled(); err != nil {
		// handle error properly in real code
		panic(err)
	}

	// Output:
	// Error without header
	// ok
}
