import React, { Component } from 'react'

import '../styles/joinchat.css'

export default class extends Component {

  state = {
    create: false,
    join: false,
    chatName: '',
    chatId: ''
  }

  selectCreate = () => {
    this.setState({
      create: true
    })
  }

  selectJoin = () => {
    this.setState({
      join: true
    })
  }

  handleChangeChatName = e => {
    this.setState({
      chatName: e.target.value
    })
  }

  handleChangeChatId = e => {
    this.setState({
      chatId: e.target.value
    })
  }

  handleCreate = () => {
    const createChat = this.props.createChat
    if (!Boolean(createChat)) {
      return
    }
    createChat(this.state.chatName)
  }

  handleJoin = () => {
    const joinChat = this.props.joinChat
    if (!Boolean(joinChat)) {
      return
    }
    joinChat(this.state.chatId)
  }

  render() {
    const { create, join, chatName, chatId } = this.state
    return (
      <div className="joinchat-container">
        {
          !create && !join ?
          <div className="joinchat-section-select">
            <div
              onClick={this.selectCreate}
              className="joinchat-section-selector joinchat-section-selectcreate"
            >
              Create Chat
            </div>
            <div
              onClick={this.selectJoin}
              className="joinchat-section-selector joinchat-section-selectjoin"
            >
              Join Chat
            </div>
          </div>
          : null
        }
        <div className={`joinchat-section ${create ? 'joinchat-section-selected' : ''}`}>
          <div className="joinchat-section-title">Create Chat</div>
          <input
            type="text"
            placeholder="Chat Name"
            className="joinchat-section-input"
            onChange={this.handleChangeChatName}
            value={chatName}
          />
          <button
            className="joinchat-section-submit"
            onClick={this.handleCreate}
          >
            Create
          </button>
        </div>
        <span className="vertical-line" />
        <div className={`joinchat-section ${join ? 'joinchat-section-selected' : ''}`}>
          <div
            className="joinchat-section-title"
            onClick={this.handleJoin}
          >
            Join Chat
          </div>
          <input
            type="text"
            placeholder="Chat ID"
            className="joinchat-section-input"
            onChange={this.handleChangeChatId}
            value={chatId}
          />
          <button className="joinchat-section-submit">Join</button>
        </div>
      </div>
    )
  }
}
