import { create } from "zustand";

type ColorModeStoreState = {
  colorMode: "light" | "dark";
  setColorMode: (mode: "light" | "dark") => void;
};

const useColorModeStore = create<ColorModeStoreState>((set) => ({
  colorMode: "light",
  setColorMode: (mode) => set({ colorMode: mode }),
}));

export default useColorModeStore;
