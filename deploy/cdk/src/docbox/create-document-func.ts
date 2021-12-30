import * as path from 'path';
import { ITable } from 'aws-cdk-lib/aws-dynamodb';
import { Construct } from 'constructs';
import { GoBaseFunc } from '../share';

export interface CreateDocumentFuncProps {
  readonly documentTable: ITable;
}

export class CreateDocumentFunc extends GoBaseFunc {
  constructor(scope: Construct, id: string, props: CreateDocumentFuncProps) {
    super(scope, id, {
      functionName: 'doc-box-createDocument',
      applicationContext: 'DocBox#CreateDocument',
      entry: path.join(__dirname, '../../../../cmd/createDocument'),
      environment: {
        DOCUMENT_TABLE_NAME: props.documentTable.tableName,
      },
    });
    props.documentTable.grantWriteData(this);
  }
}