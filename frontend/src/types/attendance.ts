export type Attendance = {
    id: number;
    rfid_id: number;
    rfid_tag: string;
    location_id: number;
    location: string;
    tap_type: 1 | 2; // 1 = Masuk, 2 = Keluar
    validation_status: 1 | 2 | 3 | 4 | 5; // 1: Hadir, 2: Terlambat, 3: Izin, 4: Sakit, 5: Bolos/Tidak Hadir
    created_at: string;
}