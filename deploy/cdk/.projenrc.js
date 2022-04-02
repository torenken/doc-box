const { awscdk } = require('projen');

const cdkVersion = '2.12.0';
const appName = 'torenken-doc-box';

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
    //docBox
    'docBucketName': `${appName}-storage`,

    '@aws-cdk/core:newStyleStackSynthesis': true,
  },
  appEntrypoint: 'doc-box-app.ts',

  //no-github-workflow
  github: false,

  license: 'MIT',
  copyrightOwner: 'Thomas Renken',
});

project.setScript('postinstall', 'touch node_modules/go.mod'); //This step is necessary so that go mod tidy ignores the cdk-go deps.
project.setScript('audit:level-high', 'yarn audit --level high --groups dependencies; [[ $? -lt 7 ]] || exit 1');
project.setScript('audit:dev-level-critical', 'yarn audit --level critical --groups devDependencies; [[ $? -lt 16 ]] || exit 1');

project.synth();