This GitHub repository contains a project that allows analyzing the email database of the Enron company, which was released after the financial scandal it was involved in. The project consists of three main parts: an indexer, a server, and a web interface.

The indexer is a program that reads email files and enters them into Zincsearch, a database. The indexer uses Golang to read the files and send relevant email data, such as sender, recipient, subject, content, and date, to Zincsearch.

The server is a web application that communicates with Zincsearch and returns the information requested by the web interface. The server uses the Chi framework to create a REST API that receives queries from the web interface, sends them to Elasticsearch, processes the results, and returns them in JSON format.

The web interface is an application that visualizes email information and allows for general searches. The web interface is built with Vue.js, a JavaScript framework for creating reactive user interfaces. The web interface displays a table with emails and their respective fields.

---

Este repositorio de GitHub contiene un proyecto que permite analizar la base de datos de correos electrónicos de la empresa Enron, que fue publicada tras el escándalo financiero que la involucró. El proyecto consta de tres partes principales: un indexador, un servidor y una interfaz web.

El indexador es un programa que lee los archivos de los correos electrónicos y los ingresa en zincsearch, una base de datos . El indexador utiliza golang para leer los archivos y enviar los datos relevantes de los correos, como el remitente, el destinatario, el asunto, el contenido, y la fecha.

El servidor es una aplicación web que se comunica con zincsearch y retorna la información solicitada por la interfaz web. El servidor utiliza el framework chi para crear una API REST que recibe las consultas de la interfaz web y las envía a elasticsearch, luego procesa los resultados y los devuelve en formato JSON.

La interfaz web es una aplicación que visualiza la información de los correos electrónicos y permite hacer búsquedas generales. La interfaz web está hecha con Vue.js, un framework de JavaScript para crear interfaces de usuario reactivas. La interfaz web muestra una tabla con los correos electrónicos y sus campos.
