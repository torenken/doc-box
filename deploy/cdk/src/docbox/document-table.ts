import { AttributeType, BillingMode, Table, TableEncryption } from 'aws-cdk-lib/aws-dynamodb';
import { Construct } from 'constructs';

export interface DocumentTableProps {
  /**
     * @default - Assigned by CloudFormation (recommended).
     */
  readonly tableName?: string;
}

export class DocumentTable extends Table {
  constructor(scope: Construct, id: string, props: DocumentTableProps = {}) {
    super(scope, id, {
      tableName: props.tableName,
      encryption: TableEncryption.DEFAULT,
      billingMode: BillingMode.PAY_PER_REQUEST,
      partitionKey: {
        name: 'id',
        type: AttributeType.STRING,
      },
    });
  }
}