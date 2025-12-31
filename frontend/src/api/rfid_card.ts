import { RfidCard } from "@/types/rfid_cards";

export async function getRfidCards(): Promise<RfidCard[]> {
    return Promise.resolve([
        {
            id: 1,
            rfid_tag: '123456',
            employee_id: 1,
            employee_name: 'John Doe',
            activation_date: '2023-01-01',
            status: 1
        }
    ])
}