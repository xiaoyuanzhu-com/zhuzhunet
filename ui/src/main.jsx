import { LocaleProvider } from '@douyinfe/semi-ui'
import en_US from '@douyinfe/semi-ui/lib/es/locale/source/en_US'
import AppRoutes from './components/app-routes.jsx'
import { initVChartSemiTheme } from '@visactor/vchart-semi-theme'
import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'

initVChartSemiTheme()

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
      <AppRoutes />
    </LocaleProvider>
  </StrictMode>,
)
