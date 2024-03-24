"use client";
import React from "react";
import { NavItem } from "@/components/navigation/NavItem";
// import { SettingsBtn } from "@/components/navigation/SettingsBtn";
import { useNavContext } from "@/components/navigation/NavContext";

export const Sidebar: React.FC = () => {
  const nav = useNavContext();
  return (
    <div className="hidden lg:fixed lg:inset-y-0 lg:z-50 lg:flex lg:w-72 lg:flex-col">
      {/* Sidebar component, swap this element with another sidebar if you like */}
      <div className="flex grow flex-col gap-y-5 overflow-y-auto bg-slate-100 border-r-2 border-gray-20 px-6 pb-4">
        <div className="flex h-16 shrink-0 items-center">
          {/* <img
          className="h-8 w-auto"
          src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=500"
          alt="Your Company"
          width=""
        /> */}
        </div>
        <nav className="flex flex-1 flex-col">
          <ul role="list" className="flex flex-1 flex-col gap-y-7">
            <li>
              <ul role="list" className="-mx-2 space-y-1">
                {nav.menu.map((item) => (
                  <NavItem
                    key={item.name}
                    name={item.name}
                    href={item.href}
                    // icon={item.icon}
                  />
                ))}
              </ul>
            </li>
            {/* <li>
              <div className="text-xs font-semibold leading-6 text-gray-400">
              </div>
              <ul role="list" className="-mx-2 mt-2 space-y-1">
                {[].map((item) => (
                  <NavItem key={""} name={""} href={""} />
                ))}
              </ul>
            </li> */}
            {/* <li className="mt-auto">
              <SettingsBtn />
            </li> */}
          </ul>
        </nav>
      </div>
    </div>
  );
};
