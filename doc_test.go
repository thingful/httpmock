package httpmock

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func ExampleRegisterStubRequest() {
	Activate()
	defer DeactivateAndReset()

	RegisterStubRequest(
		NewStubRequest(
			"GET",
			"http://example.com/",
			NewStringResponder(200, "ok"),
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

	if err = AllStubsCalled(); err != nil {
		// handle error properly in real code
		panic(err)
	}

	// Output: ok
}

func ExampleRegisterStubRequest_WithHeader() {
	Activate()
	defer DeactivateAndReset()

	RegisterStubRequest(
		NewStubRequest(
			"GET",
			"http://example.com/",
			NewStringResponder(200, "ok"),
		).WithHeader(
			&http.Header{
				"Authorization": []string{"Bearer api-key"},
			},
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

	if err = AllStubsCalled(); err != nil {
		// handle error properly in real code
		panic(err)
	}

	// Output:
	// Error without header
	// ok
}
