let host = window.location.host;

let ws = new WebSocket("ws://" + host + "/ws");
let content = document.getElementById("content");

ws.onmessage = msg => {
  let p = document.createElement("p");
  p.append(msg.data);
  content.append(p);
};

let form = document.getElementById("form");
form.addEventListener("submit", event => {
  event.preventDefault();
  let val = document.getElementById("input").value;
  console.log(val);
  ws.send(val);
});
