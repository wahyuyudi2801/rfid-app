"use client";

import { Button } from "@/components/ui/button";
import { EditIcon, Eye, Plus, Trash } from "lucide-react";

type ActionButtons = {
  create?: boolean;
  detail?: boolean;
  edit?: boolean;
  destroy?: boolean;
  create_url?: string;
  detail_url?: string;
  edit_url?: string;
  destroy_url?: string;
};

export default function ActionButtons({
  detail = true,
  edit = true,
  destroy = true,
  detail_url,
  edit_url,
  destroy_url,
}: ActionButtons) {
  return (
    <div className="flex items-center gap-2">

      {detail && (
        <Button
          variant={"default"}
          className="bg-blue-500 hover:bg-blue-500/90"
        >
          <Eye />
        </Button>
      )}

      {edit && (
        <Button
          variant={"default"}
          className="bg-yellow-500 hover:bg-yellow-500/90"
        >
          <EditIcon />
        </Button>
      )}

      {destroy && (
        <Button variant={"destructive"}>
          <Trash />
        </Button>
      )}
    </div>
  );
}
