import * as client from 'prom-client';

const Registry = client.Registry;

export const registry = new Registry();
export const gauge = {
  'a': new client.Gauge({
    name: 'gaia_a',
    help: 'gaia_a',
    registers: [registry],
  }),
  'b': new client.Gauge({
    name: 'gaia_b',
    help: 'gaia_b',
    registers: [registry],
  }),
  'c': new client.Gauge({
    name: 'gaia_c',
    help: 'gaia_c',
    registers: [registry],
  }),
};
