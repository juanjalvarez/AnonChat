import React, { Component } from 'react'
import { Icon } from 'react-fa'

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
    console.log(this.state.text)
    const { onSubmit } = this.props
    if (onSubmit) {
      onSubmit(this.state.text)
    }
  }

  render() {
    const { chat } = this.props
    return (
      <div className="chatbody-container">
        <div className="chatbody-header">
        <div className="chatbody-back">Back</div>
        <div></div>
        <div></div>
        </div>
        <div className="chatbody-body">
          <div>txt</div>
        </div>
        <div className="chatbody-footer">
          <input
            type="text"
            className="chatbody-input"
            placeholder="Type your message here"
            value={this.state.text}
            onChange={this.handleTextChange}
            onKeyPress={this.handleKeyPress}
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
