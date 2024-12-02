"use client";

import { Box, Button, Flex, Input, Text } from "@/libs/chakra-ui/react";
import { FORM_HEIGHT } from "../../constants";
import { useHooks } from "./hooks";

export function ChatForm() {
  const { error, isPending, submitAction } = useHooks();

  return (
    <Box
      p="16px"
      position="fixed"
      bottom={0}
      backgroundColor="gray.900"
      w="100%"
      h={`${FORM_HEIGHT}px`}
    >
      <form action={submitAction}>
        <fieldset disabled={isPending}>
          <Flex gap="12px">
            <Input placeholder="Type a message" name="content" isRequired />
            <Button type="submit" isLoading={isPending}>
              Send
            </Button>
          </Flex>
        </fieldset>
      </form>
      {error ? <Text>{error}</Text> : null}
    </Box>
  );
}
