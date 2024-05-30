import type {Response} from 'express';
import {Router} from 'express';
import {type VxReq} from '../middlewares';
import {getSecret} from '../libs/gcp/secrets-manager/config';

export const dataRouter = Router();

type ReqFunc = <T>(req: VxReq, res: Response) => Promise<Response<T>>;

const smz = async (req: VxReq, res: Response) => {
  console.log('Reaching smz route');
  console.log('Req headers:', req.headers);
  console.log('User id in SMZ req:', req.user);
  const secret = await getSecret(req, 'openai');
  console.log('THE secret in SMZ:', secret);
  return res.status(200).json({sup: true});
};

const postRoutes: {route: string; fn: ReqFunc}[] = [{route: '/smz', fn: smz}];
postRoutes.map((r: {route: string; fn: ReqFunc}) => {
  dataRouter.post(r.route, r.fn);
});
