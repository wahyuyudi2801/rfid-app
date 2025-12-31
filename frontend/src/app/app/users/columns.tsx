"use client"

import { User } from "@/types/user";
import { ColumnDef } from "@tanstack/react-table";
import Link from "next/link";
import ActionButtons from "../_components/action-buttons";

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
        id: "actions",
        header: "Actions",
        enableSorting: false,
        enableHiding: false,
        cell: ({ row }) => {
            const user = row.original

            return (
                <ActionButtons/>
            )
        }
    }
]