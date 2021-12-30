import * as path from 'path';
import { Construct } from 'constructs';
import { GoBaseFunc } from '../share';

export class CreateDocumentFunc extends GoBaseFunc {
  constructor(scope: Construct, id: string) {
    super(scope, id, {
      functionName: 'doc-box-createDocument',
      applicationContext: 'DocBox#CreateDocument',
      entry: path.join(__dirname, '../../../../cmd/createDocument'),
    });
  }
}