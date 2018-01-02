import React, { Component } from 'react'

import Socket from '../shared/Socket'
import Chat from '../shared/Chat'

import UserForm from '../components/UserForm'
import ChatList from '../components/ChatList'
import Modal from '../components/Modal'
import ChatBody from '../components/ChatBody'

import '../styles/app.css'

export default class extends Component {

  state = {
    modal: null,
    user: null,
    activeChat: null,
    chats: {
    }
  }

  constructor() {
    super()
    const token = localStorage.getItem('anonChatToken')
    const ws = this.socket = new Socket('ws://localhost:4000')
    let payload = {
      newUser: true
    }
    if (Boolean(token)) {
      payload = {
        token
      }
    }
    ws.send({
      type: 'authenticate',
      data: payload
    })
    ws.on('authenticate', data => {
      localStorage.setItem('anonChatToken', data.token)
      this.setState({
        user: {
          id: data.id,
          name: data.name
        }
      })
    })
    ws.on('set_user', data => {
      if (data.userId === this.state.user.id) {
        this.setState({
          user: {
            ...data
          }
        })
      }
    })
    ws.on('chat_status', data => {
      const chats = this.state.chats
      const nc = new Chat(data.id, data.name, data.owner)
      nc.users = data.users
      if (chats[data.id]) {
        nc.notificatios = chats[data.id].notifications
        nc.messages = chats[data.id].messages
      }
      chats[data.id] = nc
      this.setState({
        chats
      })
    })
    ws.on('message', data => {
      const chats = this.state.chats
      const messages = chats[data.chatId] ? chats[data.chatId].messages : []
      messages.push(data)
      chats[data.chatId].messages = messages
      if (data.chatId !== this.state.activeChat) {
        chats[data.chatId].notifications++
      }
      this.setState({
        chats
      })
    })
  }

  handleUserChange = name => {
    this.socket.send({
      type: 'set_user',
      data: {
        name
      }
    })
  }

  handleChatChange = id => {
    this.setState({
      activeChat: id,
      chats: {
        ...this.state.chats,
        [id]: {
          ...this.state.chats[id],
          notifications: 0
        }
      }
    })
  }

  handleGoBack = () => {
    this.setState({
      activeChat: null
    })
  }

  showModal = modal => {
    this.setState({
      modal
    })
  }

  handleModalClose = () => {
    this.setState({
      modal: null
    })
  }

  handleCreateChat = name => {
    this.handleModalClose()
    this.socket.send({
      type: 'new_chat',
      data: {
        name
      }
    })
  }

  handleSendMessage = text => {
    this.socket.send({
      type: 'message',
      data: {
        text,
        chatId: this.state.activeChat
      }
    })
  }

  handleJoinChat = chatId => {
    console.log('sending join request')
    this.handleModalClose()
    this.socket.send({
      type: 'join_chat',
      data: {
        chatId
      }
    })
  }
  
  render() {
    const hasActiveChat = Boolean(this.state.activeChat)
    return (
      <div>
        {
          this.state.modal ?
          <Modal onClose={this.handleModalClose}>
            {this.state.modal}
          </Modal>
          : null
        }
        <div className={`app-container ${Boolean(this.state.modal) ? 'blur' : ''}`}>
          <div className={`app-nav ${hasActiveChat ? 'unfocus' : ''}`}>
            <UserForm
              onSubmit={this.handleUserChange}
              user={this.state.user}
            />
            <ChatList
              chats={this.state.chats}
              onSelectChat={this.handleChatChange}
              showModal={this.showModal}
              createChat={this.handleCreateChat}
              joinChat={this.handleJoinChat}
            />
          </div>
          <div className={`app-body ${hasActiveChat ? '' : 'unfocus'}`}>
            <ChatBody
              chat={this.state.chats[this.state.activeChat]}
              onSubmit={this.handleSendMessage}
              user={this.state.user}
              disabled={!Boolean(this.state.activeChat)}
              handleGoBack={this.handleGoBack}
            />
          </div>
        </div>
      </div>
    )
  }
}
