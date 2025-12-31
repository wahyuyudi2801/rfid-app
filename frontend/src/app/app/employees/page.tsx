import { DataTable } from '../_components/data-table'
import getEmployees from '@/api/employee'
import { columns } from './columns'

export default async function EmployeesPage() {
  const data = await getEmployees()

  return (
    <div>
      <h1 className='mb-2'>Employees</h1>
      <DataTable data={data} columns={columns} />
    </div>
  )
}
