# datapoa
Simples CRUD de linhas e itinerários de POA

### Endpoints

* ws /ws
  * Novas mensagens indicam novos dados disponíveis
* GET /linhas
* GET /linhas/?nome={filtro}
  * Filtrar por nome case insensitive
  * Ex /linhas/?nome=navegantes
* GET /linhas/?raio={km}&lat={lat}&lng={lng}
  * Filtrar por raio em KM informado
  * Ex /linhas/?raio=2&lat=-30.124190574226&lng=-51.223783136207
* GET /linha/{codigo}

