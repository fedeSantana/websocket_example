let host = window.location.host

let ws = new WebSocket('ws://' + host + '/ws')
let main = document.getElementById('main')

ws.onmessage((msg) => {
    let p = document.createElement('p')
    p.append(msg.data)
})