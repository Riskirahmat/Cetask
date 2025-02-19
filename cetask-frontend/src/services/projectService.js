import api from "./api";

export const getProjects = async () => {
  try {
    const token = localStorage.getItem("token");
    const response = await api.get("/projects/", {
      headers: { Authorization: `Bearer ${token}` },
    });
    return response.data;
  } catch (error) {
    throw error.response?.data || { error: "Failed to fetch projects" };
  }
};

export const getProjectById = async (projectId) => {
  try {
    const token = localStorage.getItem("token");
    const response = await api.get(`/projects/${projectId}`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    return response.data;
  } catch (error) {
    throw error.response?.data || { error: "Failed to fetch project" };
  }
};

export const createProject = async (name) => {
  try {
    const token = localStorage.getItem("token");
    const response = await api.post(
      "/projects/",
      { name },
      { headers: { Authorization: `Bearer ${token}` } }
    );
    return response.data;
  } catch (error) {
    throw error.response?.data || { error: "Failed to create project" };
  }
};


export const updateProject = async (projectId, newName) => {
  try {
    const token = localStorage.getItem("token");
    const response = await api.put(
      `/projects/${projectId}`,
      { name: newName },
      { headers: { Authorization: `Bearer ${token}` } }
    );
    return response.data;
  } catch (error) {
    throw error.response?.data || { error: "Failed to update project" };
  }
};

export const deleteProject = async (projectId) => {
  try {
    const token = localStorage.getItem("token");
    await api.delete(`/projects/${projectId}`, {
      headers: { Authorization: `Bearer ${token}` },
    });
  } catch (error) {
    throw error.response?.data || { error: "Failed to delete project" };
  }
};



export const getColumns = async (projectId) => {
  try {
    const token = localStorage.getItem("token");
    const response = await api.get(`/projects/${projectId}/columns`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    return response.data;
  } catch (error) {
    throw error.response?.data || { error: "Failed to fetch columns" };
  }
};

export const createColumn = async (projectId, name) => {
  try {
    const token = localStorage.getItem("token");
    const response = await api.post(
      `/projects/${projectId}/columns`,
      { name },
      { headers: { Authorization: `Bearer ${token}` } }
    );
    return response.data;
  } catch (error) {
    throw error.response?.data || { error: "Failed to create column" };
  }
};

export const updateColumn = async (projectId, columnId, newName) => {
  try {
    const token = localStorage.getItem("token");
    const response = await api.put(
      `/projects/${projectId}/columns/${columnId}`,
      { name: newName },
      { headers: { Authorization: `Bearer ${token}` } }
    );
    return response.data;
  } catch (error) {
    throw error.response?.data || { error: "Failed to update column" };
  }
};

export const deleteColumn = async (projectId, columnId) => {
  try {
    const token = localStorage.getItem("token");
    await api.delete(`/projects/${projectId}/columns/${columnId}`, {
      headers: { Authorization: `Bearer ${token}` },
    });
  } catch (error) {
    throw error.response?.data || { error: "Failed to delete column" };
  }
};


