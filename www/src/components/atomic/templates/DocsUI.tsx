"use client";

import { NavContextProvider } from "@/components/navigation/NavContext";
import { navigationMenu } from "@/components/navigation/config";
import { Navigation } from "@/components/navigation/Navigation";
import { ReactNode } from "react";

type Props = {
  children: ReactNode;
};

export const DocsUI: React.FC<Props> = (props) => {
  return (
    <NavContextProvider navigation={{ menu: navigationMenu }}>
      <Navigation>{props.children}</Navigation>
    </NavContextProvider>
  );
};
