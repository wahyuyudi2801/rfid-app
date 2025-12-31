export type RfidCard = {
    id: number;
    rfid_tag: string;
    employee_id: number;
    employee_name: string;
    activation_date: string;
    status: 1 | 2 | 3 | 4 | 5; // 1. Aktif, 2. Tidak Aktif, 3. Hilang, 4. Rusak, 5. Lainnya
}