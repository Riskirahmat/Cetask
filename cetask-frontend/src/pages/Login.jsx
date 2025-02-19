import { useState, useContext } from "react";
import { AuthContext } from "../context/AuthContext";
import { login } from "../services/authService";
import { useNavigate } from "react-router-dom";
import { Button, Input, Box, Text, VStack, Link } from "@chakra-ui/react";

const Login = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const { login: setAuth } = useContext(AuthContext);
  const navigate = useNavigate();

const handleLogin = async () => {
  try {
    const data = await login(email, password);
    setAuth(data.token, navigate);
  } catch (err) {
    setError(err.error || "Login failed");
    console.error("Login Error:", err);
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
        LOGIN
      </Text>
      {error && <Text color="red.500">{error}</Text>}
      <VStack spacing="4">
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
        <Button colorScheme="blackAlpha" onClick={handleLogin}>
          Login
        </Button>
        <Text fontSize="sm">
          Don't have an account?{" "}
          <Link color="blue.500" onClick={() => navigate("/register")}>
            Register
          </Link>
        </Text>
      </VStack>
    </Box>
  );
};

export default Login;
