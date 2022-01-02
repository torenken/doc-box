import * as path from 'path';
import { IBucket } from 'aws-cdk-lib/aws-s3';
import { Construct } from 'constructs';
import { GoBaseFunc } from '../share';

export interface AttachDocumentFuncProps {
  readonly documentBucket: IBucket;
}

export class AttachDocumentFunc extends GoBaseFunc {
  constructor(scope: Construct, id: string, props: AttachDocumentFuncProps) {
    const useCase = 'attachDocument';
    super(scope, id, {
      useCase: useCase,
      context: `DocBox#${useCase}`,
      entry: path.join(__dirname, `../../../../cmd/${useCase}`),
      memorySize: 1024,
      environment: {
        DOCUMENT_STORAGE_NAME: props.documentBucket.bucketName,
      },
    });
    props.documentBucket.grantWrite(this);
  }
}