import React from 'react'
import { Nav, Typography, Table, Image } from '@douyinfe/semi-ui'
import { IconTooltip } from '@douyinfe/semi-icons-lab'
const { Text, Title } = Typography
const { Column } = Table

export default function TestDNS({ manifest, brands }) {
  if (!manifest) {
    return null
  }
  const pageSize = 10
  const data = manifest.dns.servers
  const nameToLogo = {}
  brands.forEach(b => {
    nameToLogo[b.name] = b.logo
  })

  const brandRenderer = (record) => {
    const logo = nameToLogo[record.brand_name]
    const logoDiv = logo ? (
      <Image className="" src={logo} preview={false} />
    ) : (
      <IconTooltip />
    )
    return (
      <div className="flex items-center gap-2">

        <div className="flex items-center justify-center w-6 h-6">
          {logoDiv}
        </div>
        <div>
          <Text>{record.brand_name}</Text>
        </div>
      </div>
    )
  }

  return (
    <div className="p-4">
      <div>
        <Title heading={4}>DNS</Title>
      </div>
      <div className="mt-4">
        <Table dataSource={data} pagination={true} pageSize={pageSize}>
          <Column title="Brand" render={brandRenderer} />
          {/* <Column title="Type" dataIndex="type" /> */}
          <Column title="Address" dataIndex="address" />
          <Column title="Latency" dataIndex="" />
        </Table>
      </div>
    </div>
  )
}
