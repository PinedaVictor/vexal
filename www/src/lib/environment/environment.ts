import { z } from "zod";

const envShema = z.object({
  // SMZ: z.string(),
  OPENAI_API_KEY: z.string().default(""),
});

export const ENV = envShema.parse(process.env);
