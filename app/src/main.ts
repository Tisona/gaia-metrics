import { WebSocket } from 'ws';
import moment from 'moment';

const url = 'ws://localhost:26657/websocket';
const wss = new WebSocket(url);

wss.on('open', () => {
  console.log('opened');

  wss.send(JSON.stringify({
    "jsonrpc": "2.0",
    "method": "subscribe",
    "id": 0,
    "params": {
      "query": "tm.event='NewBlock'"
    }
  }));
});

wss.on('message', (message: any) => {
  const data = JSON.parse(message);

  if (data.result.data) {
    const header = data.result.data.value.block.header;
    const time = moment(header.time);
    const diff = moment().diff(time, "minutes");

    console.log(header);
    console.log(diff);
  }
});
