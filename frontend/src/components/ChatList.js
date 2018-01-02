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
    const { showModal, createChat, joinChat } = this.props
    if (showModal) {
      showModal(<JoinChat
        joinChat={joinChat}
        createChat={createChat}
      />)
    }
  }

  render() {
    const { chats } = this.props
    if (Boolean(chats)) {
      return (
        <div className="chatlist-container">
          <div className="chatlist-wrapper">
            {
              Object.keys(chats).map(c => {
                const chat = chats[c]
                const lastMessage = chat.messages.length === 0 ? null : chat.messages[chat.messages.length - 1]
                const lastUser = ''
                // TODO: Change to new data model
                return (
                  <ChatItem
                    key={chat.id}
                    id={chat.id}
                    name={chat.name}
                    members={chat.users.length}
                    notifications={chat.notifications}
                    lastMessage={lastMessage ? lastMessage.text : ''}
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
