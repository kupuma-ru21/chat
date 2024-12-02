import { extendTheme, type ThemeConfig } from "@/libs/chakra-ui/react";

const config: ThemeConfig = {
  initialColorMode: "system",
  useSystemColorMode: true,
};

export const styles = {
  global: { "html, body": { overscrollBehavior: "none" } },
};

export const theme = extendTheme({ styles, config });
