var output = null, btnConnect = null, btnDisconnect = null, btnSend = null;

function enableSendMessage (active) {
    var loc = document.getElementById('location'),
        msg = document.getElementById('message');

    loc.disabled = active;
    btnConnect.disabled = active;
    btnDisconnect.disabled = !active;
    btnSend.disabled = !active;
    msg.disabled = !active
}

function onOpen () {
    writeToScreen('<span style="color:green;">[ CONNECTED ]</span>');
    enableSendMessage(true)
}

function onClose () {
    writeToScreen("<span>[ DISCONNECTED ]</span>");
    enableSendMessage(false)
}

function onMessage (e) {
    writeToScreen('<span style="color:blue;">RESPONSE: ' + e.data + '</span>');
}

function onError (e) {
    writeToScreen('<span style="color:red">ERROR:</span>' + e.data)
}

function init () {
    output = document.getElementById('output');

    btnConnect = document.getElementById('btnConnect');
    btnConnect.onclick = doConnect;

    btnDisconnect = document.getElementById('btnDisconnect');
    btnDisconnect.onclick = doDisconnect;

    btnSend = document.getElementById('btnSend');
    btnSend.onclick = doSend
}

function sendMessage (message) {
    writeToScreen('SENT: ' + message);
    client.send(message)
}

function writeToScreen (message) {
    var p = document.createElement('p');
    p.style.wordWrap = 'break-word';
    p.innerHTML = message;
    output.appendChild(p)
}

function doConnect () {
    var location = document.getElementById('location').value;

    while (output.firstChild) {
        output.removeChild(output.firstChild);
    }

    client = new WebSocket(location);
    client.onopen = function (e) { onOpen() };
    client.onclose = function (e) { onClose() };
    client.onmessage = function (e) { onMessage(e) };
    client.onerror = function (e) { onError(e) }
}

function doDisconnect () {
    document.getElementById('message').value = '';
    client.close();
}

function doSend () {
    var message = document.getElementById('message').value;
    if (message !== '') {
        sendMessage(message)
    }
}

window.addEventListener('load', init, false);
