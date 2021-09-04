# envlogdriver

[Zerolog](https://github.com/rs/zerolog) based logging libary optimized for serverless application.

## What is this package?

This package provides simple structured logger optimized for  serverless application based on [zerolog](https://github.com/rs/zerolog).

Key features of envlogdriver are:

- zerolog based simple method chaining API
- log level is set from environment variable

## Usage

First of all, initialize a logger.

```go
logger := zerodriver.NewLogger() // log level set from LOG_LEVEL environment variable. (if not exists then set to `debug`)
```

Then, write logs by using zerolog based fluent API!
```go
logger.Info().Str("key", "value").Msg("Hello World!")
// output: {"severity":"INFO","key":"value","time":"2009-11-10T23:00:00Z","message":"hello world"}
```

Here's complete example:

```go
package main

import (
    "github.com/ma91n/envlogdriver"
)

func main() {
    logger := envlogdriver.NewLogger()
    logger.Info().Str("key", "value").Msg("hello world")
}

// output: {"level":"INFO","key":"value","time":"2009-11-10T23:00:00Z","message":"hello world"}
```
