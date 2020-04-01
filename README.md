OpenAPI Specification object model
===

[![GoDoc](https://godoc.org/github.com/bentranter/go-openapi?status.svg)](https://godoc.org/github.com/bentranter/go-openapi)
[![Build Status](https://travis-ci.org/bentranter/go-openapi.svg?branch=master)](https://travis-ci.org/bentranter/go-openapi)
[![codecov](https://codecov.io/gh/bentranter/go-openapi/branch/master/graph/badge.svg)](https://codecov.io/gh/bentranter/go-openapi)

---

This is a fork of https://github.com/nasa9084/go-openapi. It was changed to add support for:
  * JSON marshalling.
  * Omitting empty fields.

**This package is still under development, so the API will be changed without any notification**

## Overview

This is an implementation of [OpenAPI Specification 3.0](https://github.com/OAI/OpenAPI-Specification) object model with some usable functions.

## Synopsis

``` go
package main

import (
    "fmt"

    "github.com/bentranter/go-openapi"
)

func main() {
    doc, _ := openapi.LoadFile("path/to/spec")
    fmt.Print(doc.Version)
}
```

## Status

* [x] Model definition
* [x] Load OpenAPI 3.0 spec file
* [ ] Resolve Reference object
  * [x] Resolve #/component reference
  * [ ] Resolve other file reference
* [ ] Validation
  * [x] Validate spec values
    * [ ] test for validation
      * [x] Document
      * [x] Info
      * [x] Contact
      * [x] License
      * [x] Server
      * [x] ServerVariable
      * [x] Paths
      * [x] PathItem
      * [x] Operation
      * [x] Parameter
      * [x] RequestBody
      * [x] Responses
      * [x] Response
      * [x] Callbacks
      * [x] Callback
      * [ ] Schema
      * [x] Example
      * [ ] MediaType
      * [ ] Header
      * [ ] Link
      * [ ] Encoding
      * [x] Discriminator
      * [x] XML
      * [x] Components
      * [x] SecurityScheme
      * [x] OAuthFlows
      * [x] OAuthFlow
      * [ ] SecurityRequirement
      * [x] Tag
      * [x] ExternalDocumentation
  * [ ] Validate HTTP Request
  * [ ] Validate HTTP Response
