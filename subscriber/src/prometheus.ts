import * as client from 'prom-client';

export class Prometheus {
  public init() {
    const collectDefaultMetrics = client.collectDefaultMetrics;
    const Registry = client.Registry;
    const register = new Registry();
    collectDefaultMetrics({ register });

    return client;
  }
}
