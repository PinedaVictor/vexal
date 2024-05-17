"use server";

import { validateToken } from "./firebase-admin/config";

/**
 *
 * @param request /
 * @returns JSON: {sup: true}
 * Sample route using nextjs
 */
export async function GET(request: Request) {
  console.log("Hitting server at /api");
  const auth = request.headers.get("Authorization");
  console.log("req in api:", auth);
  if (!auth) return;
  validateToken(auth);
  return Response.json({ sup: true });
}
