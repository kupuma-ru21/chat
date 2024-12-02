"use client";
import { Box, Divider, Flex, Text } from "@/libs/chakra-ui/react";
import dayjs from "@/libs/dayjs";
import { MessagesByDateQuery } from "@/gql/graphql";
import { useHooks } from "./hooks";
import { ChatForm } from "./components/chat-form";

export function MessagesByDate({
  messagesByDate,
}: {
  messagesByDate: MessagesByDateQuery["messagesByDate"];
}) {
  useHooks();
  return (
    <>
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
    </>
  );
}
