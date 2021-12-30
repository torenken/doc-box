import { App, Tags } from 'aws-cdk-lib';
import { DocBoxStack } from './doc-box-stack';

const app = new App();

Tags.of(app).add('domain', 'document-management');
Tags.of(app).add('owner', 'torenken');

new DocBoxStack(app, 'DocBoxStack');

app.synth();