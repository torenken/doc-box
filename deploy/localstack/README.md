# LocalStack
[LocalStack](https://github.com/localstack/localstack) provides an easy-to-use test/mocking framework for developing
Cloud applications. We use LocalStack for local integration tests.

## Requirements
* Unix System environment
* Configured [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-welcome.html) with administrator permissions
* [Docker installed](https://www.docker.com/community-edition/)
* [Docker Compose](https://docs.docker.com/compose/)

## Getting Started

### Starting LocalStack
Start localstack docker container:
```shell
$ docker-compose up -d
```

### Create cloud components & insert data
Set up the LocalStack configuration for integration tests:
```shell
$ ./localstack.sh -c
```

### Remove cloud components & data
Clean up the LocalStack configuration:
```shell
$ ./localstack.sh -d
```