import OpenAI from 'openai';
import {API_ENV} from '../../env';

const initOpenAI = (apiKey: string) => {
  return new OpenAI({
    apiKey: apiKey,
  });
};

export const generateChatResponse = async (apiKey: string, input: string) => {
  console.log('OPENAI RESP');
  // console.log('The api key', apiKey);
  // console.log('The user input:', JSON.stringify(input));
  const oai = initOpenAI(apiKey);
  const prompt = `${API_ENV.SMZ_FILE} ${JSON.stringify(input)}`;
  const chatCompletion = await oai.chat.completions.create({
    messages: [{role: 'user', content: prompt}],
    model: 'gpt-3.5-turbo',
  });
  // console.log('THe chatCompletion var:', chatCompletion);
  const aiResponse = chatCompletion.choices[0].message;
  console.log('THis is the generated reponse:', aiResponse);
  return aiResponse;
};
