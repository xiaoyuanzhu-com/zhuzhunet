import React from 'react'
import { Nav, Typography, Table } from '@douyinfe/semi-ui'
const { Text } = Typography
const { Column } = Table

export default function Credits() {
  const thanks = [
    {
      link: 'https://github.com/prometheus-community/pro-bing',
      desc: '',
    },
    {
      link: 'https://semi.design/',
      desc: '',
    },
    {
      link: 'https://vite.dev/',
      desc: '',
    },
    {
      link: 'https://tailwindcss.com/',
      desc: '',
    },
    {
      link: 'https://openlayers.org/',
      desc: '',
    },
  ]

  const handleRow = (record, index) => {
    if (index % 2 === 0) {
      return {
        style: {
          background: 'var(--semi-color-fill-0)',
        },
      };
    } else {
      return {};
    }
  }

  return (
    <div className="p-4">
      <Table dataSource={thanks} onRow={handleRow} pagination={false} showHeader={false}>
        <Column title="Link" dataIndex="link" render={(value) => <a href={value} target="_blank">{value}</a>} />
        <Column title="Description" dataIndex="desc" />
      </Table>
    </div>
  )
}
