import api from "./api";

export const createTask = async (projectId, columnId, title, desc) => {
  try {
    const token = localStorage.getItem("token");
    const response = await api.post(
      `/projects/${projectId}/columns/${columnId}/tasks`,
      { title, desc },
      { headers: { Authorization: `Bearer ${token}` } }
    );
    return response.data;
  } catch (error) {
    throw error.response?.data || { error: "Failed to create task" };
  }
};

export const getTasks = async (projectId, columnId) => {
  try {
    const token = localStorage.getItem("token");
    const response = await api.get(
      `/projects/${projectId}/columns/${columnId}/tasks`,
      { headers: { Authorization: `Bearer ${token}` } }
    );
    return response.data;
  } catch (error) {
    throw error.response?.data || { error: "Failed to fetch tasks" };
  }
};

export const updateTask = async (projectId, columnId, taskId, title, desc) => {
  if (!projectId || !columnId || !taskId) {
    console.error("Invalid parameters for updateTask:", {
      projectId,
      columnId,
      taskId,
    });
    return;
  }
  try {
    const token = localStorage.getItem("token");
    const response = await api.put(
      `/projects/${projectId}/columns/${columnId}/tasks/${taskId}`,
      { title, desc },
      { headers: { Authorization: `Bearer ${token}` } }
    );
    return response.data;
  } catch (error) {
    throw error.response?.data || { error: "Failed to update task" };
  }
};

export const deleteTask = async (projectId, columnId, taskId) => {
  if (!projectId || !columnId || !taskId) {
    console.error("Invalid parameters for deleteTask:", {
      projectId,
      columnId,
      taskId,
    });
    return;
  }
  try {
    const token = localStorage.getItem("token");
    await api.delete(
      `/projects/${projectId}/columns/${columnId}/tasks/${taskId}`,
      { headers: { Authorization: `Bearer ${token}` } }
    );
  } catch (error) {
    throw error.response?.data || { error: "Failed to delete task" };
  }
};

export const moveTask = async (projectId, columnId, taskId, newPosition) => {
  if (!projectId || !columnId || !taskId || isNaN(newPosition)) {
    console.error("Invalid parameters for moveTask:", {
      projectId,
      columnId,
      taskId,
      newPosition,
    });
    return;
  }
  try {
    const token = localStorage.getItem("token");
    const response = await api.put(
      `/projects/${projectId}/columns/${columnId}/tasks/${taskId}/move`,
      { column_id: columnId, position: newPosition },
      { headers: { Authorization: `Bearer ${token}` } }
    );
    return response.data;
  } catch (error) {
    throw error.response?.data || { error: "Failed to move task" };
  }
};
