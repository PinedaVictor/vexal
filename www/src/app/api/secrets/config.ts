import { SecretManagerServiceClient } from "@google-cloud/secret-manager";
import { API_ENV } from "../config";
// import { credential } from "firebase-admin";

const initSecretManager = () => {
  //   let client: SecretManagerServiceClient = {};
  console.log(API_ENV.GCP_PROJECT_ID);
  console.log(API_ENV.SECRETS_MAN_PRIVATE_KEY.replace(/\\n/g, "\n"));
  console.log(API_ENV.SECRETS_MAN_EMAIL);
  //   if (!client) {
  let client = {};
  try {
    client = new SecretManagerServiceClient({
      projectId: API_ENV.GCP_PROJECT_ID,
      privatekey: API_ENV.SECRETS_MAN_PRIVATE_KEY.replace(/\\n/g, "\n"),
      clientEmail: API_ENV.SECRETS_MAN_EMAIL,
    });
  } catch (error) {
    console.error("error init secrets manager:", error);
  }
  //   }
  //   console.log("SECRET client:", client);
  return client;
};

export const writeSecret = async () => {
  //   initSecretManager();
  //   const r = {
  //     parent: "emailplusSID",
  //     secretId: "OPENAI_API_KEY",
  //     secret: "THE SECRET VALUE",
  //   };
  //   const secretCreate = await client.createSecret({
  //     parent: "",
  //     secret: { labels: { sdk: "openai" }, name: "THE KEY NAME??" },
  //     secretId: "",
  //   });
  //   console.log("Secret create:", secretCreate);
};
