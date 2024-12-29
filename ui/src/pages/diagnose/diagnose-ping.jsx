import React from 'react'
import { Nav, Typography, Table, Image, Spin, Select, InputNumber, Button, Descriptions } from '@douyinfe/semi-ui'
import { IconTooltip } from '@douyinfe/semi-icons-lab'
const { Text, Title } = Typography
import api from '../../lib/api'
import { useState, useEffect } from 'react'
import Brand from './components/brand'
import { VChart } from "@visactor/react-vchart"
const { Column } = Table

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

function urlToDomain(url) {
  const urlObj = new URL(url)
  return urlObj.hostname
}

export default function DiagnosePing() {
  const [websiteList, setWebsiteList] = useState(null)
  const [domainList, setDomainList] = useState(null)
  const [selectedAddress, setSelectedAddress] = useState(null)
  const [count, setCount] = useState(10)
  const [report, setReport] = useState(null)
  const fetchWebsiteList = async () => {
    const websiteList = await api.getWebsiteList()
    setWebsiteList(websiteList)
    const domainList = websiteList.list.map(item => urlToDomain(item.address))
    setDomainList(domainList)
  }
  useEffect(() => {
    fetchWebsiteList()
  }, [])
  if (!websiteList) {
    return <div className="flex items-center justify-center h-full"><Spin /></div>
  }
  if (!selectedAddress) {
    setSelectedAddress(domainList[0])
  }

  const optionDivs = domainList.map(item => (
    <Select.Option value={item}>{item}</Select.Option>
  ))

  const onPingClicked = async () => {
    while (true) {
      const res = await api.ping(selectedAddress, count)
      setReport(res)
      if (res.meta.status !== 'in-progress') {
        break
      }
      await sleep(1000)
    }
  }

  let chartDiv = null
  if (report && report.data && report.data.rtt_ns) {
    const chartSpec = {
      type: 'line',
      data: [
        {
          id: 'ping',
          values: report.data.rtt_ns.map((rtt_ns, idx) => ({
            x: new Date(report.data.rtt_timestamps[idx]).getTime(),
            y: rtt_ns,
          })),
        }
      ],
      xField: 'x',
      yField: 'y',
      axes: [
        {
          orient: 'left',
          type: 'linear',
          label: {
            formatMethod(val) {
              return `${(val / 1000000).toFixed(2)}ms`;
            }
          }
        },
        {
          orient: 'bottom',
          type: 'time',
          layers: [
            {
              timeFormat: '%Y-%m-%d %H:%M:%S',
            }
          ]
        }
      ],
      tooltip: {
        mark: {
          title: {
            visible: false
          },
          content: {
            key: (datum) => `${new Date(datum.x).toString()}`,
            value: (datum) => `${(datum.y / 1000000).toFixed(2)}ms`,
          }
        }
      }
    }
    chartDiv = <VChart spec={chartSpec} />
  }
  let descriptionDiv = null
  if (report && report.data) {
    descriptionDiv = <Descriptions>
      <Descriptions.Item itemKey="Address">{report.data.address}</Descriptions.Item>
      <Descriptions.Item itemKey="IP">{report.data.ip}</Descriptions.Item>
      <Descriptions.Item itemKey="PacketsSent">{report.data.packets_sent}</Descriptions.Item>
      <Descriptions.Item itemKey="PacketsRecv">{report.data.packets_recv}</Descriptions.Item>
      <Descriptions.Item itemKey="PacketLoss">{report.data.packet_loss}</Descriptions.Item>
    </Descriptions>
  }

  return (
    <div className="p-4">
      <div>
        <Title heading={4}>Ping</Title>
      </div>
      <div className="mt-4">
        <div className="flex items-center justify-center">
          <div className="flex">
            <div className="w-[480px]">
              <Select className="w-full" filter allowCreate value={selectedAddress} onChange={setSelectedAddress}>
                {optionDivs}
              </Select>
            </div>
            <div className="w-[64px] ml-2">
              <InputNumber hideButtons className="w-full" value={count} onChange={setCount} />
            </div>
            <div className="w-[64px] ml-2">
              <Button type="primary" onClick={onPingClicked}>Ping</Button>
            </div>
          </div>
        </div>
        <div>
          <div>
            {descriptionDiv}
          </div>
          <div>
            {chartDiv}
          </div>
        </div>
      </div>
    </div>
  )
}
