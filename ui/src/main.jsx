import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter, Routes, Route } from 'react-router'
import './index.css'
import AppLayout from './components/app-layout.jsx'
import { LocaleProvider } from '@douyinfe/semi-ui'
import en_US from '@douyinfe/semi-ui/lib/es/locale/source/en_US'
import Test from './pages/test/test.jsx'
import Settings from './pages/settings/settings.jsx'

function setDarkMode(dark) {
  const body = document.body
  if (dark) {
    if (!body.hasAttribute('theme-mode')) {
      body.setAttribute('theme-mode', 'dark')
    }
  } else {
    if (body.hasAttribute('theme-mode')) {
      body.removeAttribute('theme-mode')
    }
  }
}

(() => {
  const mql = window.matchMedia('(prefers-color-scheme: dark)')
  setDarkMode(mql.matches)
  mql.addEventListener('change', (e) => setDarkMode(e.matches))
})()

const root = document.getElementById("root")

createRoot(root).render(
  <StrictMode>
    <LocaleProvider locale={en_US}>
      <BrowserRouter>
        <Routes>
          <Route element={<AppLayout />} >
            <Route index element={<Test />} />
            <Route path="/settings" element={<Settings />} />
          </Route>
        </Routes>
      </BrowserRouter>
    </LocaleProvider>
  </StrictMode>,
)
