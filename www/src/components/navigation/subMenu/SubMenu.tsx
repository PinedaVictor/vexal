"use client";
import Link from "next/link";
import { NavItemType, classNames } from "../config";
import { usePathname } from "next/navigation";

type props = {
  menu: NavItemType[];
};

export const SubMenu: React.FC<props> = (props) => {
  const path = usePathname();
  const tabs = props.menu;
  const active = "border-indigo-500 text-indigo-600";
  const hover =
    "border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700";
  const notActive =
    "group inline-flex items-center border-b-2 py-4 px-1 text-sm font-medium";
  return (
    <div className="fixed bottom-4 bg-slate-50 w-full">
      <div className="sm:block justify-items-end">
        <div className="border-b border-gray-200">
          <nav className="-mb-px flex space-x-8" aria-label="Tabs">
            {tabs.map((tab) => (
              <Link
                key={tab.name}
                href={tab.href}
                className={classNames(
                  tab.href === path ? active : hover,
                  notActive
                )}
              >
                {/* <tab.icon
                  className={classNames(
                    tab.href === path
                      ? "text-indigo-500"
                      : "text-gray-400 group-hover:text-gray-500",
                    "-ml-0.5 mr-2 h-5 w-5"
                  )}
                  aria-hidden="true"
                />
                <span>{tab.name}</span> */}
              </Link>
            ))}
          </nav>
        </div>
      </div>
    </div>
  );
};
