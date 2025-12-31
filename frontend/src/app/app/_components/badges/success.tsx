import { Badge } from '@/components/ui/badge'

export default function BadgeSuccess({ text }: { text: string }) {
    console.log(text)
  return (
    <Badge className="h-5 min-w-5 rounded-full px-1 font-mono bg-green-500">{text}</Badge>
  )
}
