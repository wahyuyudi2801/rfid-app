import { DataTable } from "../_components/data-table";
import { columns } from "./columns";
import { getUsers } from "@/api/user";

export default async function UserPage() {
  const data = getUsers();
  return (
    <div>
      <h1 className="mb-2">Users Page</h1>
      <DataTable columns={columns} data={data} />
    </div>
  );
}
