import 'dotenv/config';
import express, {Request, Response} from 'express';
import {writeSecret} from './libs/gcp/secrets-manager/config';

const app = express();
const port = 4000 || process.env.PORT;

app.get('/', (req: Request, res: Response) => {
  return res.status(200).json({sup: true});
});

app.listen(port, () => {
  // TODO: Update 8080 to port const
  console.log(`API listening on port http://localhost:${port}`);
  // writeSecret();
});
