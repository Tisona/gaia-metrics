// import { WebSocket } from 'ws';
import * as client from 'prom-client';

// const collectDefaultMetrics = client.collectDefaultMetrics;
// const Registry = client.Registry;
// const register = new Registry();

// collectDefaultMetrics({ register });

console.log('---->');


new client.Gauge({
  name: 'metric_name',
  help: 'metric_help',
  collect() {
    // Invoked when the registry collects its metrics' values.
    // This can be synchronous or it can return a promise/be an async function.
    Math.random();
  },
});

process.exit();

// const url = 'ws://localhost:26657/websocket';
// const wss = new WebSocket(url);

// wss.on('open', () => {
//   console.log('opened');

//   wss.send(JSON.stringify({
//     "jsonrpc": "2.0",
//     "method": "subscribe",
//     "id": 0,
//     "params": {
//       "query": "tm.event='NewBlock'"
//     }
//   }));
// });

// wss.on('message', (message: any) => {
//   const data = JSON.parse(message);

//   if (data.result.data) {
//     console.log(data.result.data.value.block.header);

//     new client.Histogram({
//       name: 'gaia',
//       help: '....',
//       labelNames: [
//         'method',
//         'route',
//         'code'
//       ],
//       buckets: [0.1, 0.3, 0.5, 0.7, 1, 3, 5, 7, 10]
//     })

//   }
// });
