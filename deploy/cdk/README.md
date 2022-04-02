# Doc Box CDK Project
This is the AWS CDK project creating the infrastructure needed for Doc Box.

## Requirements
* Unix System environment
* [Node.js](https://nodejs.org/en/download/)
* [Yarn - Package Manager](https://yarnpkg.com/)

## Getting Started

### Use projen for initial project generation
[Projen](https://github.com/projen/projen) is used for the AWS CDK project generation.

### Changing CDK configurations
The change of CDK configurations is possible in the Projen files, e.g.: *.projenrc.js*.
After that, rebuild the project with:
```shell
$ yarn projen
```

### Installing dependencies & building the target
The used build management tool is yarn, all yarn commands are executed in the /deploy/cdk folder.

```shell
$ yarn install
$ yarn build
```

### Quality Assurance and Tests
Execute the tests separately:
```bash
$ yarn test
```

#### Audit of dependencies
Execute an dependency audit:
```shell
$ yarn audit:level-high
$ yarn audit:dev-level-critical
```

## Deploy the stack
Execute the cdk deployment:
```shell
$ yarn cdk deploy --profile <<aws-account-profile>>
```

## Deploying manually lambda function on development stack
Before deploying, you should be edited the Makefile on the root folder and set the AWS_PROFILE variable.

### Build and deploy the lambda functions
The build and deploy of the lambda functions can the done on the root folder with:
```shell
$ make [lambda-name] e.g. createDocument
```