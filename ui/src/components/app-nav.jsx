import React from 'react'
import { Nav } from '@douyinfe/semi-ui'
import { IconCheckbox, IconConfig, IconHeart } from '@douyinfe/semi-icons-lab'
import { Link } from 'react-router'

export default function AppFooter() {
  const items = [
    {
      itemKey: 'test',
      text: 'Test Network',
      to: '/',
      icon: <IconCheckbox />,
    },
    {
      itemKey: 'monitoring',
      text: 'Monitoring',
      to: '/monitoring',
      icon: <IconHeart />,
    },
    {
      itemKey: 'settings',
      text: 'Settings',
      to: '/settings',
      icon: <IconConfig />,
    },
  ]
  const renderWrapper = ({ itemElement, isSubNav, isInSubNav, props }) => {
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
    <Nav className="h-full" items={items} renderWrapper={renderWrapper} footer={footer}>
    </Nav>
  )
}
