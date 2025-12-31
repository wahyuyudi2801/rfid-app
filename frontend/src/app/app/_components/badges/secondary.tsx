import { Badge } from '@/components/ui/badge'

export default function BadgeSecondary({ text }: { text: string }) {
  return (
    <Badge variant={"secondary"}>{text}</Badge>
  )
}
