import React, { Component } from 'react'

import Socket from '../shared/Socket'

import UserForm from '../components/UserForm'
import ChatList from '../components/ChatList'
import Modal from '../components/Modal'

import '../styles/app.css'

export default class extends Component {

  state = {
    modal: null,
    user: null,
    activeChat: null,
    chats: {
      a: {
        id: 'abc123',
        name: 'AnonChat Evangelists akwjdhlakwdhaklwjdhkljawd',
        users: ["1"],
        messages: [
          {
            userId: '1',
            text: 'me too!'
          }
        ],
        notifications: 4
      }
    },
    cachedUsers: {
      "1": {
        id: '1',
        name: 'Jane Doe'
      }
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
              users={this.state.cachedUsers}
              onSelectChat={this.handleChatChange}
              showModal={this.showModal}
            />
          </div>
          <div className={`app-body ${hasActiveChat ? '' : 'unfocus'}`}>
            <button onClick={() => this.setState({activeChat: null})}>back</button>
          </div>
        </div>
      </div>
    )
  }
}
