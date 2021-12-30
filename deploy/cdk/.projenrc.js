const { awscdk } = require('projen');

const cdkVersion = '2.3.0';

const project = new awscdk.AwsCdkTypeScriptApp({
  cdkVersion: cdkVersion,
  defaultReleaseBranch: 'main',
  authorName: 'Thomas Renken',
  repository: 'git@github.com:torenken/doc-box.git',
  name: 'doc-box',

  deps: [
    '@aws-cdk/aws-lambda-go-alpha',
  ],

  context: {
    '@aws-cdk/core:newStyleStackSynthesis': true,
  },
  appEntrypoint: 'doc-box-app.ts',

  //no-github-workflow
  github: false,

  license: 'MIT',
  copyrightOwner: 'Thomas Renken',
});

project.setScript('postinstall', 'touch node_modules/go.mod'); //This step is necessary so that go mod tidy ignores the cdk-go deps.

project.synth();