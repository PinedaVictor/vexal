import type {Request, Response, NextFunction} from 'express';
import {validateToken} from '../libs/gcp/firebase-admin/config';

export const authorize = async (
  req: Request,
  res: Response,
  next: NextFunction
) => {
  const token = req.headers.authorization;
  if (!token) {
    return res.status(403).json({sup: 'Unauthorized'});
  }
  const validToken = await validateToken(token);
};
