"use client";

import { RfidCard } from "@/types/rfid_cards";
import { ColumnDef } from "@tanstack/react-table";
import ActionButtons from "../_components/action-buttons";
import Link from "next/link";
import BadgeSecondary from "../_components/badges/secondary";
import BadgeWarning from "../_components/badges/warning";
import BadgeDanger from "../_components/badges/danger";
import BadgeSuccess from "../_components/badges/success";

export const columns: ColumnDef<RfidCard>[] = [
  {
    accessorKey: "rfid_tag",
    header: "RFID",
  },
  {
    accessorKey: "activation_date",
    header: "Activation Date",
  },
  {
    accessorKey: "status",
    header: "Status",
    cell: ({ row }) => {
      const status = row.original.status;
        
      if (status === 1) {
        return <BadgeSuccess text="Aktif" />;
      } else if (status === 2) {
        return <BadgeDanger text="Tidak Aktif" />;
      } else if (status === 3) {
        return <BadgeWarning text="Hilang" />;
      } else if (status === 4) {
        return <BadgeWarning text="Rusak" />;
      } else if (status === 5) {
        return <BadgeSecondary text="Lainnya" />;
      } else {
        return <BadgeSecondary text="Unknown" />;
      }
    },
  },
  {
    accessorKey: "employee_name",
    header: "Employee",
    cell: ({ row }) => (
      <Link href={`app/employees/${row.original.employee_id}`}>
        {row.original.employee_name}
      </Link>
    ),
  },
  {
    id: "actions",
    header: "Actions",
    enableSorting: false,
    enableHiding: false,
    cell: ({ row }) => <ActionButtons />,
  },
];
