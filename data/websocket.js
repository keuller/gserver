var output = null, btnConnect = null, btnDisconnect = null, btnSend = null

function toggle (active) {
    var loc = document.getElementById('location'),
        msg = document.getElementById('message')

    loc.disabled = active
    btnConnect.disabled = active
    btnDisconnect.disabled = !active
    btnSend.disabled = !active
    msg.disabled = !active
}

function onOpen (e) {
    writeToScreen('<span style="color:green;">[ CONNECTED ]</span>')
    toggle(true)
}

function onClose (e) {
    writeToScreen("<span>[ DISCONNECTED ]</span>");
    toggle(false)
}

function onMessage (e) {
    writeToScreen('<span style="color:blue;">RESPONSE: ' + e.data + '</span>');
}

function onError (e) {
    writeToScreen('<span style="color:red">ERROR:</span> ' + e.data)
}

function init () {
    output = document.getElementById('output')

    btnConnect = document.getElementById('btnConnect')
    btnConnect.onclick = doConnect

    btnDisconnect = document.getElementById('btnDisconnect')
    btnDisconnect.onclick = doDisconnect

    btnSend = document.getElementById('btnSend')
    btnSend.onclick = doSend
}

function sendMessage (message) {
    writeToScreen('SENT: ' + message)
    client.send(message)
}

function writeToScreen (message) {
    var p = document.createElement('p')
    p.style.wordWrap = 'break-word'
    p.innerHTML = message
    output.appendChild(p)
}

function doConnect (e) {
    var location = document.getElementById('location').value

    while (output.firstChild) {
        output.removeChild(output.firstChild);
    }

    client = new WebSocket(location)
    client.onopen = function (e) { onOpen(e) }
    client.onclose = function (e) { onClose(e) }
    client.onmessage = function (e) { onMessage(e) }
    client.onerror = function (e) { onError(e) }
}

function doDisconnect (e) {
    client.close()
    document.getElementById('message').value = ''
}

function doSend (e) {
    var message = document.getElementById('message').value
    if (message !== '') {
        sendMessage(message)
    }
}

window.addEventListener('load', init, false)
