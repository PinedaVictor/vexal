import type { ReactNode } from "react";
import { DocsUI } from "@/components/atomic/templates/DocsUI";

export default function Home(props: { children: ReactNode }) {
  return (
    <main className=" bg-slate-100 h-screen">
      <DocsUI>
        <div>This is my docs UI</div>
      </DocsUI>
    </main>
  );
}
