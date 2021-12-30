import { Stack, StackProps } from 'aws-cdk-lib';
import { Construct } from 'constructs';
import { CreateDocumentFunc, DocumentTable } from './docbox';

export interface DocBoxStackProps extends StackProps {
}

export class DocBoxStack extends Stack {
  constructor(scope: Construct, id: string, props: DocBoxStackProps = {}) {
    super(scope, id, props);

    // The code that defines your stack goes here
    const documentTable = new DocumentTable(this, 'DocumentTable');

    const createDocumentFunc = new CreateDocumentFunc(this, 'CreateDocumentFunc', {
      documentTable,
    });

    documentTable.grantWriteData(createDocumentFunc);
  }
}