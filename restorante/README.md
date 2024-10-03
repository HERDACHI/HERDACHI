# Restorante

Este proyecto es una aplicación para la gestión de un restaurante. Permite a un mesonero manejar reservas, menús y pedidos del restaurante.

## Características

- **Gestión de Reservas**: Crear, actualizar y eliminar reservas.
- **Gestión de Menús**: Añadir, actualizar y eliminar elementos del menú.
- **Gestión de Pedidos**: Tomar, actualizar y gestionar pedidos de los clientes.
- **Administración del Restaurante**: Herramientas para la gestión de mesas, horarios y personal.

## Tecnologías Utilizadas

- **Node.js**: Entorno de ejecución para JavaScript en el servidor.
- **Express**: Framework web para Node.js.
- **MySQL**: Base de datos relacional.
- **Sequelize**: ORM para MySQL.
- **HTML/CSS/JavaScript**: Para la interfaz de usuario.

## Requisitos Previos

- **Node.js**: Asegúrate de tener Node.js instalado en tu sistema.
- **MySQL**: Asegúrate de tener MySQL instalado y en ejecución.

## Instalación

1. Clona el repositorio:
    ```bash
    git clone https://github.com/HERDACHI/HERDACHI.git
    ```
2. Navega al directorio del proyecto:
    ```bash
    cd HERDACHI/restorante/servidor
    ```
3. Instala las dependencias necesarias:
    ```bash
    npm install
    ```
4. Configura la base de datos en `config/config.json`:
    ```json
    {
      "development": {
        "username": "tu_usuario",
        "password": "tu_contraseña",
        "database": "nombre_base_de_datos",
        "host": "127.0.0.1",
        "dialect": "mysql"
      }
    }
    ```

## Uso

1. Asegúrate de que MySQL esté en ejecución.
2. Ejecuta las migraciones para crear las tablas en la base de datos:
    ```bash
    npx sequelize-cli db:migrate
    ```
3. Ejecuta el servidor:
    ```bash
    npm start
    ```
4. Abre tu navegador y ve a `http://localhost:3000` para acceder a la aplicación.

## Endpoints Principales

### Reservas

- **GET /reservas**: Obtener todas las reservas.
- **POST /reservas**: Crear una nueva reserva.
- **GET /reservas/{id}**: Obtener una reserva por ID.
- **PUT /reservas/{id}**: Actualizar una reserva por ID.
- **DELETE /reservas/{id}**: Eliminar una reserva por ID.

### Menús

- **GET /menus**: Obtener todos los elementos del menú.
- **POST /menus**: Añadir un nuevo elemento al menú.
- **GET /menus/{id}**: Obtener un elemento del menú por ID.
- **PUT /menus/{id}**: Actualizar un elemento del menú por ID.
- **DELETE /menus/{id}**: Eliminar un elemento del menú por ID.

### Pedidos

- **GET /pedidos**: Obtener todos los pedidos.
- **POST /pedidos**: Crear un nuevo pedido.
- **GET /pedidos/{id}**: Obtener un pedido por ID.
- **PUT /pedidos/{id}**: Actualizar un pedido por ID.
- **DELETE /pedidos/{id}**: Eliminar un pedido por ID.

## Licencia

Este proyecto está licenciado bajo la Licencia MIT. Consulta el archivo LICENSE para más detalles.


