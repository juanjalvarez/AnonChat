import React, { Component } from 'react'
import { Icon } from 'react-fa'

import '../styles/modal.css'

export default class extends Component {

  render() {
    const { onClose, children } = this.props
    return (
      <div className="modal-space">
        <div className="modal-container">
          <div className="modal-header">
            <span />
            <Icon name="times" onClick={onClose} className="modal-close" />
          </div>
          <div className="modal-body">{children}</div>
        </div>
      </div>
    )
  }
}
