import { RemovalPolicy } from 'aws-cdk-lib';
import { AccessLogFormat, EndpointType, LogGroupLogDestination, RestApi } from 'aws-cdk-lib/aws-apigateway';
import { LogGroup, RetentionDays } from 'aws-cdk-lib/aws-logs';
import { Construct } from 'constructs';

export class DocumentApi extends RestApi {
  constructor(scope: Construct, id: string) {
    super(scope, id, {
      endpointConfiguration: {
        types: [EndpointType.REGIONAL],
      },
      deployOptions: {
        stageName: 'api',
        accessLogDestination: new LogGroupLogDestination(new LogGroup(scope, 'DocBoxAccessLogGroup', {
          retention: RetentionDays.THREE_MONTHS,
          removalPolicy: RemovalPolicy.DESTROY,
        })),
        accessLogFormat: AccessLogFormat.jsonWithStandardFields(),
      },
    });
  }
}