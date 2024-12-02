"use server";

import { createApolloClient } from "@/apollo-client";
import { LoginDocument } from "@/gql/graphql";
import { cookies } from "next/headers";

export async function setCookie() {
  const cookieStore = await cookies();
  const userId = cookieStore.get("userId");
  if (userId) return;

  const client = createApolloClient();
  const { data, errors } = await client.mutate({ mutation: LoginDocument });
  if (errors) throw new Error(errors[0].message);
  if (data == null) throw new Error("data is null");

  cookieStore.set("userId", data.login);
}
