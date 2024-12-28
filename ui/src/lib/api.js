import axios from 'axios'

const getJSON = async (url) => {
  const res = await axios.get(url)
  return res.data
}

export default {
  getManifest: () => getJSON('/api/manifest'),
  getBrandList: () => getJSON('/api/brands'),
  getDNSList: () => getJSON('/api/dns'),
  getWebsiteList: () => getJSON('/api/websites'),
  ping: (address, count) => getJSON(`/api/ping?address=${address}&count=${count}&async=true`),
}
