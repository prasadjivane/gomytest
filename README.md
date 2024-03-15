# gomytest [![Go Reference](https://pkg.go.dev/badge/github.com/prasadjivane/gomytest.svg)](https://pkg.go.dev/github.com/prasadjivane/gomytest)
<img width="622" alt="gomytest" src="https://github.com/prasadjivane/gomytest/assets/26869583/48b83a90-9c2f-4d58-ad90-a7a140319da9">

Go package for testing APIs from the command line. `gomytest` is a command-line tool written in Go that allows you to test various HTTP methods (GET, POST, PUT, DELETE) against RESTful APIs.

## Features

- Supports testing of GET, POST, PUT, and DELETE HTTP methods.
- Prints response data, including status code, response body, and execution time.
- Allows testing of APIs from the command line.

## Installation

To install `gomytest`, you can either build it from source or download the pre-built binary for your platform from the [Releases](https://github.com/prasadjivane/gomytest/releases) page.

### Build from Source

To build from the source, make sure you have Go installed. Then, run the following commands:

```bash
git clone https://github.com/prasadjivane/gomytest.git

go build -o gomytest cmd/gomytest/main.go    
```


## Usage
To use gomytest, run the executable with the desired HTTP method, URL, and optional request data:

```bash

./gomytest <method> <url> [data]
```

Replace <method> with the desired HTTP method (GET, POST, PUT, DELETE), <url> with the endpoint URL you want to test, and [data] with any request data (for POST and PUT methods).

## Examples:

```bash
# GET request
./gomytest GET https://jsonplaceholder.typicode.com/posts/1

# POST request
./gomytest POST https://jsonplaceholder.typicode.com/posts '{"title": "gomytest", "body": "now", "userId": 1}'

# PUT request
./gomytest PUT https://jsonplaceholder.typicode.com/posts/1 '{"title": "now", "body": "gomytest", "userId": 1}'

# DELETE request
./gomytest DELETE https://jsonplaceholder.typicode.com/posts/1
```

## Contributing

Contributions are welcome! If you find any bugs or want to add new features, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License.
