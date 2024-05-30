import type {Request, Response, NextFunction} from 'express';
import type {DecodedIdToken} from 'firebase-admin/auth';
import {validateToken} from '../libs/gcp/firebase-admin/config';

export interface VxReq extends Request {
  user?: DecodedIdToken;
}

export const authorize = async (
  req: VxReq,
  res: Response,
  next: NextFunction
) => {
  const token = req.headers.authorization;
  if (!token) {
    return res.status(403).json({sup: 'Unauthorized'});
  }
  const validToken = await validateToken(token);
  if (!validToken) {
    return res.status(403).json({sup: 'Unauthorized'});
  }
  req.user = validToken;
  return next();
};
