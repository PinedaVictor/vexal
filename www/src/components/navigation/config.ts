// import {
//   CalendarIcon,
//   UsersIcon,
//   FolderIcon,
//   HomeIcon,
//   RocketLaunchIcon,
//   NewspaperIcon,
//   CreditCardIcon,
//   UserIcon,
//   ShareIcon,
//   PencilIcon,
// } from "@heroicons/react/24/outline";

export function classNames(...classes: string[]) {
  return classes.filter(Boolean).join(" ");
}

export type NavItemType = {
  name: string;
  href: string;
};

// export const dftNav = [{ name: "NA", href: "/", icon: HomeIcon }];
export const dftNav = [{ name: "NA", href: "/" }];

export const navigationMenu = [
  { name: "Getting started", href: "/" },
  { name: "Deps", href: "/deps" },
];
