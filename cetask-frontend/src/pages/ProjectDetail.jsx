import { Box } from "@chakra-ui/react";
import KanbanBoard from "../components/KanbanBoard";

const ProjectDetail = () => {
  return (
    <Box maxW="800px" mx="auto" mt="10">
      <KanbanBoard />
    </Box>
  );
};

export default ProjectDetail;
