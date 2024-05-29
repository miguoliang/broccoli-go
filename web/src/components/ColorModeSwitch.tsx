import { Button } from "@headlessui/react";
import useColorModeStore from "../stores/ColorModeStore";
import { useCallback } from "react";
import { FaMoon, FaSun } from "react-icons/fa";

const ColorModeSwitch = () => {
  const { colorMode, setColorMode } = useColorModeStore();
  const handleModeChange = useCallback(() => {
    setColorMode(colorMode === "light" ? "dark" : "light");
  }, [colorMode, setColorMode]);
  return (
    <Button onClick={handleModeChange}>
      {colorMode === "light" ? <FaMoon /> : <FaSun />}
    </Button>
  );
};

export default ColorModeSwitch;
