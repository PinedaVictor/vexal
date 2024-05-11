"use client";
import type { ReactNode } from "react";
import { DocsUI } from "@/components/atomic/templates/DocsUI";
import { vexalFirebaseApp } from "./firebase";
import { getAuth, signOut } from "firebase/auth";

export default function Home(props: { children: ReactNode }) {
  const auth = getAuth(vexalFirebaseApp);
  return (
    <main className=" bg-slate-100 h-screen">
      <DocsUI>
        <div>This is my docs UI</div>
        <button
          onClick={() => {
            signOut(auth)
              .then(() => {
                // Sign-out successful.
                console.log("User signed out");
              })
              .catch((error) => {
                // An error happened.
              });
          }}
        >
          Sign Out
        </button>
      </DocsUI>
    </main>
  );
}
