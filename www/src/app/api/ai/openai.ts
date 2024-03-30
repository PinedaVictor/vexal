import OpenAI from "openai";
import { ENV } from "@/lib/environment/environment";

const initOpenAI = (apiKey: string) => {
  return new OpenAI({
    apiKey: apiKey || ENV.OPENAI_API_KEY,
  });
};

export const g = async (apiKey: string) => {
  console.log("The openai api key:", ENV.OPENAI_API_KEY);
  const oai = initOpenAI(ENV.OPENAI_API_KEY);
  const chatCompletion = await oai.chat.completions.create({
    messages: [{ role: "user", content: "Say this is a test" }],
    model: "gpt-3.5-turbo",
  });
};

g("");

async function main() {}
