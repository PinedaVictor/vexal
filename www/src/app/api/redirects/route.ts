"use server";
import { NextResponse, NextRequest } from "next/server";
import { redirect } from "next/navigation";

export async function GET(request: NextRequest) {
  return NextResponse.json({ sup: true });
}

export const redirectUser = (token: string, userId: string) => {
  const callbackUrl = process.env.CALLBACK_URL;

  if (!callbackUrl || !token || !userId) {
    return NextResponse.json(
      { error: "Missing required parameters" },
      { status: 400 }
    );
  }
  const redirectUrl = `${callbackUrl}?token=${token}&userId=${userId}`;
  redirect(redirectUrl);
};
