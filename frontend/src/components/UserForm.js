import React, { Component } from 'react'

import { Icon } from 'react-fa'

import '../styles/userform.css'

export default class extends Component {

  state = {
    text: 'anonymous'
  }

  onChange = e => {
    this.setState({
      text: e.target.value
    })
  }

  onKey = e => {
    if (e.key === 'Enter') {
      this.onSubmit()
    }
  }

  onBlur = e => {
    this.onSubmit()
  }

  onSubmit = () => {
    const { onSubmit } = this.props;
    if (onSubmit) {
      onSubmit(this.state.text)
    }
  }

  render() {
    return (
      <div className="userform-container">
        <input value={this.state.text} className="userform-input" onChange={this.onChange} onKeyPress={this.onKey} onBlur={this.onBlur} />
        <div className="userform-edit-wrapper">
          <Icon name="pencil-square-o" className="userform-edit" />
        </div>
      </div>
    )
  }
}
