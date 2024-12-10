import React from 'react'
import { Nav, Typography, Table, Image, Spin } from '@douyinfe/semi-ui'
import { IconTooltip } from '@douyinfe/semi-icons-lab'
const { Text, Title } = Typography
import api from '../../lib/api'
import { useState, useEffect } from 'react'
import Brand from './components/brand'
const { Column } = Table

export default function DiagnoseDNS() {
  const [dnsList, setDNSList] = useState(null)
  const [brandList, setBrandList] = useState(null)
  const fetchDNSList = async () => {
    const dnsList = await api.getDNSList()
    setDNSList(dnsList)
  }
  const fetchBrandList = async () => {
    const brandList = await api.getBrandList()
    setBrandList(brandList)
  }
  useEffect(() => {
    fetchDNSList()
    fetchBrandList()
  }, [])
  if (!dnsList || !brandList) {
    return <div className="flex items-center justify-center h-full"><Spin /></div>
  }

  const pageSize = 10
  const data = dnsList.list

  const brandRenderer = (record) => {
    return <Brand brandList={brandList} brandID={record.brands.id} />
  }

  return (
    <div className="p-4">
      <div>
        <Title heading={4}>DNS</Title>
      </div>
      <div className="mt-4">
        <Table dataSource={data} pagination={true} pageSize={pageSize}>
          <Column title="Brand" render={brandRenderer} />
          <Column title="Address" dataIndex="address" />
          <Column title="Latency" dataIndex="" />
        </Table>
      </div>
    </div>
  )
}
