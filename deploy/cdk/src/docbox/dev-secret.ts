import { RemovalPolicy } from 'aws-cdk-lib';
import { Key } from 'aws-cdk-lib/aws-kms';
import { Secret } from 'aws-cdk-lib/aws-secretsmanager';
import { Construct } from 'constructs';

export class DevSecret extends Secret {
  constructor(scope: Construct, id: string) {
    super(scope, id, {
      removalPolicy: RemovalPolicy.RETAIN,
      encryptionKey: new Key(scope, 'DevSecretKey', {
        removalPolicy: RemovalPolicy.RETAIN,
      }),
      generateSecretString: {
        secretStringTemplate: JSON.stringify({
          user: 'test123456',
        }),
        generateStringKey: 'api_key',
        excludePunctuation: true,
        passwordLength: 64,
      },
    });
  }
}