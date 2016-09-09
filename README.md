# httpmock [![Build Status](https://travis-ci.org/thingful/httpmock.png?branch=master)](https://travis-ci.org/thingful/httpmock)

This library is a fork of https://github.com/jarcoal/httpmock, which now has a
somewhat different API from the original library. All of the clever stuff was
worked out in the original version, this fork just adds some extra fussiness in
how it registers and responds to requests.

## What is the purpose of this fork?

The reason for creating this fork was that while the original httpmock library
provided a very neat mechanism for inserting mocked responses into the default
Go net/http client, it was intentionally designed to be very tolerant about the
requests it matched which lead to us seeing some bugs in production.

To give a specific example let's say our code requires requesting data from
some external resource (http://api.example.com/resource/1), and for correct
operation it must include in that request a specific authorization header
containing a value passed in from somewhere else within the program. For this
scenario the previous library would let us create a responder for
http://api.example.com/resource/1, but wouldn't allow us to force a test
failure if that resource was requested without the authorization header being
present.

## Usage examples

For usage examples, please see the documentation at
https://godoc.org/github.com/thingful/httpmock.
