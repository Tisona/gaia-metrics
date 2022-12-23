import * as client from 'prom-client';

const Registry = client.Registry;

export const registry = new Registry();
export const gauge = {
  latest_block_height: new client.Gauge({
    name: 'latest_block_height',
    help: 'latest_block_height',
    registers: [registry],
  }),
  n_peers: new client.Gauge({
    name: 'n_peers',
    help: 'n_peers',
    registers: [registry],
  }),
  desync: new client.Gauge({
    name: 'desync',
    help: 'desync',
    registers: [registry],
  }),
};
