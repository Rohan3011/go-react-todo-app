import useSWR from "swr";
import {
  ActionIcon,
  Alert,
  Box,
  Group,
  List,
  ThemeIcon,
  Title,
} from "@mantine/core";
import AddTodo from "./components/AddTodo";
import { CheckCircleFillIcon } from "@primer/octicons-react";
import { CheckCircleIcon } from "@primer/octicons-react";
import { useEffect } from "react";

export interface Todo {
  _id: string;
  title: string;
  body: string;
  done: boolean;
}

export const END_POINT = "http://localhost:8080";

const fetcher = (url: string) =>
  fetch(`${END_POINT}/${url}`).then((res) => res.json());

function App() {
  const { data, mutate } = useSWR<Todo[]>("api/todos", fetcher);

  const markTodoAsDone = async (todo: Todo) => {
    const updated = await fetch(`${END_POINT}/api/todos/${todo._id}`, {
      method: "PATCH",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ done: true }),
    }).then((res) => res.json());
    mutate(updated);
  };

  useEffect(() => {
    console.log(data);
  });

  return (
    <Box
      sx={(theme) => ({
        padding: "24",
        width: "100%",
        maxWidth: "40rem",
        margin: "0 auto",
      })}
    >
      <Group mb={20} position="center">
        <Title order={1}>Go-React Todo Application</Title>
      </Group>
      <List spacing={"xs"} size="sm" mb={12} center>
        {data?.length ? (
          data.map((todo) => (
            <List.Item
              key={`todo_item__${todo._id}`}
              onClick={() => markTodoAsDone(todo)}
              icon={
                todo.done ? (
                  <ActionIcon color={"blue"} radius="xl">
                    <CheckCircleFillIcon size={20} />
                  </ActionIcon>
                ) : (
                  <ActionIcon color={"gray"} radius="xl">
                    <CheckCircleFillIcon size={20} />
                  </ActionIcon>
                )
              }
            >
              {todo.title}
            </List.Item>
          ))
        ) : (
          <Alert title="Todo List is empty!" color="yellow">
            todos will be listed here.
          </Alert>
        )}
      </List>
      <AddTodo mutate={mutate} />
    </Box>
  );
}

export default App;
