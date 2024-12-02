import { createApolloClient } from "@/apollo-client";
import { Box, Divider, Flex, Text } from "@/libs/chakra-ui/react";
import dayjs from "@/libs/dayjs";
import { FORM_HEIGHT } from "./constants";
import { ChatForm } from "./components/chat-form";
import { MessagesByDateDocument } from "@/gql/graphql";

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
      {messagesByDate.map(({ id, createdAt, messages }) => {
        return (
          <Box key={id}>
            <Flex alignItems="center">
              <Divider />
              <Text px="16px" whiteSpace="nowrap">
                <time dateTime={dayjs(createdAt).format("YYYY-MM-DD")}>
                  {dayjs(createdAt).format("YYYY-MM-DD")}
                </time>
              </Text>
              <Divider />
            </Flex>
            {messages?.map(({ id, userID, createdAt, content }) => {
              return (
                <Box px="16px" key={id} pb="16px">
                  <Flex gap="8px">
                    <Text fontWeight={700}>{userID}</Text>
                    {/* TODO: make time fixed using position: sticky or fixed */}
                    <time dateTime={dayjs(createdAt).format("HH:mm")}>
                      {dayjs(createdAt).format("HH:mm")}
                    </time>
                  </Flex>
                  <Text>{content}</Text>
                </Box>
              );
            })}
          </Box>
        );
      })}
      <ChatForm />
    </Box>
  );
}
