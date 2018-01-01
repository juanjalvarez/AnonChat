import React from 'react'
import ReactDOM from 'react-dom'

ReactDOM.render(
  <div>
    <div>{JSON.stringify(window.location)}</div>
  </div>,
  document.getElementById('root')
)