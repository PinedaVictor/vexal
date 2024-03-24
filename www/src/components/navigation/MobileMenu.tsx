import { NavItem } from "./NavItem";
// import { NavItem } from "@/components/navigation/NavItem";
import { SettingsBtn } from "./SettingsBtn";
import { useNavContext } from "./NavContext";

type props = {};
export const MobileMenu: React.FC<props> = (props) => {
  const nav = useNavContext();
  return (
    <div className="flex grow flex-col gap-y-5 overflow-y-auto bg-slate-100 px-6 pb-4 ring-1 ring-white/10">
      <div className="flex h-16 shrink-0 items-center">item</div>
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
              Menu
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
  );
};
