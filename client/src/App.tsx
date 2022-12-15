import useSWR from "swr";
import { Alert, Box, Group, List, ThemeIcon, Title } from "@mantine/core";
import AddTodo from "./components/AddTodo";
import { CheckCircleFillIcon } from "@primer/octicons-react";
import { CheckCircleIcon } from "@primer/octicons-react";

export interface Todo {
  id: number;
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
    const updated = await fetch(`${END_POINT}/api/todos/${todo.id}/done`, {
      method: "PATCH",
    }).then((res) => res.json());
    mutate(updated);
  };

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
              key={`todo_item__${todo.id}`}
              onClick={() => markTodoAsDone(todo)}
              icon={
                todo.done ? (
                  <ThemeIcon color={"blue"} radius="xl">
                    <CheckCircleFillIcon size={20} />
                  </ThemeIcon>
                ) : (
                  <ThemeIcon color={"gray"} radius="xl">
                    <CheckCircleFillIcon size={20} />
                  </ThemeIcon>
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
