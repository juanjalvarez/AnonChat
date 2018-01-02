import React, { Component } from 'react'

import '../styles/joinchat.css'

export default class extends Component {

  state = {
    create: false,
    join: false
  }

  render() {
    return (
      <div className="joinchat-container">
        <div className="joinchat-create-container">
          <div>Create Chat</div>
          <input type="text" placeholder="Chat Name" />
          <button>Create</button>
        </div>
        <div className="joinchat-join-container">
          <div>Join Chat</div>
          <input type="text" placeholder="Chat Name" />
          <button>Join</button>
        </div>
      </div>
    )
  }
}
