import React, { Component } from 'react'
import { Icon } from 'react-fa'

import '../styles/chatitem.css'

export default class extends Component {

  onSelect = () => {
    const { chat, onSelect } = this.props
    const { id } = chat
    if (Boolean(id)) {
      onSelect(id)
    }
  }

  render() {
    const { chat, onSelect } = this.props
    const { id = "", name = "N/A", messages = [], users = {}, notifications = 0 } = chat
    const lastMessage = messages.length > 0 ? messages[messages.length - 1] : null
    return (
      <div className="chatitem-container" onClick={this.onSelect}>
        <div className="chatitem-header">
          <div className="chatitem-name">{name}</div>
          <div className="chatitem-counts">
            {
              notifications > 0 ?
              <span className="chatitem-notifications"><span>{notifications}</span></span>
              : null
            }
            <span><Icon name="user" /><span>{Object.keys(users).length}</span></span>
          </div>
        </div>
        {
          lastMessage ?
          <div className="chatitem-lastmessage">
            <div>{users[lastMessage.userId].name}</div>
            <div>{lastMessage.text}</div>
          </div>
          : null
        }
      </div>
    )
  }
}
