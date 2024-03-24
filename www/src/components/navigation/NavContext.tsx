import { ReactNode, createContext, useContext } from "react";
import { NavItemType, dftNav } from "./config";

type Navigation = {
  menu: NavItemType[];
};
export const NavContext = createContext<Navigation>({ menu: dftNav });

type props = {
  children: ReactNode;
  navigation: Navigation;
};

export const NavContextProvider: React.FC<props> = (props) => {
  return (
    <NavContext.Provider value={{ menu: props.navigation.menu }}>
      {props.children}
    </NavContext.Provider>
  );
};

export const useNavContext = () => {
  const ctx = useContext(NavContext);
  if (!ctx) {
    throw new Error("useNavContext must be used within a NavContext provider");
  }
  return ctx;
};
