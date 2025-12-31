import { DataTable } from '../_components/data-table'
import { columns } from './columns'
import { getLocations } from '@/api/location'

export default async function LocationsPage() {
  const data = await getLocations()
  return (
    <div>
      <h1 className='mb-2'>Locations</h1>
      <DataTable data={data} columns={columns} />
    </div>
  )
}
