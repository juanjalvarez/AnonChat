import React, { Component } from 'react'

import ChatItem from './ChatItem'
import JoinChat from './JoinChat'

import '../styles/chatlist.css'

export default class extends Component {

  handleSelectChat = id => {
    const { onSelectChat } = this.props
    if (onSelectChat) {
      onSelectChat(id)
    }
  }

  showModal = () => {
    const { showModal } = this.props
    if (showModal) {
      showModal(<JoinChat />)
    }
  }

  render() {
    const { users, chats } = this.props
    if (Boolean(chats)) {
      return (
        <div className="chatlist-container">
          <div className="chatlist-wrapper">
            {
              Object.keys(chats).map(c => {
                const chat = chats[c]
                const lastMessage = chat.messages[chat.messages.length - 1]
                const lastUser = users[lastMessage.userId].name
                return (
                  <ChatItem
                    key={chat.id}
                    id={chat.id}
                    name={chat.name}
                    members={chat.users.length}
                    notifications={chat.notifications}
                    lastMessage={lastMessage.text}
                    lastUser={lastUser}
                    onSelect={this.handleSelectChat}
                  />
                )
              })
            }
          </div>
          <div className="chatlist-add" onClick={this.showModal}>+</div>
        </div>
      )
    }
  }
}
