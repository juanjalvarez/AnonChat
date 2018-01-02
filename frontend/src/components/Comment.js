import React, { Component } from 'react'

import '../styles/comment.css'

export default class extends Component {
  
  render() {
    const { text, author, timestamp, self } = this.props
    const date = new Date(timestamp)
    const dateString = `${date.getHours()}:${date.getMinutes()}`
    return (
      <div className={`comment-container ${self ? 'comment-self' : ''}`}>
        <div className="comment-header">
          <span>{author}</span>
          <span>{dateString}</span>
        </div>
        <div className="comment-text">{text}</div>
      </div>
    )
  }
}
