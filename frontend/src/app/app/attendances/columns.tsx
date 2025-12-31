"use client";

import { Attendance } from "@/types/attendance";
import { ColumnDef } from "@tanstack/react-table";
import Link from "next/link";
import BadgeSuccess from "../_components/badges/success";
import BadgeDanger from "../_components/badges/danger";
import BadgeSecondary from "../_components/badges/secondary";
import BadgeWarning from "../_components/badges/warning";
import BadgeInfo from "../_components/badges/info";
import ActionButtons from "../_components/action-buttons";

export const columns: ColumnDef<Attendance>[] = [
  {
    accessorKey: "tap_type",
    header: "Tap Type",
    cell: ({ row }) => {
      const tapType = row.original.tap_type;

      if (tapType == 1) {
        return <BadgeSuccess text="Masuk" />;
      } else if (tapType == 2) {
        return <BadgeDanger text="Keluar" />;
      } else {
        return <BadgeSecondary text="Unknown" />;
      }
    },
  },
  {
    accessorKey: "validation_status",
    header: "Validation Status",
    cell: ({ row }) => {
      const validation_status = row.original.validation_status;

      if (validation_status == 1) {
        return <BadgeSuccess text="Hadir" />;
      } else if (validation_status == 2) {
        return <BadgeWarning text="Terlambat" />;
      } else if (validation_status == 3) {
        return <BadgeInfo text="Izin" />;
      } else if (validation_status == 4) {
        return <BadgeInfo text="Sakit" />;
      } else if (validation_status == 5) {
        return <BadgeDanger text="Bolos/Tidak Hadir" />;
      } else {
        return <BadgeSecondary text="Unknown" />;
      }
    },
  },
  {
    accessorKey: "location",
    header: "Location",
    cell: ({ row }) => {
      const attendance = row.original;

      return (
        <Link href={`app/locations/${attendance.location_id}`}>
          {attendance.location}
        </Link>
      );
    },
  },
  {
    accessorKey: "rfid_tag",
    header: "RFID",
    cell: ({ row }) => {
      const attendance = row.original;

      return (
        <Link href={`app/rfid-cards/${attendance.rfid_id}`}>
          {attendance.rfid_tag}
        </Link>
      );
    },
  },
  {
    id: "actions",
    enableSorting: false,
    enableHiding: false,
    header: "Actions",
    cell: ({ row }) => {
        const attendance = row.original;

        return (
            <ActionButtons edit={false} />
        )
    }
  }
];
