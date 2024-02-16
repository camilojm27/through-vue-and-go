This GitHub repository contains a project that allows analyzing the email database of the Enron company, which was released after the financial scandal it was involved in. The project consists of three main parts: an indexer, a server, and a web interface.

The indexer is a program that reads email files and enters them into Zincsearch, a database. The indexer uses Golang to read the files and send relevant email data, such as sender, recipient, subject, content, and date, to Zincsearch.

The server is a web application that communicates with Zincsearch and returns the information requested by the web interface. The server uses the Chi framework to create a REST API that receives queries from the web interface, sends them to Elasticsearch, processes the results, and returns them in JSON format.

The web interface is an application that visualizes email information and allows for general searches. The web interface is built with Vue.js, a JavaScript framework for creating reactive user interfaces. The web interface displays a table with emails and their respective fields.

# Requirements

- [zincsearch](https://zincsearch-docs.zinc.dev/)
- Go 1.21+
- Nodejs 18+
- [Enron Mail Dataset (423MB)](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz)

# Installation

## Indexer

- Install and run the zincsearch tool.
- Rename the `.env.example` file to `.env` and add the environment variables as needed.
- Download the Enron database into the root folder of this repository.
- Compile the indexer: navigate to the folder and run the command `go build main.go`. Then, execute the program with `./main.go`.

## Backend

- Rename the `.env.example` file to `.env` and add the environment variables as needed.
- Compile and run `go build main.go`, then execute the program with `./main.go`.

## Frontend

- Rename the `.env.example` file to `.env` and add the environment variables as needed.
- Install the necessary dependencies with `npm install`.
- Run the development server with `npm run dev`.

---

Este repositorio de GitHub contiene un proyecto que permite analizar la base de datos de correos electrónicos de la empresa Enron, que fue publicada tras el escándalo financiero que la involucró. El proyecto consta de tres partes principales: un indexador, un servidor y una interfaz web.

El indexador es un programa que lee los archivos de los correos electrónicos y los ingresa en zincsearch, una base de datos . El indexador utiliza golang para leer los archivos y enviar los datos relevantes de los correos, como el remitente, el destinatario, el asunto, el contenido, y la fecha.

El servidor es una aplicación web que se comunica con zincsearch y retorna la información solicitada por la interfaz web. El servidor utiliza el framework chi para crear una API REST que recibe las consultas de la interfaz web y las envía a elasticsearch, luego procesa los resultados y los devuelve en formato JSON.

La interfaz web es una aplicación que visualiza la información de los correos electrónicos y permite hacer búsquedas generales. La interfaz web está hecha con Vue.js, un framework de JavaScript para crear interfaces de usuario reactivas. La interfaz web muestra una tabla con los correos electrónicos y sus campos.

# Requerimientos

- [zincsearch](https://zincsearch-docs.zinc.dev/)
- Go 21+
- Nodejs 18+
- [Enron Mail Dataset (423MB)](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz)

# Instalacion

## Indexer

- Instalar y ejecutar la herramienta zincsearch.
- Renombrar el archivo .env.example a .env y agregar las variables de entorno según el caso.
- Descargar la base de datos enron en la carpeta raíz de este repositorio.
- Compilar el indexador: ingresar a la carpeta y ejecutar el comando `go build main.go`. Luego, ejecutar el programa con `./main.go`

## Backend

- Renombrar el archivo `.env.example` a `.env` y agregar las variables de entorno según el caso.
- Compilar y ejecutar `go build main.go`, después ejecutar el programa con `./main.go`

## Frontend

- Renombrar el archivo .env.example a .env y agregar las variables de entorno según el caso.
- Instalar las dependencias necesarias con `npm install`
- Ejecutar el servidor de desarrollo con `npm run dev`
