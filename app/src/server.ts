import http from 'http';
import moment from 'moment';

import { gauge, registry } from './prometheus';
import { request } from './request';

const hostname = '127.0.0.1';
const port = 3000;

const server = http.createServer(async (_req, res) => {
  let latest_block_height = -1;
  let desync = -1;
  let n_peers = -1;

  try {
    const status = await request('status');
    const net_info = await request('net_info');


    const data = status.data.result.sync_info;

    const now = moment().unix();
    const latest_block_time = moment(data.latest_block_time).unix();

    desync = now - latest_block_time;
    n_peers = Number(net_info.data.result.n_peers);
    latest_block_height = Number(data.latest_block_height);
  } catch (e: any) {
    console.log(e.code);
  }

  gauge.latest_block_height.set(latest_block_height);
  gauge.n_peers.set(n_peers);
  gauge.desync.set(desync);

  const m = await registry.metrics();
  res.statusCode = 200;
  res.end(m);
});

server.listen(port, hostname, () => {
  console.log(`Server is running at http://${hostname}:${port}/`);
});
