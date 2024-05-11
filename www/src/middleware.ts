"use server";
import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";
import { headers } from "next/headers";
// import { validateToken } from "./app/api/firebase-admin/config";
// import { vexalFirebaseApp } from "./app/firebase";
// import admin from "firebase-admin";

// This function can be marked `async` if using `await` inside
export async function middleware(request: NextRequest) {
  // console.log("Middleware function", request.headers.get("Firebase"));
  const authorization = headers().get("Authorization");
  console.log("Requests header Auth:", authorization);
  if (authorization === null) {
    return NextResponse.redirect(new URL("/", request.url));
  }
  // const auth = getAuth(vexalFirebaseApp);
  // console.log(auth);
  try {
    // const tokenValid = await validateToken(authorization);
    // console.log("Token Valid:", tokenValid);
  } catch (error) {
    console.error(error);
  }
}

// See "Matching Paths" below to learn more
export const config = {
  matcher: "/api/:path*",
};
