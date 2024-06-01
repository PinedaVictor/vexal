import {z} from 'zod';

const envSchema = z.object({
  // GCP
  GCP_PROJECT_ID: z.string(),

  // Firebase Admin
  FB_ADMIN_CLIENT_EMAIL: z.string(),
  FB_ADMIN_PRIVATE_KEY: z.string(),

  // Secrets Manager
  SECRETS_MAN_PRIVATE_KEY: z.string(),
  SECRETS_MAN_EMAIL: z.string(),

  // Prompts
  SMZ_FILE: z.string(),
});

export const API_ENV = envSchema.parse(process.env);
