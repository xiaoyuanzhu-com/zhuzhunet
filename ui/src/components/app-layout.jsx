import React from 'react'
import { Layout } from '@douyinfe/semi-ui'
import { Outlet } from 'react-router'
import AppNav from './app-nav'
import AppFooter from './app-footer'

export default function AppLayout() {
  const { Header, Footer, Sider, Content } = Layout
  return (
    <Layout className="w-screen h-screen">
      <Sider>
        <AppNav />
      </Sider>
      <Layout>
        <Content>
          <Outlet />
        </Content>
        <Footer>
          <AppFooter />
        </Footer>
      </Layout>
    </Layout>
  )
}
