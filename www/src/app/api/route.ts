"use server";

import { validateToken } from "./firebase-admin/config";
import { getSecret, writeSecret } from "./secrets/config";
import { headers } from "next/headers";
// import { initSecretManager } from "./secrets/config";

/**
 *
 * @param request /
 * @returns JSON: {sup: true}
 * Sample route using nextjs
 */
export async function GET(request: Request) {
  console.log("Hitting server at /api");
  // const auth = request.headers.get("Authorization");
  // console.log("req in api:", auth);

  try {
    const i = await getSecret("openai");
    console.log("secret: getSecret:::", i);
  } catch (error) {
    console.log("ERROR::", error);
  }
  // if (!auth) return;
  // validateToken(auth);
  return Response.json({ sup: true });
}
