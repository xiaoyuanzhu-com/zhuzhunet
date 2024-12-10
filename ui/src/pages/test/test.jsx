import React from 'react'
import { Nav, Typography } from '@douyinfe/semi-ui'
import TestDNS from './test-dns'
import TestLatency from './test-latency'
import TestSpeed from './test-speed'
import TestIP from './test-ip'
import TestAccess from './test-access'
import { useState, useEffect } from 'react'
import api from '../../lib/api'
const { Text } = Typography

export default function Test() {
  const [manifest, setManifest] = useState(null)
  const [brands, setBrands] = useState(null)
  const fetchManifest = async () => { 
    const manifest = await api.getManifest()
    setManifest(manifest)
  }
  const fetchBrands = async () => {
    const brands = await api.getBrands()
    setBrands(brands)
  }
  useEffect(() => {
    fetchManifest()
    fetchBrands()
  }, [])
  return (
    <div>
      <TestLatency manifest={manifest} brands={brands} />
      <TestSpeed manifest={manifest} brands={brands} />
      <TestAccess manifest={manifest} brands={brands} />
      <TestIP manifest={manifest} brands={brands} />
      <TestDNS manifest={manifest} brands={brands} />
    </div>
  )
}
