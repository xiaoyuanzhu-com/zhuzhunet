import React from 'react'
import { Nav } from '@douyinfe/semi-ui'
import { IconCheckbox, IconRating, IconToast, IconConfig, IconHeart, IconGettingStarted } from '@douyinfe/semi-icons-lab'
import { Link } from 'react-router'

export default function AppFooter() {
  const items = [
    {
      itemKey: 'test',
      text: 'Test My Network',
      to: '/',
      icon: <IconGettingStarted />,
    },
    {
      itemKey: 'monitoring',
      text: 'Monitoring',
      to: '/monitoring',
      icon: <IconHeart />,
    },
    {
      itemKey: 'diagnose',
      text: 'Diagnose',
      to: '/diagnose',
      icon: <IconToast />,
      items: [
        {
          itemKey: 'dns',
          text: 'DNS',
          to: '/diagnose/dns',
          icon: <IconRating />,
        },
      ],
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
    <Nav className="h-full" items={items} renderWrapper={renderWrapper} footer={footer} limitIndent={false}>
    </Nav>
  )
}
