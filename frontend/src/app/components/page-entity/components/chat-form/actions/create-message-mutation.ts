"use server";

import { createApolloClient } from "@/apollo-client";
import { cookies } from "next/headers";
import { revalidatePath } from "next/cache";
import { CreateMessageDocument } from "@/gql/graphql";

export async function createMessageMutation(
  _prev: unknown,
  formData: FormData
) {
  const cookieStore = await cookies();
  const userId = cookieStore.get("userId");
  if (userId === undefined) throw new Error("userid not found");

  const client = createApolloClient();
  const { errors } = await client.mutate({
    mutation: CreateMessageDocument,
    variables: {
      userId: userId.value,
      content: String(formData.get("content")),
    },
  });

  if (errors) {
    console.log(errors);
    return { error: errors[0].message };
  }

  revalidatePath("/");
  return { error: null };
}
