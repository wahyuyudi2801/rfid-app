"use client"

import { Employee } from "@/types/employee";
import { ColumnDef } from "@tanstack/react-table";
import ActionButtons from "../_components/action-buttons";
import BadgeSuccess from "../_components/badges/success";
import BadgeDanger from "../_components/badges/danger";
import BadgeSecondary from "../_components/badges/secondary";

export const columns: ColumnDef<Employee>[] = [
    {
        accessorKey: "registration_number",
        header: "Registration Number"
    },
    {
        accessorKey: "name",
        header: "Name"
    },
    {
        accessorKey: "work_unit",
        header: "Work Unit"
    },
    {
        accessorKey: "position",
        header: "Position"
    },
    {
        accessorKey: "status",
        header: "Status",
        cell: ({ row }) => {
            const status = row.original.status

            if (status == 1) {
                return <BadgeSuccess text="Aktif" />
            } else if (status == 2) {
                return <BadgeDanger text="Tidak Aktif" />
            } else {
                return <BadgeSecondary text="Unknown" />
            }
        }
    },
    {
        id: "actions",
        header: "Actions",
        enableSorting: false,
        enableHiding: false,
        cell: ({ row }) => {
            const employee = row.original

            return (
                <ActionButtons/>
            )
        }
    }
]