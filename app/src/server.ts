import http from 'http';
import { gauge, registry } from './prometheus';

const hostname = '127.0.0.1';
const port = 3000;

const server = http.createServer(async (_req, res) => {
  const m = await registry.metrics();

  gauge.a.set(Math.random());
  gauge.b.set(Math.random());
  gauge.c.set(Math.random());

  res.statusCode = 200;
  res.end(m);
});

server.listen(port, hostname, () => {
  console.log(`Server is running at http://${hostname}:${port}/`);
});
