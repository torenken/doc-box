import { App } from 'aws-cdk-lib';

export class TestApp extends App {
  constructor() {
    super({
      context: {
        'aws:cdk:bundling-stacks': ['NoStack'], //disable bundling lambda (asset) to reduce the test-runtime.
        '@aws-cdk/core:newStyleStackSynthesis': 'true',
      },
    });
  }
}