"use server";
import { generateChatResponse } from "./openai";
// Initial route for ai based packages.

/**
 *
 * @param request /
 * @returns JSON: {sup: true}
 * Sample route using nextjs
 */
export async function GET(request: Request) {
  console.log("Hitting server at /api/ai");

  return Response.json({ sup: true });
}
