import Link from "next/link";
import { usePathname } from "next/navigation";
import { classNames, NavItemType } from "./config";

export const NavItem: React.FC<NavItemType> = (item) => {
  const path = usePathname();
  const active = "underline";
  const hover = "hover:underline";
  const notActive =
    "text-slate-700 group flex gap-x-3 rounded-md p-2 text-sm leading-6 font-semibold";
  return (
    <li key={item.name}>
      <Link
        href={item.href}
        className={classNames(item.href === path ? active : hover, notActive)}
      >
        {/* ICON? */}
        {/* <item.icon className="h-6 w-6 shrink-0" aria-hidden="true" /> */}
        {item.name}
      </Link>
    </li>
  );
};
