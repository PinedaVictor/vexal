"use server";

import { validateToken } from "./firebase-admin/config";
import { writeSecret } from "./secrets/config";
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
  writeSecret();
  // if (!auth) return;
  // validateToken(auth);
  return Response.json({ sup: true });
}
