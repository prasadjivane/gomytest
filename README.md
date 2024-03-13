# gomytest

Go package for testing APIs from the command line.

`gomytest` is a command-line tool written in Go that allows you to test various HTTP methods (GET, POST, PUT, DELETE) against RESTful APIs.

## Features

- Supports testing of GET, POST, PUT, and DELETE HTTP methods.
- Prints response data, including status code, response body, and execution time.
- Allows testing of APIs from the command line.

## Installation

To install `apitester`, you can either build it from source or download the pre-built binary for your platform from the [Releases](https://github.com/prasadjivane/apitester/releases) page.

### Build from Source

To build from the source, make sure you have Go installed. Then, run the following commands:

```bash
git clone https://github.com/prasadjivane/gomytest.git

go build -o gomytest cmd/gomytest/main.go    
```


## Usage
To use gomytest, run the executable with the desired HTTP method, URL, and optional request data:

```bash

./apitester <method> <url> [data]
```

Replace <method> with the desired HTTP method (GET, POST, PUT, DELETE), <url> with the endpoint URL you want to test, and [data] with any request data (for POST and PUT methods).

## Examples:

```bash
# GET request
./apitester GET https://jsonplaceholder.typicode.com/posts/1

# POST request
./apitester POST https://jsonplaceholder.typicode.com/posts '{"title": "foo", "body": "bar", "userId": 1}'

# PUT request
./apitester PUT https://jsonplaceholder.typicode.com/posts/1 '{"title": "foo", "body": "bar", "userId": 1}'

# DELETE request
./apitester DELETE https://jsonplaceholder.typicode.com/posts/1
```

## Contributing

Contributions are welcome! If you find any bugs or want to add new features, feel free to open an issue or submit a pull request.

## License

This project is licensed under the ISC License.
