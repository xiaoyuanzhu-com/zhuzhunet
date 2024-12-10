import { IconTooltip } from "@douyinfe/semi-icons-lab"
import { Image, Typography } from "@douyinfe/semi-ui"
const { Text } = Typography

export default function Brand({ brandList, brandID }) {
  if (!brandList) {
    return null
  }
  const brand = brandList.list.find(b => b.id === brandID)
  if (!brand) {
    return null
  }
  const logoURL = brand.logo && brand.logo.length > 0 ? brand.logo[0].signed_url : null
  const logoDiv = logoURL ? (
    <Image className="" src={logoURL} preview={false} />
  ) : (
    <IconTooltip />
  )
  return (
    <div className="flex items-center gap-2">
      <div className="flex items-center justify-center w-6 h-6">
        {logoDiv}
      </div>
      <div>
        <Text>{brand.name}</Text>
      </div>
    </div>
  )
}
