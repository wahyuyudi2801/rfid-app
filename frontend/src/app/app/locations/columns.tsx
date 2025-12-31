"use client"

import { Location } from "@/types/location";
import { ColumnDef } from "@tanstack/react-table";
import ActionButtons from "../_components/action-buttons";

export const columns: ColumnDef<Location>[] = [
    {
        accessorKey: "location",
        header: "Location"
    },
    {
        accessorKey: "reader_code",
        header: "Reader Code"
    },
    {
        id: "actions",
        header: "Actions",
        enableSorting: false,
        enableHiding: false,
        cell: ({ row }) => {
            const location = row.original

            return <ActionButtons/>
        }
    }
]