# api-test
Go package for testing APIs from the command line.

# gomytest

`gomytest` is a simple command-line tool written in Go that allows you to test various HTTP methods (GET, POST, PUT, DELETE) against RESTful APIs.

## Features

- Supports testing of GET, POST, PUT, and DELETE HTTP methods.
- Prints response data, including status code, response body, and execution time.
- Allows testing of APIs from the command line.

## Installation

To install `apitester`, you can either build it from source or download the pre-built binary for your platform from the [Releases](https://github.com/prasadjivane/apitester/releases) page.

### Build from Source

To build from the source, make sure you have Go installed. Then, run the following commands:

```bash
git clone https://github.com/prasadjivane/apitester.git
cd apitester/cmd/apitester
go build -o apitester main.go
