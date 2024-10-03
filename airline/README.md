# AIRLINE
# AIRLINE API

Este proyecto es una API para la gestión de una aerolínea ficticia llamada Andes Airline. La API permite realizar operaciones relacionadas con vuelos, pasajeros y reservas.

## Características

- **Gestión de Vuelos**: Crear, actualizar, eliminar y consultar vuelos.
- **Gestión de Pasajeros**: Registrar, actualizar y consultar información de pasajeros.
- **Gestión de Reservas**: Crear, actualizar y consultar reservas de vuelos.

## Tecnologías Utilizadas

- **Java 11**: Lenguaje de programación principal.
- **Spring Boot**: Framework para la creación de aplicaciones web.
- **Hibernate**: Framework de mapeo objeto-relacional (ORM).
- **MySQL**: Base de datos relacional.
- **Maven**: Herramienta de gestión de dependencias y construcción de proyectos.

## Requisitos Previos

- **Java 11** o superior.
- **Maven**.
- **MySQL**.

## Instalación

1. Clonar el repositorio:
    ```bash
    git clone https://github.com/HERDACHI/AIRLINE.git
    ```
2. Navegar al directorio del proyecto:
    ```bash
    cd AIRLINE
    ```
3. Configurar la base de datos en `src/main/resources/application.properties`:
    ```properties
    spring.datasource.url=jdbc:mysql://localhost:3306/airline
    spring.datasource.username=tu_usuario
    spring.datasource.password=tu_contraseña
    ```
4. Construir el proyecto con Maven:
    ```bash
    mvn clean install
    ```
5. Ejecutar la aplicación:
    ```bash
    mvn spring-boot:run
    ```

## Endpoints Principales

### Vuelos

- **GET /vuelos**: Obtener todos los vuelos.
- **POST /vuelos**: Crear un nuevo vuelo.
- **GET /vuelos/{id}**: Obtener un vuelo por ID.
- **PUT /vuelos/{id}**: Actualizar un vuelo por ID.
- **DELETE /vuelos/{id}**: Eliminar un vuelo por ID.

### Pasajeros

- **GET /pasajeros**: Obtener todos los pasajeros.
- **POST /pasajeros**: Registrar un nuevo pasajero.
- **GET /pasajeros/{id}**: Obtener un pasajero por ID.
- **PUT /pasajeros/{id}**: Actualizar un pasajero por ID.
- **DELETE /pasajeros/{id}**: Eliminar un pasajero por ID.

### Reservas

- **GET /reservas**: Obtener todas las reservas.
- **POST /reservas**: Crear una nueva reserva.
- **GET /reservas/{id}**: Obtener una reserva por ID.
- **PUT /reservas/{id}**: Actualizar una reserva por ID.
- **DELETE /reservas/{id}**: Eliminar una reserva por ID.

## Contribuciones

¡Las contribuciones son bienvenidas! Por favor, sigue los siguientes pasos:

1. Haz un fork del repositorio.
2. Crea una nueva rama (`git checkout -b feature/nueva-funcionalidad`).
3. Realiza tus cambios y haz commit (`git commit -am 'Añadir nueva funcionalidad'`).
4. Sube tus cambios (`git push origin feature/nueva-funcionalidad`).
5. Abre un Pull Request.

## Licencia

Este proyecto está licenciado bajo la Licencia MIT. Consulta el archivo LICENSE para más detalles.

