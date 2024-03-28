import { z } from "zod";

const envShema = z.object({
  SMZ: z.string(),
});

export const ENV = envShema.parse(process.env);
