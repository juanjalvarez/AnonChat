import React, { Component } from 'react'

import ChatItem from './ChatItem'

export default class extends Component {

  handleSelectChat = id => {
    console.log(id)
  }

  render() {
    const { chats } = this.props
    if (Boolean(chats)) {
      return (
        <div>
          {
            Object.keys(chats).map(c => {
              const chat = chats[c]
              return (
                <ChatItem key={chat.id} chat={chat} onSelect={this.handleSelectChat} />
              )
            })
          }
        </div>
      )
    }
  }
}
