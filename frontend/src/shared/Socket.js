export default class {

  constructor(url) {
    this.ws = new WebSocket(url)
    this.ws.onopen = e => {
      console.log(`Connected to ${url}`)
    }
    this.ws.onmessage = this.handleMessage
    this.handlers = new Map()
  }

  handleMessage = e => {
    const event = JSON.parse(e.data)
    if (!Boolean(event)) {
      return
    }
    const { type, data } = event
    const handler = this.handlers.get(type)
    if (!Boolean(handler)) {
      return
    }
    handler(data)
  }

  on = (event, handler) => {
    this.handlers.set(event, handler)
  }

  onClose = handler => {
    this.ws.onclose = handler
  }

  send = event => {
    this.ws.send(event)
  }
}
