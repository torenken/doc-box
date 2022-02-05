import { SecretValue, Stack, StackProps } from 'aws-cdk-lib';
import { AuthorizationType, CognitoUserPoolsAuthorizer, LambdaIntegration } from 'aws-cdk-lib/aws-apigateway';
import { Construct } from 'constructs';
import {
  AttachDocumentFunc,
  CreateDocumentFunc,
  DocumentApi,
  DocumentBucket,
  DevUsagePlan,
  DocumentTable,
  DocumentUserPool,
  DevSecret,
} from './docbox';

export interface DocBoxStackProps extends StackProps {
  /**
   * Name of the Bucket for document attachments
   */
  readonly documentBucketName: string;
}

export class DocBoxStack extends Stack {
  constructor(scope: Construct, id: string, props: DocBoxStackProps) {
    super(scope, id, props);

    // The code that defines your stack goes here

    const devSecret = new DevSecret(this, 'DevSecret');

    const documentBucket = new DocumentBucket(this, 'DocumentBucket', {
      bucketName: props.documentBucketName,
    });

    const documentTable = new DocumentTable(this, 'DocumentTable');

    const createDocumentFunc = new CreateDocumentFunc(this, 'CreateDocumentFunc', {
      documentTable,
    });

    const attachDocumentFunc = new AttachDocumentFunc(this, 'AttachDocumentFunc', {
      documentBucket,
      documentTable,
    });

    const documentUserPool = new DocumentUserPool(this, 'DocumentUserPool');

    const cognitoUserPoolsAuthorizer = new CognitoUserPoolsAuthorizer(this, 'DocumentCognitoUserPoolsAuthorizer', {
      cognitoUserPools: [documentUserPool],
    });

    const documentApi = new DocumentApi(this, 'DocumentRestApi');

    new DevUsagePlan(this, 'DocumentDevUsagePlan', {
      documentApi,
      devApiKeys: [
        documentApi.addApiKey('DevApiKey', {
          value: SecretValue.secretsManager(devSecret.secretArn, { jsonField: 'api_key' }).toString(),
        }),
      ],
    });

    const documentResource = documentApi.root.addResource('documentManagement').addResource('document');
    const showDocumentResource = documentResource.addResource('{docId}');
    const attachmentResource = showDocumentResource.addResource('attachment');

    documentResource.addMethod('POST', new LambdaIntegration(createDocumentFunc), {
      authorizationScopes: ['subscriber/docbox'],
      authorizationType: AuthorizationType.COGNITO,
      authorizer: cognitoUserPoolsAuthorizer,
      apiKeyRequired: true,
    });
    attachmentResource.addMethod('POST', new LambdaIntegration(attachDocumentFunc));
  }
}