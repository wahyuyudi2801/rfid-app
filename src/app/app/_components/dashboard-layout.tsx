import { AppSidebar } from "@/components/ui/app-sidebar";
import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar";
import { ReactNode } from "react";

export default function DashboardLayout({ children }: { children: ReactNode }) {
  return (
    <SidebarProvider>
      <AppSidebar/>
      <main className="p-4 w-full flex flex-col gap-4">
        <section className="bg-white w-full p-4 rounded-lg shadow flex items-center gap-2 ">
          <SidebarTrigger />
          <span>Dashboard</span>
        </section>

        <section className="bg-white w-full p-4 rounded-lg shadow">
          {children}
        </section>
      </main>
    </SidebarProvider>
  );
}
