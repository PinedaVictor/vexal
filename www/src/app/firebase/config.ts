"use client";
// Import the functions you need from the SDKs you need
import { initializeApp, getApps, getApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyBX9VtsHErsdfIQtOwiC-lJCERBN66yeds",
  authDomain: "vexal-9a246.firebaseapp.com",
  projectId: "vexal-9a246",
  storageBucket: "vexal-9a246.appspot.com",
  messagingSenderId: "435966026827",
  appId: "1:435966026827:web:38cbed5ef93e6d6a7899ce",
  measurementId: "G-W9ZG10X3DZ",
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
