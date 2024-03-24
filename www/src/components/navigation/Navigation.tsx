"use client";
import { useState } from "react";
import { Menu } from "@headlessui/react";
import {
  Bars3Icon,
  BellIcon,
  MagnifyingGlassIcon,
} from "@heroicons/react/24/outline";
import { UserButton } from "@clerk/nextjs";
import { MobileDrawer, Sidebar } from "../atomic/organisms";

type props = {
  children: React.ReactNode;
};

export const Navigation: React.FC<props> = (props) => {
  const [sidebarOpen, setSidebarOpen] = useState(false);
  return (
    <div>
      <MobileDrawer
        toggle={sidebarOpen}
        setToggle={() => setSidebarOpen(!sidebarOpen)}
      />
      {/* Static sidebar for desktop */}
      <Sidebar />
      <div className="lg:pl-72">
        <div className="sticky top-0 z-40 flex h-16 shrink-0 items-center gap-x-4 border-b border-gray-200 bg-white px-4 shadow-sm sm:gap-x-6 sm:px-6 lg:px-8">
          <button
            type="button"
            className="-m-2.5 p-2.5 text-gray-700 lg:hidden"
            onClick={() => setSidebarOpen(true)}
          >
            <span className="sr-only">Open sidebar</span>
            <Bars3Icon className="h-6 w-6" aria-hidden="true" />
          </button>

          {/* Separator */}
          <div
            className="h-6 w-px bg-gray-900/10 lg:hidden"
            aria-hidden="true"
          />

          <div className="flex flex-1 gap-x-4 self-stretch lg:gap-x-6">
            {/* Search Form */}
            <form className="relative flex flex-1" action="#" method="GET">
              {/* TODO: Search Form */}
              {/* <label htmlFor="search-field" className="sr-only">
                Search
              </label>
              <MagnifyingGlassIcon
                className="pointer-events-none absolute inset-y-0 left-0 h-full w-5 text-gray-400"
                aria-hidden="true"
              />
              <input
                id="search-field"
                className="block h-full w-full border-0 py-0 pl-8 pr-0 text-gray-900 placeholder:text-gray-400 focus:ring-0 sm:text-sm"
                placeholder="Search..."
                type="search"
                name="search"
              /> */}
            </form>
            <div className="flex items-center gap-x-4 lg:gap-x-6">
              {/* TODO: Notifications */}
              {/* <button
                type="button"
                className="-m-2.5 p-2.5 text-gray-400 hover:text-gray-500"
              >
                <span className="sr-only">View notifications</span>
                <BellIcon className="h-6 w-6" aria-hidden="true" />
              </button> */}

              {/* Separator */}
              <div
                className="hidden lg:block lg:h-6 lg:w-px lg:bg-gray-900/10"
                aria-hidden="true"
              />

              {/* Profile dropdown */}
              <Menu as="div" className="relative">
                <Menu.Button className="-m-1.5 flex items-center p-1.5">
                  <span className="lg:flex lg:items-center">
                    <UserButton afterSignOutUrl="/sign-in" />
                  </span>
                </Menu.Button>
              </Menu>
            </div>
          </div>
        </div>

        <main className="py-10 ">
          <div className="px-4 sm:px-6 lg:px-8 ">{props.children}</div>
        </main>
      </div>
    </div>
  );
};
