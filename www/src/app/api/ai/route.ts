"use server";
import { g } from "./openai";
// Initial route for ai based packages.

/**
 *
 * @param request /
 * @returns JSON: {sup: true}
 * Sample route using nextjs
 */
export async function GET(request: Request) {
  console.log("Hitting server at /api/ai");
  g("");
  return Response.json({ sup: true });
}
