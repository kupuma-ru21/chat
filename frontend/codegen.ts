import type { CodegenConfig } from "@graphql-codegen/cli";

const config: CodegenConfig = {
  schema: process.env.GRAPHQL_SCHEMA_URL,
  documents: ["**/*.graphql"],
  ignoreNoDocuments: true, // for better experience with the watcher
  generates: {
    "./src/gql/": {
      preset: "client",
      plugins: [],
      presetConfig: {
        // ref: https://the-guild.dev/graphql/codegen/plugins/presets/preset-client#the-usefragment-helper
        fragmentMasking: { unmaskFunctionName: "getFragmentData" },
      },
    },
  },
};

export default config;
