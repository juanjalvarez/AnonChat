export default class {

  constructor(url) {
    this.ws = new WebSocket(url)
    this.pending = [
      () => {
        this.ws.onmessage = this.handleMessage
      }
    ]
    this.open = false
    this.ws.onopen = () => {
      this.open = true
      this.pending.forEach(p => p())
    }
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
      console.log('failed to find handler for event type', type)
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
    const f = () => this.ws.send(JSON.stringify(event))
    if (this.open) {
      f()
    } else {
      this.pending.push(f)
    }
  }
}
