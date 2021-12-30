# <img src="doc/doc-box.png" width="220" height="80">

## Purpose
The Doc Box is a serverless application that implements the business capabilities of the business domain
[Document Management](https://www.tmforum.org/resources/standard/tmf667-document-management-api-rest-specification-r17-0-1/).

Currently, the following business capabilities are supported:
- Create Document-Metadata and add Attachments

## Project structure
The project contains of source code and files for a serverless application that can be deployed via AWS CDK.
See **deploy/cdk/README.md** for build of the AWS CDK project.

| Folder              | Description                                                                          |
|---------------------|:-------------------------------------------------------------------------------------|
 | doc-box             | Root folder for building, testing, etc.                                              |
| doc-box/**cnd**     | Contains the handlers for the implemented serverless functions.                      |
| doc-box/**deploy**  | Contains the AWS CDK stack, local test infrastructure (localstack) and needed tools. |
| doc-box/**doc**     | Contains the specification and documentation.                                        |
| doc-box/**pkg**     | Contains the internal source code for the implemented serverless functions.          |
| doc-box/**pkg/tmf** | Contains the generated tmf api model.                                                |

## Requirements
* Unix System environment
* Configured [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-welcome.html) with administrator permissions
* [Docker installed](https://www.docker.com/community-edition)
* [Golang](https://golang.org)
* [Node.js](https://nodejs.org/en/download/)
* [Yarn - Package Manager](https://yarnpkg.com/)

## Getting Started
### Installing dependencies & building the target
The used build management tool is Make, all Make commands are executed in the root folder.

#### Building the target
Complete Build of Golang project via Makefile
```shell
$ make
```