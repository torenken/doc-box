# OpenAPI for Doc Box

## Purpose
The OpenAPI section generates corresponding data models and API documentation from the 
OpenAPI 3.0 specification.

## Requirements
* Unix System environment
* [Golang](https://golang.org)
* [Java11](https://www.oracle.com/de/java/technologies/javase-jdk11-downloads.html)

## Getting Started
### Downloading openapi-generator & generate the data model and documentation

```shell
$ ./openapi-cli -d (first time downloading the openapi-generator-cli)
```

The target folders for generation can be configured in header of the script `openapi-cli` and in the config folder `./config`. 
```shell
$ ./openapi-cli -g
```
