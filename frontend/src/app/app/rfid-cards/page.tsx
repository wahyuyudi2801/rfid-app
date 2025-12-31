import { getRfidCards } from '@/api/rfid_card'
import { DataTable } from '../_components/data-table'
import { columns } from './columns'

export default async function RFIDCardsPage() {
  const data = await getRfidCards()
  return (
    <div>
      <h1 className='mb-2'>RFID Cards</h1>
      <DataTable data={data} columns={columns} />
    </div>
  )
}
