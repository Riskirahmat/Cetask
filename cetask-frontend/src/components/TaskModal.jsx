import { useState } from "react";
import {
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalBody,
  ModalFooter,
  Button,
  Input,
  VStack,
  HStack,
  IconButton,
  Text,
} from "@chakra-ui/react";
import { CloseIcon, DeleteIcon } from "@chakra-ui/icons";
import { updateTask, deleteTask } from "../services/taskService";
import ReactQuill from "react-quill";
import "react-quill/dist/quill.snow.css";

const TaskModal = ({
  isOpen,
  onClose,
  projectId,
  columnId,
  task,
  onUpdate,
  onDelete,
  onCreate,
}) => {
  const [title, setTitle] = useState(task?.title || "");
  const [desc, setDesc] = useState(task?.desc || "");
  const [isSaving, setIsSaving] = useState(false);

  const handleSave = async () => {
    if (!title.trim()) return;
    setIsSaving(true);
    try {
      if (task) {
        await handleUpdate();
      } else {
        await onCreate(title, desc);
      }
      onClose();
    } catch (err) {
      console.error("Failed to save task:", err);
    } finally {
      setIsSaving(false);
    }
  };

  const handleUpdate = async () => {
    if (!title.trim()) return;
    if (!task?.id || !columnId) {
      console.error("Invalid task or column ID:", {
        columnId,
        taskId: task?.id,
      });
      return;
    }
    setIsSaving(true);
    try {
      const updatedTask = await updateTask(
        projectId,
        columnId,
        task.id,
        title,
        desc
      );
      if (onUpdate) onUpdate(updatedTask);
      onClose();
    } catch (err) {
      console.error("Failed to update task:", err);
    } finally {
      setIsSaving(false);
    }
  };

  const handleDelete = async () => {
    if (!task?.id || !columnId) return;
    const confirmDelete = window.confirm(
      "Are you sure you want to delete this task?"
    );
    if (!confirmDelete) return;
    try {
      await deleteTask(projectId, columnId, task.id);
      if (onDelete) onDelete(task.id);
      onClose();
    } catch (err) {
      console.error("Failed to delete task:", err);
    }
  };

  return (
    <Modal isOpen={isOpen} onClose={onClose} isCentered>
      <ModalOverlay />
      <ModalContent
        p="4"
        border="2px solid black"
        boxShadow="4px 4px 0px black"
        bg="white"
      >
        <ModalHeader>
          <HStack justify="space-between">
            <Input
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              placeholder="Enter task title"
            />
            <IconButton icon={<CloseIcon />} onClick={onClose} size="sm" />
          </HStack>
        </ModalHeader>
        <ModalBody>
          <VStack spacing="4" align="stretch">
            <Text fontSize="sm" fontWeight="bold">
              Description
            </Text>
            <ReactQuill
              value={desc}
              onChange={setDesc}
              theme="snow"
              placeholder="Write a description..."
            />
          </VStack>
        </ModalBody>
        <ModalFooter>
          <HStack justify="space-between" w="100%">
            {task && (
              <Button
                colorScheme="red"
                leftIcon={<DeleteIcon />}
                onClick={handleDelete}
              >
                Delete
              </Button>
            )}
            <Button
              colorScheme="blue"
              onClick={handleSave}
              isLoading={isSaving}
            >
              {task ? "Save" : "Add Task"}
            </Button>
          </HStack>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};

export default TaskModal;
