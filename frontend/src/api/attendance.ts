import { Attendance } from '@/types/attendance'

export default function getAttendances(): Promise<Attendance[]> {
  return Promise.resolve([
    {
      id: 1,
      rfid_id: 1,
      rfid_tag: '123456',
      location_id: 1,
      location: 'Location 1',
      tap_type: 1,
      validation_status: 1,
      created_at: '2023-01-01'
    }
  ])
}
