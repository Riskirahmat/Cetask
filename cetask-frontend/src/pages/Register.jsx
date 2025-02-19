import { useState } from "react";
import { register } from "../services/authService";
import { useNavigate } from "react-router-dom";
import { Button, Input, Box, Text, VStack, Link } from "@chakra-ui/react";

const Register = () => {
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const handleRegister = async () => {
    try {
      await register(username, email, password);
      navigate("/");
    } catch (err) {
      setError(err.error || "Registration failed");
    }
  };

  return (
    <Box
      maxW="400px"
      mx="auto"
      mt="10"
      p="6"
      border="2px solid black"
      boxShadow="8px 8px 0px black"
    >
      <Text fontSize="2xl" fontWeight="bold" mb="4" textAlign="center">
        REGISTER
      </Text>
      {error && <Text color="red.500">{error}</Text>}
      <VStack spacing="4">
        <Input
          placeholder="Username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
        <Input
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <Input
          placeholder="Password"
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <Button colorScheme="blackAlpha" onClick={handleRegister}>
          Register
        </Button>
        <Text fontSize="sm">
          Already have an account?{" "}
          <Link color="blue.500" onClick={() => navigate("/")}>
            Login
          </Link>
        </Text>
      </VStack>
    </Box>
  );
};

export default Register;