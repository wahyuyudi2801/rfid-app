import { User } from "@/types/user";

export function getUsers(): User[] {
    return [
        {
            id: 1,
            username: "admin",
            role_id: 1,
            employee_id: 1,
            role: "Admin",
            registration_number: "1234"
        }
    ]
}