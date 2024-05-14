"use client";
import { useState } from "react";
import {
  getAuth,
  GoogleAuthProvider,
  signInWithRedirect,
  signOut,
  onAuthStateChanged,
} from "firebase/auth";
import { vexalFirebaseApp } from "@/app/firebase";
import { useEffect } from "react";
import { redirectUser } from "@/app/api/redirects/route";

export default function SignInCLI() {
  const [loading, setLoading] = useState(true);
  const auth = getAuth(vexalFirebaseApp);

  useEffect(() => {
    const SignInOAuth2 = async () => {
      const provider = new GoogleAuthProvider();
      try {
        signInWithRedirect(auth, provider);
      } catch (error) {
        console.log("Error redirecting user:", error);
      }
    };

    // Listen for the redirect result
    const unregisterAuthObserver = onAuthStateChanged(auth, async (user) => {
      if (user) {
        // User is signed in
        const token = await user.getIdToken();
        setLoading(false);
        redirectUser(token, user.uid);
      } else {
        // No user signed in, continue authentication flow
        SignInOAuth2();
      }
    });
    return () => {
      unregisterAuthObserver();
    };
  }, [auth]);

  return (
    <>
      <div className="flex h-screen">
        <div className="m-auto">
          {loading ? (
            <p>Authenticating...</p>
          ) : (
            <>
              <p>Redirecting you</p>
            </>
          )}
        </div>
      </div>
    </>
  );
}
