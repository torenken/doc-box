import * as path from 'path';
import { ITable } from 'aws-cdk-lib/aws-dynamodb';
import { IBucket } from 'aws-cdk-lib/aws-s3';
import { Construct } from 'constructs';
import { GoBaseFunc } from './go-base-func';

export interface AttachDocumentFuncProps {
  readonly documentBucket: IBucket;
  readonly documentTable: ITable;
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
        DOCUMENT_TABLE_NAME: props.documentTable.tableName,
      },
    });
    props.documentBucket.grantReadWrite(this);
    props.documentTable.grantReadWriteData(this);
  }
}