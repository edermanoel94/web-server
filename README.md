# web-server

A very simple implementation of a multi-thread static webserver using HTTP over TCP. The main goal here is compare with rust.

## Getting Started

Install golang

### Running

1. Clone the repo
2. Run the server

```sh
go build
./web-server
```

### Usage

1. Open your browser
2. Use one tab to access http://127.0.0.1:8080/sleep, which will open a slow page (2 seconds sleeping thread).
3. Use another tab to access http://127.0.0.1:8080/, which will open a quick page (no sleep).
