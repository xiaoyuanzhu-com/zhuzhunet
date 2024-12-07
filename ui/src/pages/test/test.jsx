import React from 'react'
import { Nav, Typography } from '@douyinfe/semi-ui'
import TestDNS from './test-dns'
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
      <TestDNS manifest={manifest} brands={brands} />
    </div>
  )
}
