package httpsrv

import (
	"html/template"
	"net/http"
)

func (s *Server) handlerHome(w http.ResponseWriter, r *http.Request) {
	csrfToken, err := GenerateCSRFToken()
	if err != nil {
		http.Error(w, "Could not generate CSRF token", http.StatusInternalServerError)
		return
	}

	// Set the CSRF token as a cookie
	SetCSRFCookie(w, csrfToken)

	varmap := map[string]interface{}{
		"host":                  "ws://" + r.Host + "/goapp/ws?csrf_token=" + csrfToken,
		"numOfConnections":      s.numOfConnections,
		"concurrentConnections": s.concurrentWSConnections,
	}

	template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws = [];
    var numConnections = Number("{{.numOfConnections}}");
    var concurrentConnections = "{{.concurrentConnections}}";
    var print = function(message, i) {
        let prefix = "";
        if (concurrentConnections === "true" && numConnections > 1) {
            prefix = "[conn #" + i + "] ";
        }
        var d = document.createElement("div");
        d.textContent = prefix + message;
        output.appendChild(d);
        output.scroll(0, output.scrollHeight);
    };
    document.getElementById("open").onclick = function(evt) {
        if (ws.length !== 0) {
            return false;
        }
        for (let i = 0; i < numConnections; i++) {
            let newWs = new WebSocket("{{.host}}");
            ws.push(newWs);
            newWs.onopen = function(evt) {
                print("OPEN", i);
            }
            newWs.onclose = function(evt) {
                print("CLOSE", i);
            }
            newWs.onmessage = function(evt) {
                print("RESPONSE: " + evt.data, i);
            }
            newWs.onerror = function(evt) {
                print("ERROR: " + evt.data, i);
            }
        }
        return false;
    };
    document.getElementById("send").onclick = function(evt) {
        if (ws.length === 0) {
            return false;
        }
        for (let i = 0; i < numConnections; i++) {
            print("SEND: " + input.value, i);
            ws[i].send(input.value);
        }
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
        if (ws.length === 0) {
            return false;
        }
        ws.forEach((websocket) => websocket.close());
        ws = [];
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="{}">
<button id="send">Reset</button>
</form>
</td><td valign="top" width="50%">
<div id="output" style="max-height: 70vh;overflow-y: scroll;"></div>
</td></tr></table>
</body>
</html>
`)).Execute(w, varmap)
}
