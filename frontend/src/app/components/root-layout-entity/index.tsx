"use client";
import { useEffect } from "react";
import { setCookie } from "./actions/set-cookie";

export function RootLayoutEntity({ children }: { children: React.ReactNode }) {
  useEffect(() => {
    (async () => {
      await setCookie();
    })();
  }, []);
  return <>{children}</>;
}
