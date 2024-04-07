import OpenAI from "openai";
import { ENV } from "@/lib/environment/environment";

const initOpenAI = (apiKey: string) => {
  return new OpenAI({
    apiKey: apiKey,
  });
};

// TODO: Add content param
export const generateChatResponse = async (apiKey: string) => {
  console.log("The testing openai api key:", ENV.OPENAI_API_KEY);
  console.log("The api key give from cli user:", apiKey);
  // TODO: This can likely be handled inside initOpenAI
  let oai;
  if (apiKey !== "") {
    oai = initOpenAI(apiKey);
  } else {
    oai = initOpenAI(ENV.OPENAI_API_KEY);
  }
  try {
    const chatCompletion = await oai.chat.completions.create({
      messages: [{ role: "user", content: "Say this is a test" }],
      model: "gpt-3.5-turbo",
    });
    console.log("THe chatCompletion var:", chatCompletion);
    const aiResponse = chatCompletion.choices[0].message;
    console.log("THis is the generated reponse:", aiResponse);
  } catch (error) {
    console.error("Error generating text: ", error);
  }
};

async function main() {}
