# datapoa
Simples CRUD de linhas e itinerários de POA

### Docker
Imagem disponivel ``docker pull docker.pkg.github.com/airtongit/datapoa/datapoa:latest``

Porta: 8000

Execução: ``docker run -p8000:8000 docker.pkg.github.com/airtongit/datapoa/datapoa:latest``

### Endpoints

* ws /ws
  * websocket, mensagem recebida indica novos dados disponíveis
    ```html
    <!DOCTYPE html>
    <html lang="en">
        <head>
            <title>Exemplo WebSocket</title>
        </head>
        <body>
            <pre id="data">{{.Data}}</pre>
            <script type="text/javascript">
                (function() {
                    var data = document.getElementById("data");
                    var conn = new WebSocket("ws://host.com/ws");
                    conn.onclose = function(evt) {
                        data.textContent = 'Connection closed';
                    }
                    conn.onmessage = function(evt) {
                        console.log('novos dados disponiveis');
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

