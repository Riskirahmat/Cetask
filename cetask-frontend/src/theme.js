import { extendTheme } from "@chakra-ui/react";

const neobrutalismTheme = extendTheme({
  styles: {
    global: {
      "html, body": {
        background: "#ffffff", 
        color: "#000000", 
        fontFamily: "Arial, sans-serif",
      },
    },
  },
  components: {
    Button: {
      baseStyle: {
        borderRadius: "0px",
        fontWeight: "bold",
        border: "2px solid black",
        boxShadow: "4px 4px 0px black",
        _hover: {
          transform: "translateY(2px)",
          boxShadow: "2px 2px 0px black",
        },
        _active: {
          transform: "translateY(4px)",
          boxShadow: "none",
        },
      },
    },
    Input: {
      baseStyle: {
        borderRadius: "0px",
        border: "2px solid black",
        boxShadow: "4px 4px 0px black",
      },
    },
    Card: {
      baseStyle: {
        border: "2px solid black",
        boxShadow: "6px 6px 0px black",
        borderRadius: "0px",
        padding: "10px",
        _hover: {
          boxShadow: "3px 3px 0px black",
        },
      },
    },
  },
});

export default neobrutalismTheme;
