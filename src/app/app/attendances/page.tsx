import React from 'react'
import { DataTable } from '../_components/data-table'
import { columns } from './columns'
import getAttendances from '@/api/attendance'

export default async function AttendancesPage() {
  const data = await getAttendances()
  return (
    <div>
      <h1 className='mb-2'>Attendances</h1>
      <DataTable data={data} columns={columns} />
    </div>
  )
}
