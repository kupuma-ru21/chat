"use client";

import { theme } from "@/theme";
import { ChakraProvider } from "@/libs/chakra-ui/react";

export function Providers({ children }: { children: React.ReactNode }) {
  // TODO: when dark mode is enabled, white screen appears on page load
  return <ChakraProvider theme={theme}>{children}</ChakraProvider>;
}
