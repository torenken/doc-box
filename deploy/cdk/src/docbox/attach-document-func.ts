import * as path from 'path';
import { Construct } from 'constructs';
import { GoBaseFunc } from '../share';

export class AttachDocumentFunc extends GoBaseFunc {
  constructor(scope: Construct, id: string) {
    const useCase = 'attachDocument';
    super(scope, id, {
      useCase: useCase,
      context: `DocBox#${useCase}`,
      entry: path.join(__dirname, `../../../../cmd/${useCase}`),
      memorySize: 1024,
    });
  }
}