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
    return (
      <div className="chatlist-container">
        {
          Boolean(chats) && Object.keys(chats).length > 0 ?
          <div className="chatlist-wrapper">
            {
              Object.keys(chats).map(c => {
                const chat = chats[c]
                const lastMessage = chat.messages.length === 0 ? null : chat.messages[chat.messages.length - 1]
                const lastUser = chat.messages.length === 0 ? null : chat.users[lastMessage.userId].name
                // TODO: Change to new data model
                return (
                  <ChatItem
                    key={chat.id}
                    id={chat.id}
                    name={chat.name}
                    members={Object.keys(chat.users).length}
                    notifications={chat.notifications}
                    lastMessage={lastMessage ? lastMessage.text : ''}
                    lastUser={lastUser}
                    onSelect={this.handleSelectChat}
                  />
                )
              })
            }
          </div>
          :
          <div className="chatlist-nochat">
            <span>You are not a member of any chat(s), please create/join one below.</span>
          </div>
        }
        <div className="chatlist-add" onClick={this.showModal}>+</div>
      </div>
    )
  }
}
