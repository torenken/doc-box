import { Stack, StackProps } from 'aws-cdk-lib';
import { LambdaIntegration } from 'aws-cdk-lib/aws-apigateway';
import { Construct } from 'constructs';
import { AttachDocumentFunc, CreateDocumentFunc, DocumentApi, DocumentBucket, DocumentTable } from './docbox';

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
    const documentBucket = new DocumentBucket(this, 'DocumentBucket', {
      bucketName: props.documentBucketName,
    });

    const documentTable = new DocumentTable(this, 'DocumentTable');

    const createDocumentFunc = new CreateDocumentFunc(this, 'CreateDocumentFunc', {
      documentTable,
    });

    const attachDocumentFunc = new AttachDocumentFunc(this, 'AttachDocumentFunc', {
      documentBucket,
    });

    const documentApi = new DocumentApi(this, 'DocumentRestApi');
    const documentResource = documentApi.root.addResource('documentManagement').addResource('document');
    const showDocumentResource = documentResource.addResource('{docId}');
    const attachmentResource = showDocumentResource.addResource('attachment');

    documentResource.addMethod('POST', new LambdaIntegration(createDocumentFunc));
    attachmentResource.addMethod('POST', new LambdaIntegration(attachDocumentFunc));
  }
}