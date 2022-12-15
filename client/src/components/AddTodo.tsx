import { Button, Group, Modal, Textarea, TextInput } from "@mantine/core";
import { useForm } from "@mantine/form";
import { useDisclosure } from "@mantine/hooks";
import { KeyedMutator } from "swr";
import { END_POINT, Todo } from "../App";

function AddTodo({ mutate }: { mutate: KeyedMutator<Todo[]> }) {
  const [opened, handlers] = useDisclosure(false);

  const form = useForm({
    initialValues: {
      title: "",
      body: "",
    },
  });

  const createTodo = async (data: { title: string; body: string }) => {
    const updated = await fetch(`${END_POINT}/api/todos`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    }).then((res) => res.json());
    mutate(updated);

    form.reset();
    handlers.close();
  };

  return (
    <>
      <Modal
        centered
        opened={opened}
        onClose={handlers.close}
        title={"Create todo ✏️"}
      >
        <form onSubmit={form.onSubmit(createTodo)} onReset={form.onReset}>
          <TextInput
            label="Title"
            placeholder="Title"
            required
            withAsterisk
            mb={12}
            {...form.getInputProps("title")}
          />
          <Textarea
            mb={12}
            placeholder="Todo Description"
            label="Description"
            {...form.getInputProps("body")}
          />
          <Group position="right">
            <Button color="gray" type="reset">
              Reset
            </Button>
            <Button type="submit">Create</Button>
          </Group>
        </form>
      </Modal>
      <Button fullWidth onClick={handlers.open}>
        Create Todo
      </Button>
    </>
  );
}

export default AddTodo;
