import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";

export default function Home() {
  return (
    <div className="p-4">
      <h1>Welcome to RFID App</h1>
      <br />
      <div className="w-56">
        <Input className="mb-2" placeholder="Masukkan input..." />
        
        <Button>Klik</Button>
      </div>
    </div>
  );
}
