import { App, Tags } from 'aws-cdk-lib';
import { DocBoxStack } from './doc-box-stack';

const app = new App();

Tags.of(app).add('domain', 'document-management');
Tags.of(app).add('owner', 'torenken');

const docBucketName = app.node.tryGetContext('docBucketName') as string;

new DocBoxStack(app, 'DocBoxStack', {
  documentBucketName: docBucketName,
});

app.synth();