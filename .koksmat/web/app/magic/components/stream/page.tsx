"use client";
import { useEffect, useState } from "react";
import { readStreamableValue } from "ai/rsc";
import { getWeather } from "./stream-server";

export default function Page() {
  const [weather, setWeather] = useState<any>(null);
  useEffect(() => {
    const load = async () => {
      const weatherUI = await getWeather();
      setWeather(weatherUI);
    };
    load();
  }, []);

  return <div>{weather}</div>;
}
