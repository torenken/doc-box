import { RemovalPolicy } from 'aws-cdk-lib';
import { OAuthScope, UserPool } from 'aws-cdk-lib/aws-cognito';
import { Construct } from 'constructs';

export class DocumentUserPool extends UserPool {
  constructor(scope: Construct, id: string) {
    super(scope, id, {
      removalPolicy: RemovalPolicy.DESTROY,
    });

    this.addDomain('DocumentDomain', {
      cognitoDomain: {
        domainPrefix: 'docbox',
      },
    });

    this.addResourceServer('DocumentResourceServer', {
      identifier: 'subscriber',
      scopes: [{
        scopeName: 'docbox',
        scopeDescription: 'docbox-endpoint',
      }],
    });
    this.addClient('DocumentClient', {
      generateSecret: true,
      oAuth: {
        flows: {
          clientCredentials: true,
        },
        scopes: [OAuthScope.custom('subscriber/docbox')],
      },
    });
  }
}
