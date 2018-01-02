import React, { Component } from 'react'
import { Icon } from 'react-fa'

import '../styles/chatitem.css'

export default class extends Component {

  onSelect = () => {
    const { id, onSelect } = this.props
    if (Boolean(id)) {
      onSelect(id)
    }
  }

  render() {
    const { name, members, notifications, lastMessage, lastUser } = this.props
    return (
      <div className="chatitem-container" onClick={this.onSelect}>
        <div className="chatitem-header">
          <div className="chatitem-name nowrap">{name}</div>
          <div className="chatitem-counts">
            {
              notifications > 0 ?
              <span className="chatitem-notifications nowrap"><span>{notifications}</span></span>
              : null
            }
            <span className="nowrap"><Icon name="user" />{members}</span>
          </div>
        </div>
        {
          lastMessage ?
          <div className="chatitem-lastmessage">
            <div>{lastUser}</div>
            <div>{lastMessage}</div>
          </div>
          : null
        }
      </div>
    )
  }
}
