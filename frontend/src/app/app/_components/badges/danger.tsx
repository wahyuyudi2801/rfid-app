import { Badge } from '@/components/ui/badge'

export default function BadgeDanger({ text }: { text : string }) {
  return (
    <Badge variant="destructive">{text}</Badge>
  )
}
