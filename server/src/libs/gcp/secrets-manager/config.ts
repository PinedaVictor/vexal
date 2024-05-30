import {SecretManagerServiceClient} from '@google-cloud/secret-manager';
import {API_ENV} from '../../../env';
import {VxReq} from '../../../middlewares';

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

export const getSecret = async (req: VxReq, sdk: string) => {
  const client = initSecretManager();
  try {
    const secretVersion = await client.accessSecretVersion({
      name: `projects/${API_ENV.GCP_PROJECT_ID}/secrets/${req.user?.uid}_${sdk}/versions/latest`,
    });
    // console.log('The secret:', secretVersion[0].payload?.data?.toString());
    return secretVersion[0].payload?.data?.toString();
  } catch (error) {
    console.error(error);
    // return 'secret not found';
  }
};

// TODO: Likely ready to delete
// export const writeSecret = async (secretId: string) => {
//   const client = initSecretManager();
//   try {
//     const secretCreate = await client.createSecret({
//       parent: `projects/${API_ENV.GCP_PROJECT_ID}`,
//       secret: {
//         labels: {sdk: 'openai'},
//         // name: 'topsecretValue',
//         replication: {
//           userManaged: {
//             replicas: [{location: 'us-central1'}],
//           },
//         },
//       },
//       // This will be the sid+[SDK-SUFFIX]
//       secretId: secretId,
//     });
//     console.log('Secret create:', secretCreate);
//   } catch (error) {
//     console.error('error creating secret:', error);
//   }
// };
