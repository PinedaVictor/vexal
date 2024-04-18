import express, { Request, Response } from "express";

const app = express();
const port = 8080;

app.get("/", (req: Request, res: Response) => {
  return res.status(200).json({ sup: true });
});

app.get("/service", (req: Request, res: Response) => {
  return res.status(200).json({ sup: true });
});

app.get("/serviceB", (req: Request, res: Response) => {
  return res.status(200).json({ sup: true });
});

app.get("/service3", (req: Request, res: Response) => {
  return res.status(200).json({ sup: true });
});

app.listen(port, () => {
  // TODO: Update 8080 to port const
  console.log("API listening on port http://localhost:8080/");
});
