"use server";

import { createStreamableUI } from "ai/rsc";

import { NatsConnection, connect, StringCodec } from "nats";

export async function getWeather() {
  const weatherUI = createStreamableUI();
  async function server(subject: string) {
    let nc: NatsConnection | null = null;

    const connectionString = process.env.NATS ?? "nats://127.0.0.1:4222";
    nc = await connect({
      servers: [connectionString],
    });

    const sc = StringCodec();
    const sub = nc.subscribe("hello");

    (async () => {
      for await (const m of sub) {
        const msg = sc.decode(m.data);
        weatherUI.append(<div>{msg}</div>);
        if (msg === "done") {
          weatherUI.done(<div>It&apos;s a sunny day!</div>);
          break;
        }
      }
      console.log("subscription closed");
    })();
  }
  weatherUI.update(<div style={{ color: "gray" }}>Connecting</div>);
  server("hello");
  // setTimeout(() => {
  //   weatherUI.done(<div>It's a sunny day!</div>);
  // }, 3000);
  setTimeout(() => {
    weatherUI.append(<div>Preparing forecast</div>);
  }, 2000);
  setTimeout(() => {
    weatherUI.update(<div>Connected</div>);
  }, 200);

  return weatherUI.value;
}
