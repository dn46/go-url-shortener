# URL Shortener Service

This is a simple URL shortener service built with Go, Gin and Redis. It allows you to shorten long URLs.

## Features

- Generate short URLs
- Redirect to the original URL using the short URL

## Installation

1. Install Redis. You can download it from the [official Redis website](https://redis.io/download) - make sure to donwload the correct version for your OS.

2. Clone the repository:

`
git clone https://github.com/dn46/go-url-shortener.git
`

3. Navigate to the project directory:

`
cd go-url-shortener
`

4. Install the dependencies:

`
go mod download
`

## Usage

1. Start the server:

`
make run
`

2. Open your browser and navigate to `http://localhost:9808`

3. Create a short URL by using the input box with any URL and clicking `submit`.

4. Navigate to the new generated short URL presented to you.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.