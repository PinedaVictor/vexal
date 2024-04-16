"use server";

import { NextRequest, NextResponse } from "next/server";
import { generateChatResponse } from "../openai";

export async function POST(request: NextRequest) {
  const key = request.headers.get("openai");
  if (!key) {
    return NextResponse.json({ error: "open api key not provided" });
  }
  const data = await request.json();
  try {
    const aiResponse = await generateChatResponse(
      key,
      `${process.env.SMZ_FILE} file: ${data.entity} content: ${data.content}`
    );
    console.log("This is the summary:", aiResponse?.content);
    return NextResponse.json(aiResponse.content);
  } catch (error) {
    console.error("error generating SMZ response");
    return NextResponse.json({ error: error });
  }
}
