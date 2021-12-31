import * as path from 'path';
import { ITable } from 'aws-cdk-lib/aws-dynamodb';
import { Construct } from 'constructs';
import { GoBaseFunc } from '../share';

export interface CreateDocumentFuncProps {
  readonly documentTable: ITable;
}

export class CreateDocumentFunc extends GoBaseFunc {
  constructor(scope: Construct, id: string, props: CreateDocumentFuncProps) {
    const useCase = 'createDocument';
    super(scope, id, {
      useCase: useCase,
      context: `DocBox#${useCase}`,
      entry: path.join(__dirname, `../../../../cmd/${useCase}`),
      environment: {
        DOCUMENT_TABLE_NAME: props.documentTable.tableName,
      },
      memorySize: 1024,
    });
    props.documentTable.grantWriteData(this);
  }
}