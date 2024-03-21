# consignas-go-db

## Despliegue

Para empezar la ejecucion de la apliocac, ejecutamos el siguiente comando ubicados en la raiz del proyecto

<pre><code> go run cmd/server/main.go </code></pre>

docker run --name digitalstore-mysql -p 33060:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql:latest
docker exec -i $id_generado digitalstore -uroot -psecret digitalstore < build_database.sql
