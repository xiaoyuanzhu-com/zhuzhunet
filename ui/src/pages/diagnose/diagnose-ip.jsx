import React from 'react'
import { Nav, Typography, Table, Image, Spin, Select, InputNumber, Button, Descriptions, Input } from '@douyinfe/semi-ui'
import { IconTooltip } from '@douyinfe/semi-icons-lab'
const { Text, Title } = Typography
import api from '../../lib/api'
import { useState, useEffect } from 'react'
import Brand from './components/brand'
import { VChart } from "@visactor/react-vchart"
import GeoMap from '@/components/geo-map'
import { useSearchParams } from "react-router"
const { Column } = Table

export default function DiagnoseIP() {
  const [searchParams, setSearchParams] = useSearchParams()
  const ip = searchParams.get('ip')
  const setIP = (value) => {
    setSearchParams({ ip: value })
  }
  const [ipInfo, setIPInfo] = useState(null)
  const fetchIPInfo = async () => {
    if (!ip) {
      return
    }
    const data = await api.getIPInfo(ip)
    if (data && data.length > 0) {
      setIPInfo(data[0])
    }
  }
  useEffect(() => {
    fetchIPInfo()
  }, [])

  let descriptionDiv = null
  if (ipInfo) {
    descriptionDiv = <Descriptions>
      <Descriptions.Item itemKey="IP">{ipInfo.ip}</Descriptions.Item>
      <Descriptions.Item itemKey="Country">{ipInfo.country}</Descriptions.Item>
      <Descriptions.Item itemKey="City">{ipInfo.city}</Descriptions.Item>
      <Descriptions.Item itemKey="ASN">{ipInfo.asn}</Descriptions.Item>
      <Descriptions.Item itemKey="AS">{ipInfo.as}</Descriptions.Item>
      <Descriptions.Item itemKey="Geo">
        <div className="w-[640px] h-[360px]">
          <GeoMap longitude={ipInfo.longitude} latitude={ipInfo.latitude} />
        </div>
      </Descriptions.Item>
    </Descriptions>
  }

  return (
    <div className="p-4">
      <div>
        <Title heading={4}>IP Info</Title>
      </div>
      <div className="mt-4">
        <div className="flex items-center justify-center">
          <div className="flex">
            <div className="w-[480px]">
              <Input value={ip} onChange={setIP}></Input>
            </div>
            <div className="w-[64px] ml-2">
              <Button type="primary" onClick={fetchIPInfo}>Lookup</Button>
            </div>
          </div>
        </div>
        <div>
          {descriptionDiv}
        </div>
      </div>
    </div>
  )
}
