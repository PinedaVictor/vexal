import type {Response} from 'express';
import {Router} from 'express';
import {type VxReq} from '../middlewares';
import {getSecret} from '../libs/gcp/secrets-manager/config';
import {generateChatResponse} from '../libs/openai/openai';

export const dataRouter = Router();

type ReqFunc = <T>(req: VxReq, res: Response) => Promise<Response<T>>;

const smz = async (req: VxReq, res: Response) => {
  try {
    const secret = await getSecret(req, 'openai');
    console.log('THE secret in SMZ:', secret);
    if (!secret) {
      return res.status(400).json({error: 'no openai key provided'});
    }
    console.log('Getting opena response');
    const chatRes = await generateChatResponse(secret, req.body);
    console.log(chatRes);
    return res.status(200).json(chatRes.content);
  } catch (error) {
    console.error('Error creating smz response', error);
    return res.status(400).json(error);
  }
};

const postRoutes: {route: string; fn: ReqFunc}[] = [{route: '/smz', fn: smz}];
postRoutes.map((r: {route: string; fn: ReqFunc}) => {
  dataRouter.post(r.route, r.fn);
});
