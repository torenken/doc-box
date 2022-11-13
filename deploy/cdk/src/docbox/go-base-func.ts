import { GoFunction } from '@aws-cdk/aws-lambda-go-alpha';
import { Duration } from 'aws-cdk-lib';
import { Architecture } from 'aws-cdk-lib/aws-lambda';
import { RetentionDays } from 'aws-cdk-lib/aws-logs';
import { Construct } from 'constructs';

export interface GoBaseFuncProps {
  readonly context: string;
  readonly useCase: string;
  readonly entry: string;
  readonly environment?: Record<string, string>;
  readonly memorySize?: number;
  readonly timeout?: Duration;
}

export class GoBaseFunc extends GoFunction {
  constructor(scope: Construct, id: string, props: GoBaseFuncProps) {
    super(scope, id, {
      functionName: `doc-box-${props.useCase}`,
      entry: props.entry,
      logRetention: RetentionDays.ONE_MONTH,
      architecture: Architecture.ARM_64,
      bundling: {
        goBuildFlags: ['-ldflags "-s -w"'],
        cgoEnabled: false,
      },
      memorySize: props.memorySize,
      environment: {
        APPLICATION_CONTEXT: props.context,
        ...props.environment,
      },
    });
  }
}