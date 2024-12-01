// REF: https://www.apollographql.com/blog/next-js-getting-started
import { ApolloClient, InMemoryCache } from "@apollo/client";

export const createApolloClient = () => {
  return new ApolloClient({
    uri: process.env.GRAPHQL_SCHEMA_URL,
    cache: new InMemoryCache(),
  });
};
