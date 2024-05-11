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
import { redirect } from "next/navigation";

export default function SignInCLI() {
  const [loading, setLoading] = useState(true);
  const auth = getAuth(vexalFirebaseApp);

  useEffect(() => {
    const SignInOAuth2 = async () => {
      const provider = new GoogleAuthProvider();
      try {
        console.log("Redirecting user");
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
        console.log("The user:", user);
        console.log("Token:", token);
        setLoading(false);
        window.location.href = `http://localhost:8080/callback?token=${token}&userId=${user.uid}`;
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
              <p>Authentication successful</p>
              <button
                onClick={() => {
                  signOut(auth)
                    .then(() => {
                      // Sign-out successful.
                      console.log("User signed out");
                    })
                    .catch((error) => {
                      // An error happened.
                      console.log("Error signing out");
                    });
                }}
              >
                Sign Out
              </button>
            </>
          )}
        </div>
      </div>
    </>
  );
}
