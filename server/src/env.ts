import {z} from 'zod';

const envSchema = z.object({
  // OAuth2.0
  // CALLBACK_URL: z.string(),

  // // Firebase
  // NEXT_PUBLIC_FB_apiKey: z.string(),
  // NEXT_PUBLIC_FB_authDomain: z.string(),
  // NEXT_PUBLIC_FB_projectId: z.string(),
  // NEXT_PUBLIC_FB_storageBucket: z.string(),
  // NEXT_PUBLIC_FB_messagingSenderId: z.string(),
  // NEXT_PUBLIC_FB_appId: z.string(),
  // NEXT_PUBLIC_FB_measurementId: z.string(),

  // // Firebase Admin
  // FB_ADMIN_PROJECT_ID: z.string(),
  // FB_ADMIN_CLIENT_EMAIL: z.string(),
  // FB_ADMIN_PRIVATE_KEY: z.string(),

  // GCP
  GCP_PROJECT_ID: z.string(),

  // Secrets Manager
  SECRETS_MAN_PRIVATE_KEY: z.string(),
  SECRETS_MAN_EMAIL: z.string(),
});

export const API_ENV = envSchema.parse(process.env);
