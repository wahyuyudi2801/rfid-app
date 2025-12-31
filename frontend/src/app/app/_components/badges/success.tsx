import { Badge } from "@/components/ui/badge";

export default function BadgeSuccess({ text }: { text: string }) {
  return (
    <Badge className="h-5 min-w-5 rounded-full px-1 font-mono bg-green-500">
      {text}
    </Badge>
  );
}
