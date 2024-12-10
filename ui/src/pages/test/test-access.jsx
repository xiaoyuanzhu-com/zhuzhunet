import React from 'react'
import { Nav, Typography, Table, Image } from '@douyinfe/semi-ui'
import { IconTooltip } from '@douyinfe/semi-icons-lab'
const { Text, Title } = Typography
const { Column } = Table

export default function TestAccess({ manifest, brands }) {
  if (!manifest) {
    return null
  }

  return (
    <div className="p-4">
      <div>
        <Title heading={4}>Access</Title>
      </div>
    </div>
  )
}
