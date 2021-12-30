import { GoFunction, GoFunctionProps } from '@aws-cdk/aws-lambda-go-alpha';
import { RetentionDays } from 'aws-cdk-lib/aws-logs';
import { Construct } from 'constructs';

export interface GoBaseFuncProps extends GoFunctionProps {
  readonly functionName: string;
  readonly applicationContext: string;
}

export class GoBaseFunc extends GoFunction {
  constructor(scope: Construct, id: string, props: GoBaseFuncProps) {
    super(scope, id, {
      functionName: props.functionName,
      entry: props.entry,
      logRetention: RetentionDays.ONE_MONTH,
      bundling: {
        goBuildFlags: ['-ldflags "-s -w"'],
        cgoEnabled: false,
      },
      memorySize: props.memorySize,
      environment: {
        APPLICATION_CONTEXT: props.applicationContext,
        ...props.environment,
      },
    });
  }
}