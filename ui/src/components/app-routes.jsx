import { BrowserRouter, Route, Routes } from 'react-router'
import AppLayout from './app-layout.jsx'
import DiagnoseDNS from '@/pages/diagnose/diagnose-dns.jsx'
import DiagnosePing from '@/pages/diagnose/diagnose-ping.jsx'
import Monitoring from '@/pages/monitoring/monitoring.jsx'
import Settings from '@/pages/settings/settings.jsx'
import Test from '@/pages/test/test.jsx'

export default function AppRoutes() {
  return (
    <BrowserRouter>
      <Routes>
        <Route element={<AppLayout />} >
          <Route index element={<Test />} />
          <Route path="/monitoring" element={<Monitoring />} />
          <Route path="/diagnose/dns" element={<DiagnoseDNS />} />
          <Route path="/diagnose/ping" element={<DiagnosePing />} />
          <Route path="/settings" element={<Settings />} />
        </Route>
      </Routes>
    </BrowserRouter>
  )
}
