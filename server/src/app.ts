import 'dotenv/config';
import express, {Request, Response} from 'express';
import cors from 'cors';
import {authorize} from './middlewares';
import {dataRouter} from './routes';

const app = express();
const port = 4000 || process.env.PORT;
app.use(cors({origin: '*'}));
app.use(express.json());
app.use(express.urlencoded({extended: true}));

app.get('/', (req: Request, res: Response) => {
  return res.status(200).json({sup: true});
});

app.use(authorize);
app.use('/data', dataRouter);

app.listen(port, () => {
  console.log(`API listening on port http://localhost:${port}`);
});
