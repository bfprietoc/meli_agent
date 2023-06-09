# meli_agent

# Prueba Mercadolibre  - Compliance continuo de servidores
Prueba tecnica para para MercadoLibre (Backend). 
Basicamente dividi el problema en 2 partes, asi que se crearon dos repos para la solucion, uno donde se encuentra la aplicacion-server la cual se consume y que almacena y consulta la informacion solicitada, desplegado en una instancia de AWS que consume tambien una instancia Mysql desplegada en AWS.  La segunda parte es un archivo ejecutable de acuerdo a la arquitectura del sistema operativo, este es quien recolecta la informacion y la envia al servidor para ser almacenada, ambos repos estan desarrollados en Golang.

## Si se quiere ejecutar:
### Base de datos (Mysql)
1. Descargamos y ejecutamos la imagen docker de Mysql con el siguiente comando:
``` 
docker run -d \
    -e MYSQL_ROOT_PASSWORD=root \
    -e MYSQL_DATABASE=meli \
    -e MYSQL_USER=admin \
    -e MYSQL_PASSWORD=Meli2023 \
    --name mysql-container \
    mysql/mysql-server:latest
```

3. Conectados a la db, procedemos a crear tabla con los campos necesarios, para eso se ejecutan la siquiente query:
```
CREATE TABLE IF NOT EXISTS server_data (
	id serial PRIMARY KEY,	
	cpu_info VARCHAR (100),
	process_info JSON,
	users_info JSON,
	os_info VARCHAR(100)
);
```
La base de datos se expone por defecto en el puerto 3306, asi que usamos el mismo para la conexion.

5. Finalmente para el despliegue en la nube, se creo una instancia RDS de AWS para la creacion de la base de datos mysql, la cual estara disponible para no tener que crear la DB desde cero.

## App Go
1. Descargar o clonar el repositorio a la maquina donde se desea correr.
2. Ubicarse en el directorio meli_agent
3. Ejecutar el siguiente comando para construir la imagen docker:
```
docker build -t meli .
```
4. Luego se puede correr la imagen ejecutando:
```
docker run -p 8080:8080 meli
```

5. La aplicacion se ejecutara en la ruta http://localhost:8080

6. Se crearon dos endpoints, uno tipo GET donde se consutla la informacion almacenada en la base de datos y el otro tipo POST para la creacion de los datos, ambos se hacen contra la url http://localhost:8080/data.

Tambien si se quiere, se puede consultar la informacion en el endpoint que ya se encuentra desplegado en AWS

http://ec2-18-188-229-228.us-east-2.compute.amazonaws.com:8080/data

La misma Url anterior es la que se consume para almacenar la informacion pero con el metodo POST y la estructura json especifica:

```
{
    "cpu_info": "Intel(R) Core(TM) i30-7360U CPU @ 2.30GHz\n",
    "process_info": "[{\"pid\":\"1\",\"name\":\"/sbin/launchd\"}]",
    "users_info": "[\"fabianprieto     console      Mar  8 19:56 \",\"fabianprieto     ttys001      Mar  9 22:11 \",\"\"]",
    "os_info": "osInfo"
}
```
* En el campo cpu_info se almacena la informacion del procesador.
* En el campo process_info se almacena la informacion de los procesos que se encuentran activos, se almacena un arreglo con objetos que tienen el PID y el nombre del proceso.
* En el campo users_info se almacena la informacion de los usuarios activos, contando las terminales en ejecucion.
* Finalmente en el campo os_info almaceno la informacion del sistema operativo, su nombre y version. 

## Docker 
1. Se usa docker para la gestion y el despliegue de la aplicacion que almacena y consulta la informacion enviada por el agente. 
3. Se creo una instancia EC2 en AWS con una maquina Ubuntu y se desplego la imagen de docker alli con los pasos descritos anteriormente.

## Implementacion y tecnologias usadas

- [Golang 1.18](https://go.dev/)
- [MySql](https://www.mysql.com)
- [Echo](https://github.com/labstack/echo)
- [Docker](https://www.docker.com)
- [AWS](https://aws.amazon.com/)


## Agente

Link al repo del agente:

https://github.com/bfprietoc/agent

1. Se crea un script en Golang para extraer la informacion solicitada por medio de comandos del sistema (Dependiendo el sistema operativo se usa uno u otro).
2. El script reune la informacion y la envia directamente al servidor (Se elige enviarla inmediatamente y no almacenarla para despues procesarla) donde se almacena por medio de un metodo POST definido, inmediatamente despues de la ejecucion, la libreria curl se usa para hacer el consumo del endpoint desplegado.
3. Se agregan diferentes archivos ejecutables del agente de acuerdo a la arquitectura y sistema operativo, para nuestro caso se usa el file linuxAmd64_agent, el cual es usado en debian como se solicito. 
4. Adicional se deja el archivo main.go donde se encuentra la logica usada para crear el script ejecutable.


### Requisitos

En el sistema operativo se la maquina donde se peinsa ejecutar el agente, es necesario tener instalado una libreria llamada curl, con la que se hace el envio de la informacion al servidor, adicional a eso es necesario dar algunos permisos extrar al script ejecutable.

### Instalacion de curl
Ejecutar el siguiente comando en la consola

```
sudo apt install curl
```
Verificar la instalacion con el comando:
```
curl --version
```
Nota. Normalmente esta libreria se instala cuando se hace un update de todo el sistema operativo.

### Otorgar permisos de ejecucion

Desde la terminal, nos situamos en la ubicacion del file descargado y ejecutamos el siguiente comando:

```
sudo chmod +x linuxAmd64_agent
```
### Ejecucion

1. Se descarga el ejecutable especifico para la arquitectura, en nuestro caso usaremos una maquina virtual con Debian 11 entonces elegimos el file linuxAmd64_agent. 
2. Abrimos la terminal y nos ubicamos en la ruta donde se encuentra el archivo descargado, procedemos a ejecutalo de la siguiente manera:

```
./linuxAmd64_agent
```
3. Obtenemos un mensaje en la consola se Successfull cuando termine la ejecucion.

### Consideraciones
1. Para macOs es necesario dar permisos adicionales, para ellos ejecutamos el archivo que corresponde a la arquitectura utilizada, normalmente salta un error en la ejecucion, cerramos la patalla que informa del error, procedemos a ir a las configuraciones de seguridad y vemos un mensaje que nos pregunta si confiamos en el archivo, procedemos a aceptarlo y  automaticamente se vuelve a ejecutar.



