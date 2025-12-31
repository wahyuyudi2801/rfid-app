export type Employee = {
    id: number;
    name: string;
    phone: string;
    email: string;
    registration_number: string;
    position: string;
    work_unit: string;
    status: 1 | 2; // 1. Aktif, 2. Tidak Aktif
}