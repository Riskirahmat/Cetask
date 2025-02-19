import { useEffect, useState } from "react";
import { DragDropContext, Droppable, Draggable } from "react-beautiful-dnd";
import {
  Box,
  Button,
  HStack,
  VStack,
  Text,
  IconButton,
  Spinner,
  Input,
} from "@chakra-ui/react";
import { EditIcon, DeleteIcon, AddIcon } from "@chakra-ui/icons";
import { getTasks, moveTask } from "../services/taskService";
import { createColumn } from "../services/projectService"; 

const KanbanBoard = ({
  projectId,
  columns = [],
  setColumns,
  handleOpenEditColumnModal,
  handleOpenConfirmModal,
  handleOpenAddTaskModal,
  handleOpenEditTaskModal,
}) => {
  const [tasks, setTasks] = useState({});
  const [loading, setLoading] = useState(true);
  const [newColumnName, setNewColumnName] = useState(""); 

  useEffect(() => {
    const fetchTasks = async () => {
      if (!columns || columns.length === 0) {
        setLoading(false);
        return;
      }

      try {
        const tasksData = {};
        for (const column of columns) {
          tasksData[column.id] = await getTasks(projectId, column.id);
        }
        setTasks(tasksData);
      } catch (error) {
        console.error("Failed to fetch tasks:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchTasks();
  }, [projectId, columns]);

  const onDragEnd = async (result) => {
    if (!result.destination) return;
    const { source, destination, draggableId } = result;
    const taskId = draggableId;
    const oldColumnId = source.droppableId;
    const newColumnId = destination.droppableId;
    const newPosition = destination.index;

    if (oldColumnId === newColumnId && source.index === newPosition) return;

    try {
      await moveTask(projectId, newColumnId, taskId, newPosition);
      setTasks((prev) => {
        const updatedOldColumn =
          prev[oldColumnId]?.filter((task) => task.id !== taskId) || [];
        const movedTask = prev[oldColumnId]?.find((task) => task.id === taskId);
        if (!movedTask) return prev;

        movedTask.position = newPosition;
        movedTask.column_id = newColumnId;

        const updatedNewColumn = [...(prev[newColumnId] || []), movedTask];
        return {
          ...prev,
          [oldColumnId]: updatedOldColumn,
          [newColumnId]: updatedNewColumn.sort(
            (a, b) => a.position - b.position
          ),
        };
      });
    } catch (err) {
      console.error("Failed to move task:", err);
    }
  };

  const handleAddColumn = async () => {
    if (!newColumnName.trim()) return;
    try {
      const newColumn = await createColumn(projectId, newColumnName);
      setColumns((prev) => [...prev, newColumn]); 
      setNewColumnName(""); 
    } catch (error) {
      console.error("Failed to create column:", error);
    }
  };

  if (loading) {
    return (
      <Box textAlign="center" mt="4">
        <Spinner size="xl" />
        <Text mt="2">Loading columns...</Text>
      </Box>
    );
  }

  if (!columns || columns.length === 0) {
    return (
      <Box textAlign="center" mt="4">
        <Text>No columns available. Please add a column.</Text>
      </Box>
    );
  }

  return (
    <>
      <HStack mt="4" spacing="2">
        <Input
          placeholder="Enter column name"
          value={newColumnName}
          onChange={(e) => setNewColumnName(e.target.value)}
          w="250px"
        />
        <Button
          leftIcon={<AddIcon />}
          colorScheme="blue"
          onClick={handleAddColumn}
        >
          Add Column
        </Button>
      </HStack>

      <DragDropContext onDragEnd={onDragEnd}>
        <HStack spacing="4" overflowX="auto" mt="4">
          {columns.map((column) => (
            <Droppable key={column.id} droppableId={String(column.id)}>
              {(provided) => (
                <Box
                  {...provided.droppableProps}
                  ref={provided.innerRef}
                  p="4"
                  w="250px"
                  minH="300px"
                  border="2px solid black"
                  bg="gray.100"
                >
                  <HStack justify="space-between">
                    <Text fontWeight="bold">{column.name}</Text>
                    <HStack>
                      <IconButton
                        icon={<EditIcon />}
                        size="xs"
                        onClick={() => handleOpenEditColumnModal(column)}
                      />
                      <IconButton
                        icon={<DeleteIcon />}
                        size="xs"
                        colorScheme="red"
                        onClick={() => handleOpenConfirmModal(column.id)}
                      />
                    </HStack>
                  </HStack>
                  <VStack spacing="2">
                    {tasks[column.id]?.map((task, index) => (
                      <Draggable
                        key={task.id}
                        draggableId={String(task.id)}
                        index={index}
                      >
                        {(provided) => (
                          <Box
                            ref={provided.innerRef}
                            {...provided.draggableProps}
                            {...provided.dragHandleProps}
                            p="2"
                            border="1px solid black"
                            w="100%"
                            bg="white"
                            onClick={() => handleOpenEditTaskModal(task)}
                          >
                            <Text fontSize="sm" fontWeight="bold">
                              {task.title}
                            </Text>
                          </Box>
                        )}
                      </Draggable>
                    ))}
                    {provided.placeholder}
                  </VStack>
                  <Button
                    size="sm"
                    mt="4"
                    colorScheme="blackAlpha"
                    onClick={() => handleOpenAddTaskModal(column.id)}
                  >
                    + Add Task
                  </Button>
                </Box>
              )}
            </Droppable>
          ))}
        </HStack>
      </DragDropContext>
    </>
  );
};

export default KanbanBoard;
