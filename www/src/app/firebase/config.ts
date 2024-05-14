"use client";
// Import the functions you need from the SDKs you need
import { initializeApp, getApps, getApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: process.env.NEXT_PUBLIC_FB_apiKey,
  authDomain: process.env.NEXT_PUBLIC_FB_authDomain,
  projectId: process.env.NEXT_PUBLIC_FB_projectId,
  storageBucket: process.env.NEXT_PUBLIC_FB_storageBucket,
  messagingSenderId: process.env.NEXT_PUBLIC_FB_messagingSenderId,
  appId: process.env.NEXT_PUBLIC_FB_appId,
  measurementId: process.env.NEXT_PUBLIC_FB_measurementId,
};

// Initialize Firebase
const initFirebase = () => {
  if (getApps().length === 0) {
    return initializeApp(firebaseConfig);
  }
  return getApp();
};

export const vexalFirebaseApp = initFirebase();
// const analytics = getAnalytics(vexalFirebaseApp);
