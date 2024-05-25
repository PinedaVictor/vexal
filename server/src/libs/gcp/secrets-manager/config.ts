import {SecretManagerServiceClient} from '@google-cloud/secret-manager';
import {API_ENV} from '../../../env';

let client: SecretManagerServiceClient;
const initSecretManager = () => {
  if (!client) {
    try {
      client = new SecretManagerServiceClient({
        credentials: {
          projectId: API_ENV.GCP_PROJECT_ID,
          private_key: API_ENV.SECRETS_MAN_PRIVATE_KEY.replace(/\\n/g, '\n'),
          client_email: API_ENV.SECRETS_MAN_EMAIL,
        },
      });
    } catch (error) {
      console.error('error initiating secrets manager:', error);
    }
  }
  return client;
};

export const writeSecret = async () => {
  const client = initSecretManager();
  try {
    const secretCreate = await client.createSecret({
      parent: `projects/${API_ENV.GCP_PROJECT_ID}`,
      secret: {
        labels: {sdk: 'openai'},
        name: 'topsecretValue',
        replication: {
          userManaged: {
            replicas: [{location: 'us-central1'}],
          },
        },
      },
      // This will be the sid+[SDK-SUFFIX]
      secretId: 'topsecret',
    });
  } catch (error) {
    console.error('error creating secret:', error);
  }
};
