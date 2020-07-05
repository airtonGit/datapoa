# datapoa
Simples CRUD de linhas e itinerários de POA

### Endpoints

* ws /ws
  * websocket, mensagem recebida indica novos dados disponíveis
    ```html
    <!DOCTYPE html>
    <html lang="en">
        <head>
            <title>WebSocket Example</title>
        </head>
        <body>
            <pre id="fileData">{{.Data}}</pre>
            <script type="text/javascript">
                (function() {
                    var data = document.getElementById("fileData");
                    var conn = new WebSocket("ws://{{.Host}}/ws?lastMod={{.LastMod}}");
                    conn.onclose = function(evt) {
                        data.textContent = 'Connection closed';
                    }
                    conn.onmessage = function(evt) {
                        console.log('file updated');
                        data.textContent = evt.data;
                    }
                })();
            </script>
        </body>
    </html>
    ```
* GET /linhas
* GET /linhas/?nome={filtro}
  * Filtrar por nome case insensitive
  * Ex /linhas/?nome=navegantes
* GET /linhas/?raio={km}&lat={lat}&lng={lng}
  * Filtrar por raio em KM informado
  * Ex /linhas/?raio=2&lat=-30.124190574226&lng=-51.223783136207
* GET /linha/{codigo}

