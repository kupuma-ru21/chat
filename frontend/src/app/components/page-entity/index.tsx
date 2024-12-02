import { createApolloClient } from "@/apollo-client";
import { Box } from "@/libs/chakra-ui/react";
import dayjs from "@/libs/dayjs";
import { FORM_HEIGHT } from "./constants";
import { MessagesByDateDocument } from "@/gql/graphql";
import { MessagesByDate } from "./components/messages_by_date";

export async function PageEntity() {
  const client = createApolloClient();
  const { data, error } = await client.query({ query: MessagesByDateDocument });
  if (error) throw new Error(error.message);

  const messagesByDate = data.messagesByDate.map(
    ({ id, createdAt, messages }) => {
      return {
        id,
        createdAt,
        messages:
          // TODO: remove this sort and sort on the server
          messages?.slice().sort((a, b) => {
            return dayjs(a.createdAt).diff(dayjs(b.createdAt));
          }) ?? [],
      };
    }
  );
  return (
    <Box position="relative" pb={`${FORM_HEIGHT}px`}>
      <MessagesByDate messagesByDate={messagesByDate} />
    </Box>
  );
}
