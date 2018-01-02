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
      const messages = chats[data.id] ? chats[data.id] : []
      const nc = new Chat(data.id, data.name, data.owner)
      nc.users = data.users
      chats[data.id] = nc
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
      activeChat: id
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

  handleJoinChat = id => {
    console.log('user requested to join chat with id', id)
    this.handleModalClose()
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
            />
          </div>
        </div>
      </div>
    )
  }
}
