import React, { Component } from 'react'

import UserForm from '../components/UserForm'
import ChatList from '../components/ChatList'

import '../styles/app.css'

export default class extends Component {

  state = {
    activeChat: false,
    chats: {
      a: {
        id: 'abc123',
        name: 'AnonChat Evangelists',
        users: {
          '1': {
            id: '1',
            name: 'Jane Doe'
          }
        },
        messages: [
          {
            userId: '1',
            text: 'me too!'
          }
        ],
        notifications: 4
      }
    }
  }

  handleUserChange = name => {
    console.log(name)
  }
  
  render() {
    return (
      <div className="app-container">
        <div className={`app-nav ${this.state.activeChat ? 'unfocus' : ''}`}>
          <UserForm
            onSubmit={this.handleUserChange}
          />
          <ChatList
            chats={this.state.chats}
          />
          <button onClick={() => {this.setState({activeChat: true})}}>enter chat</button>
        </div>
        <div className={`app-body ${this.state.activeChat ? '' : 'unfocus'}`}>
          body
          <button onClick={() => this.setState({activeChat: false})}>back</button>
        </div>
      </div>
    )
  }
}
