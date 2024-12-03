import React from 'react'
import { Nav } from '@douyinfe/semi-ui'
import { IconGettingStarted, IconConfig } from '@douyinfe/semi-icons-lab'
import { Link } from 'react-router'

export default function AppFooter() {
  const items = [
    {
      itemKey: 'test',
      text: 'Test Network',
      to: '/',
      icon: <IconGettingStarted />,
    },
    {
      itemKey: 'settings',
      text: 'Settings',
      to: '/settings',
      icon: <IconConfig />,
    },
  ]
  const renderWrapper = ({ itemElement, isSubNav, isInSubNav, props }) => {
    console.log(props)
    return (
      <Link to={props.to}>
        {itemElement}
      </Link>
    )
  }
  const footer = {
    collapseButton: true,
  }
  return (
    <Nav items={items} renderWrapper={renderWrapper} footer={footer}>
    </Nav>
  )
}
