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
} from "@chakra-ui/react";
import { useState } from "react";

const EditProjectModal = ({ isOpen, onClose, onConfirm, currentName }) => {
  const [newName, setNewName] = useState(currentName);

  const handleSave = () => {
    if (!newName.trim() || newName === currentName) return;
    onConfirm(newName);
    onClose();
  };

  return (
    <Modal isOpen={isOpen} onClose={onClose} isCentered>
      <ModalOverlay />
      <ModalContent
        p="4"
        border="2px solid black"
        boxShadow="6px 6px 0px black"
      >
        <ModalHeader>Edit Project</ModalHeader>
        <ModalBody>
          <VStack spacing="4">
            <Input
              value={newName}
              onChange={(e) => setNewName(e.target.value)}
              placeholder="Enter new project name"
            />
          </VStack>
        </ModalBody>
        <ModalFooter>
          <Button colorScheme="gray" onClick={onClose}>
            Cancel
          </Button>
          <Button colorScheme="blue" ml="3" onClick={handleSave}>
            Save
          </Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};

export default EditProjectModal;
