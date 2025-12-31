"use client"

import { Button } from "@/components/ui/button";
import { User } from "@/types/user";
import { ColumnDef } from "@tanstack/react-table";
import { EditIcon, Trash } from "lucide-react";
import Link from "next/link";

export const columns: ColumnDef<User>[] = [
    {
        accessorKey: "username",
        header: "Username",
    },
    {
        accessorKey: "role",
        header: "Role",
    },
    {
        accessorKey: "registration_number",
        header: "Registration Number",
        cell: ({ row }) => {
            const user = row.original

            return (
                <Link href={`app/employees/${user.employee_id}`} className="text-blue-500 underline hover:text-blue-500/90">
                    {user.registration_number}
                </Link>
            )
        }
    },
    {
        header: "Actions",
        enableSorting: false,
        enableHiding: false,
        id: "actions",
        cell: ({ row }) => {
            const user = row.original

            return (
                <div className="flex items-center gap-2">
                    <Button variant={'default'} className="bg-yellow-500 hover:bg-yellow-500/90">
                        <EditIcon/>
                    </Button>
                    <Button variant={'destructive'}>
                        <Trash/>
                    </Button>
                </div>
            )
        }
    }
]