import React, { Component } from 'react'
import { Icon } from 'react-fa'

import Comment from './Comment'

import '../styles/chatbody.css'

export default class extends Component {

  state = {
    text: ''
  }

  handleTextChange = e => {
    this.setState({
      text: e.target.value
    })
  }

  handleKeyPress = e => {
    if (e.key === 'Enter') {
      this.onSubmit()
    }
  }

  onSubmit = () => {
    if (this.props.disabled) {
      return
    }
    console.log(this.state.text)
    const { onSubmit } = this.props
    if (onSubmit) {
      onSubmit(this.state.text)
    }
    this.setState({
      text: ''
    })
  }

  render() {
    const { chat, user, disabled = false, handleGoBack } = this.props
    return (
      <div className="chatbody-container">
        <div className="chatbody-header">
          <div className="chatbody-back" onClick={handleGoBack}>Back</div>
          <div className="chatbody-title">{chat ? `${chat.name} @ ${chat.id}` : ''}</div>
          <div></div>
        </div>
        <div className="chatbody-body">
          {
            chat ? chat.messages.map((m, index) => {
              return (
                <Comment
                  key={index}
                  text={m.text}
                  timestamp={m.timestamp}
                  author={m.userId === user.id ? 'you' : chat.users[m.userId].name}
                  self={m.userId === user.id}
                />
              )
            }) : null
          }
        </div>
        <div className="chatbody-footer">
          <input
            type="text"
            className="chatbody-input"
            placeholder={disabled ? 'Select a chat...' : 'Type your message here'}
            value={this.state.text}
            onChange={this.handleTextChange}
            onKeyPress={this.handleKeyPress}
            disabled={disabled}
          />
          <Icon
            name="arrow-right"
            className="chatbody-send"
            onClick={this.onSubmit}
          />
        </div>
      </div>
    )
  }
}
