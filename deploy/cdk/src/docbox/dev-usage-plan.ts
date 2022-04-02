import { IApiKey, IRestApi, UsagePlan } from 'aws-cdk-lib/aws-apigateway';
import { Construct } from 'constructs';

export interface DevUsagePlanProps {
  readonly documentApi: IRestApi;
  readonly devApiKeys: IApiKey[];
}

export class DevUsagePlan extends UsagePlan {
  constructor(scope: Construct, id: string, props: DevUsagePlanProps) {
    super(scope, id, {
      apiStages: [{
        api: props.documentApi,
        stage: props.documentApi.deploymentStage,
      }],
      throttle: {
        rateLimit: 5,
        burstLimit: 10,
      },
    });
    props.devApiKeys.forEach((key) => {
      this.addApiKey(key);
    });
  }
}