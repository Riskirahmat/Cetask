import { Box, Text, HStack, Button, Spacer, VStack } from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";

const Header = ({ onLogout }) => {
  const navigate = useNavigate();

  return (
    <Box
      bg="white"
      color="black"
      px="6"
      py="4"
      border="4px solid black"
      boxShadow="4px 4px 4px 0px black"
    >
      <HStack>
        {/* Logo CeTask */}
        <HStack onClick={() => navigate("/dashboard")} cursor="pointer">
          <VStack align="start" spacing="0">
            <Text fontSize="xl" fontWeight="bold">
              CeTask
            </Text>
            <Text fontSize="sm" color="gray.300">
              Make your task easy
            </Text>
          </VStack>
        </HStack>

        <Spacer />

        {/* Button Logout */}
        <Button colorScheme="red" onClick={onLogout}>
          Logout
        </Button>
      </HStack>
    </Box>
  );
};

export default Header;
