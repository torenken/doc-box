import { Stack, StackProps } from 'aws-cdk-lib';
import { Construct } from 'constructs';
import { CreateDocumentFunc } from './doc/create-document-func';

export interface DocBoxStackProps extends StackProps {
}

export class DocBoxStack extends Stack {
  constructor(scope: Construct, id: string, props: DocBoxStackProps = {}) {
    super(scope, id, props);

    // The code that defines your stack goes here

    new CreateDocumentFunc(this, 'CreateDocumentFunc');
  }
}