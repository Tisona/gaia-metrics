import { WebSocket } from 'ws';

export class Ws {
  public init() {
    const url = 'ws://localhost:26657/websocket';
    const wss = new WebSocket(url);

    return wss;
  }
}
