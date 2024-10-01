# email

[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/dmitrymomot/email)](https://github.com/dmitrymomot/email/tags)
[![Go Reference](https://pkg.go.dev/badge/github.com/dmitrymomot/email.svg)](https://pkg.go.dev/github.com/dmitrymomot/email)
[![License](https://img.shields.io/github/license/dmitrymomot/email)](https://github.com/dmitrymomot/email/blob/main/LICENSE)


[![Tests](https://github.com/dmitrymomot/email/actions/workflows/tests.yml/badge.svg)](https://github.com/dmitrymomot/email/actions/workflows/tests.yml)
[![CodeQL Analysis](https://github.com/dmitrymomot/email/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/dmitrymomot/email/actions/workflows/codeql-analysis.yml)
[![GolangCI Lint](https://github.com/dmitrymomot/email/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/dmitrymomot/email/actions/workflows/golangci-lint.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/dmitrymomot/email)](https://goreportcard.com/report/github.com/dmitrymomot/email)

Email Validation and Sanitization Module

## Summary

This Go module provides comprehensive utilities for validating and sanitizing email addresses. It includes functions to validate email formats, check domain blacklists, sanitize email addresses, and more. The module is designed to be flexible and extendable, allowing you to add custom validators and sanitizers as needed.

## Installation

To install this module, use the following `go get` command:

```sh
go get github.com/dmitrymomot/email
```

## Usage

### Importing the Module

First, import the module in your Go code:

```go
import (
    "github.com/dmitrymomot/email"
)
```

### Validating Email Addresses

You can validate email addresses using the `ValidateEmail` function. This function accepts the email address and a list of custom validators.

```go
package main

import (
    "fmt"
    "github.com/dmitrymomot/email"
)

func main() {
    emailAddr := "test@example.com"
    err := email.ValidateEmail(emailAddr, email.ValidateUsernameFormat, email.ValidateIcanSuffix, email.IsAddressBlacklisted)
    if err != nil {
        fmt.Println("Invalid email:", err)
    } else {
        fmt.Println("Valid email")
    }
}
```

### Sanitizing Email Addresses

You can sanitize email addresses using the `SanitizeEmail` function. This function accepts the email address and a list of custom sanitizers.

```go
package main

import (
    "fmt"
    "github.com/dmitrymomot/email"
)

func main() {
    emailAddr := " Test+spam@example.com "
    sanitizedEmail := email.SanitizeEmail(emailAddr, email.TrimSpace, email.ToLower, email.RemoveAfterPlus)
    fmt.Println("Sanitized email:", sanitizedEmail)
}
```

### Managing Blacklisted Domains

You can manage blacklisted domains using the provided functions:

```go
package main

import (
    "fmt"
    "github.com/dmitrymomot/email"
)

func main() {
    // Add a domain to the blacklist
    email.AddBlacklist("example.com")

    // Check if a domain is blacklisted
    isBlacklisted := email.IsDomainBlacklisted("example.com")
    fmt.Println("Is example.com blacklisted?", isBlacklisted)

    // Remove a domain from the blacklist
    email.RemoveBlacklist("example.com")
}
```

## Contributing

Contributions are welcome! If you find a bug or have a feature request, please open an issue on the GitHub repository. If you want to contribute code, please fork the repository and submit a pull request.

### Steps to Contribute

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Make your changes.
4. Commit your changes with a descriptive commit message.
5. Push your changes to your fork.
6. Open a pull request to the main repository.

## License

This module is licensed under the Apache 2.0 License. See the [LICENSE](LICENSE) file for more details.
