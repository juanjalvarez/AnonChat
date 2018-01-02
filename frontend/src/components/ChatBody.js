import React, { Component } from 'react'

import '../styles/chatbody.css'

export default class extends Component {

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
          <input type="text" />
        </div>
      </div>
    )
  }
}
