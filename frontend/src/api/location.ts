import { Location } from "@/types/location";

export async function getLocations(): Promise<Location[]> {
    return Promise.resolve([
        {
            id: 1,
            location: "abc",
            reader_code: "RC123"
        }
    ])
}