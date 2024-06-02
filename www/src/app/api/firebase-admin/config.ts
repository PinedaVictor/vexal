"use server";
import { initializeApp, cert, getApp } from "firebase-admin/app";
import { getAuth } from "firebase-admin/auth";

const placeholderRandomKey =
  "-----BEGIN PRIVATE KEY-----\nMIIBOAIBAAJAb6SCGCg9YJbGxqQuZVgiTfwhcInjpyNwGd4CiZ9mQhIXqznJIn0n\n3WFqmmSUyp/8FbuQJ5P/9c31uEBMHemHrwIDAQABAkAZhur8LQ9Dq5YNy3KUV5+Y\nZ0MaIV09VTwHnhzEbP8LshvjLTxtkew8hxiYH8vs16YDq+VBUOQcdUOYwrckJGlB\nAiEAs/EUfRcuaMQD8qd0sKVeZCg3DWT4E32WaNxzmP57I2ECIQCe1P/OEmiQNqzW\naxj1rfwZZmu9l8/RHNABMtGjefuVDwIgVifhCn/V7b2kskNxgL69MjF7IoOssBBa\nuvyKsL9wECECIE5igefaL95UDVI4QnYkqhCC+lLz0+y4UnL0+H68GYvDAiAX3dYR\nwXynTwFUZ352Usl8+/Bmc7OFuucLzb7pSAiskA==\n-----END PRIVATE KEY-----\n";

const projectId = process.env.FB_ADMIN_PROJECT_ID;
const privateKey = process.env.FB_ADMIN_PRIVATE_KEY ?? placeholderRandomKey;
const clientEmail = process.env.FB_ADMIN_CLIENT_EMAIL;

// const initFirebaseAdmin = () => {
//   initializeApp({
//     credential: cert({
//       privateKey: privateKey.replace(/\\n/g, "\n"),
//       clientEmail,
//       projectId,
//     }),
//   });
// };

// Initialize Firebase
export const getFirebaseAdmin = () => {
  const app = getApp();
  if (!app) {
    return initializeApp({
      credential: cert({
        privateKey: privateKey.replace(/\\n/g, "\n"),
        clientEmail,
        projectId,
      }),
    });
  }
  return app;
};

export const validateToken = async (token: string) => {
  try {
    const tokenValid = await getAuth().verifyIdToken(token);
    console.log("token valid", tokenValid);
    return tokenValid;
  } catch (error) {
    console.error(error);
    return false;
  }
};
