const { awscdk } = require('projen');

const appName = 'torenken-doc-box';

const project = new awscdk.AwsCdkTypeScriptApp({
  cdkVersion: '2.60.0',
  defaultReleaseBranch: 'main',
  authorName: 'Thomas Renken',
  repository: 'git@github.com:torenken/doc-box.git',
  name: appName,

  deps: [
    '@aws-cdk/aws-lambda-go-alpha',
  ],

  context: {
    //docBox
    'docBucketName': `${appName}-storage`,

    '@aws-cdk/core:newStyleStackSynthesis': true,
  },
  appEntrypoint: 'doc-box-app.ts',

  jestOptions: {
    jestVersion: '28',
  },

  //no-github-workflow
  github: false,

  license: 'MIT',
  copyrightOwner: 'Thomas Renken',
});

project.setScript('postinstall', 'touch node_modules/go.mod'); //This step is necessary so that go mod tidy ignores the cdk-go deps.
project.setScript('audit:level-high', '/bin/bash -c \'yarn audit; [[ $? -ge 8 ]] && exit 1 || exit 0\'');

project.synth();