import OpenAI from "openai";

const initOpenAI = (apiKey: string) => {
  return new OpenAI({
    apiKey: apiKey,
  });
};

// TODO: Add content param
export const generateChatResponse = async (apiKey: string, content: string) => {
  const oai = initOpenAI(apiKey);
  const chatCompletion = await oai.chat.completions.create({
    messages: [{ role: "user", content: content }],
    model: "gpt-3.5-turbo",
  });
  console.log("THe chatCompletion var:", chatCompletion);
  const aiResponse = chatCompletion.choices[0].message;
  console.log("THis is the generated reponse:", aiResponse);
  return aiResponse;
};
