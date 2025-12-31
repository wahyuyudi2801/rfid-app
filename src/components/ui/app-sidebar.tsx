"use client"

import React, { ForwardRefExoticComponent, RefAttributes } from "react";
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "./sidebar";
import { BookA, IdCard, LayoutGrid, Locate, LucideProps, User, Users } from "lucide-react";
import { paths } from "@/config/paths";
import Link from "next/link";
import Image from "next/image";

type Item = {
  title: string;
  url: string;
  icon: ForwardRefExoticComponent<
    Omit<LucideProps, "ref"> & RefAttributes<SVGSVGElement>
  >;
};

const items: Item[] = [
  {
    title: "Dashboard",
    url: paths.app.dashboard.getHref(),
    icon: LayoutGrid,
  },
  {
    title: "Users",
    url: paths.app.users.getHref(),
    icon: User,
  },
  {
    title: "Employees",
    url: paths.app.employees.getHref(),
    icon: Users,
  },
  
  {
    title: "Attendances",
    url: paths.app.attendances.getHref(),
    icon: BookA,
  },
  {
    title: "RFID Cards",
    url: paths.app.rfidCards.getHref(),
    icon: IdCard,
  },
  {
    title: "Locations",
    url: paths.app.locations.getHref(),
    icon: Locate,
  },
];
import Logo from "@/assets/logo.png";
import { usePathname } from "next/navigation";

export function AppSidebar() {
  const pathname = usePathname();
  
  return (
    <Sidebar collapsible="icon">
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton asChild>
              <Link href="/app" className="py-8">
                <Image src={Logo} alt="base-logo" />
                <span className="font-bold text-lg">RFID App</span>
              </Link>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupContent>
            <SidebarMenu>
              {items.map((item) => (
                <SidebarMenuItem key={item.title}>
                  <SidebarMenuButton isActive={pathname === item.url.toLowerCase()} className="mb-2" asChild>
                    <Link href={item.url} className="opacity-55">
                      <item.icon />
                      <span>{item.title}</span>
                    </Link>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
    </Sidebar>
  );
}
