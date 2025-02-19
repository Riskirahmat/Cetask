import { useEffect, useState } from "react";
import {
  getProjects,
  createProject,
  updateProject,
  deleteProject,
} from "../services/projectService";
import {
  Box,
  Button,
  Input,
  VStack,
  Text,
  HStack,
  IconButton,
} from "@chakra-ui/react";
import { EditIcon, DeleteIcon } from "@chakra-ui/icons";
import { useNavigate } from "react-router-dom";
import ConfirmModal from "../components/ConfirmModal";
import EditProjectModal from "../components/EditProjectModal";
import Header from "../components/Header";

const Dashboard = () => {
  const [projects, setProjects] = useState([]);
  const [newProjectName, setNewProjectName] = useState("");
  const [isConfirmOpen, setIsConfirmOpen] = useState(false);
  const [isEditOpen, setIsEditOpen] = useState(false);
  const [projectToEdit, setProjectToEdit] = useState(null);
  const [projectToDelete, setProjectToDelete] = useState(null);
  const navigate = useNavigate();

useEffect(() => {
  const fetchProjects = async () => {
    try {
      const data = await getProjects();
      setProjects(Array.isArray(data) ? data : []);
    } catch (err) {
      console.error("Failed to fetch projects:", err);
      setProjects([]);
    }
  };
  fetchProjects();
}, []);

  const handleCreateProject = async () => {
    if (!newProjectName) return;
    try {
      const newProject = await createProject(newProjectName);
      setProjects([...projects, newProject]);
      setNewProjectName("");
    } catch (err) {
      console.error("Failed to create project:", err);
    }
  };
  const handleOpenEditModal = (project) => {
    setProjectToEdit(project);
    setIsEditOpen(true);
  };

  const handleEditProject = async (newName) => {
    if (!projectToEdit) return;
    try {
      await updateProject(projectToEdit.id, newName);
      setProjects(
        projects?.map((p) =>
          p.id === projectToEdit.id ? { ...p, name: newName } : p
        )
      );
      setIsEditOpen(false);
    } catch (err) {
      console.error("Failed to update project:", err);
    }
  };

  const handleOpenConfirmModal = (projectId) => {
    setProjectToDelete(projectId);
    setIsConfirmOpen(true);
  };

  const handleDeleteProject = async () => {
    if (!projectToDelete) return;
    try {
      await deleteProject(projectToDelete);
      setProjects(projects.filter((p) => p.id !== projectToDelete));
      setIsConfirmOpen(false);
    } catch (err) {
      console.error("Failed to delete project:", err);
    }
  };

    const handleLogout = () => {
      localStorage.removeItem("token");
      navigate("/");
    };

  return (
    <Box maxW="800px" mx="auto" mt="10">
      <Header onLogout={handleLogout} />
      <Text fontSize="2xl" fontWeight="bold" textAlign="center" mb="4">
        Dashboard
      </Text>
      <VStack spacing="4">
        <HStack>
          <Input
            placeholder="New Project Name"
            value={newProjectName}
            onChange={(e) => setNewProjectName(e.target.value)}
          />
          <Button colorScheme="blackAlpha" onClick={handleCreateProject}>
            Create
          </Button>
        </HStack>
      </VStack>
      <HStack spacing="4" overflowX="auto" mt="4">
        {projects?.map((project) => (
          <Box
            key={project.id}
            p="4"
            w="200px"
            border="2px solid black"
            boxShadow="6px 6px 0px black"
            bg="gray.100"
            cursor="pointer"
            onClick={() => navigate(`/projects/${project.id}`)}
          >
            <HStack justify="space-between">
              <Text fontWeight="bold">{project.name}</Text>
              <HStack>
                <IconButton
                  icon={<EditIcon />}
                  size="xs"
                  onClick={(e) => {
                    e.stopPropagation();
                    handleOpenEditModal(project);
                  }}
                />
                <IconButton
                  icon={<DeleteIcon />}
                  size="xs"
                  colorScheme="red"
                  onClick={(e) => {
                    e.stopPropagation();
                    handleOpenConfirmModal(project.id);
                  }}
                />
              </HStack>
            </HStack>
          </Box>
        ))}
      </HStack>
      {/* Modal Edit Project */}
      <EditProjectModal
        isOpen={isEditOpen && projectToEdit !== null}
        onClose={() => setIsEditOpen(false)}
        onConfirm={handleEditProject}
        currentName={projectToEdit ? projectToEdit.name : ""}
      />

      {/* Modal Konfirmasi Hapus Project */}
      <ConfirmModal
        isOpen={isConfirmOpen}
        onClose={() => setIsConfirmOpen(false)}
        onConfirm={handleDeleteProject}
        title="Delete Project"
        message="Are you sure you want to delete this project?"
      />
    </Box>
  );
};

export default Dashboard;
