import { Badge } from '@/components/ui/badge'

export default function BadgeInfo({ text }: { text: string }) {
  return (
    <Badge className="h-5 min-w-5 rounded-full px-1 font-mono bg-blue-500">{text}</Badge>
  )
}
