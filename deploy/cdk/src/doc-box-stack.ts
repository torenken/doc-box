import { Stack, StackProps } from 'aws-cdk-lib';
import { LambdaIntegration } from 'aws-cdk-lib/aws-apigateway';
import { Construct } from 'constructs';
import { CreateDocumentFunc, DocumentApi, DocumentTable } from './docbox';

export interface DocBoxStackProps extends StackProps {}

export class DocBoxStack extends Stack {
  constructor(scope: Construct, id: string, props: DocBoxStackProps = {}) {
    super(scope, id, props);

    // The code that defines your stack goes here
    const documentTable = new DocumentTable(this, 'DocumentTable');

    const createDocumentFunc = new CreateDocumentFunc(this, 'CreateDocumentFunc', {
      documentTable,
    });

    const documentApi = new DocumentApi(this, 'DocumentRestApi');
    const documentResource = documentApi.root.addResource('documentManagement').addResource('document');

    documentResource.addMethod('POST', new LambdaIntegration(createDocumentFunc));
  }
}