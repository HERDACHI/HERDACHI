# Website Downloader

Este proyecto es una herramienta para descargar el contenido completo de un sitio web, incluyendo todos sus activos (imágenes, CSS, JavaScript, etc.). Es útil para crear copias locales de sitios web para análisis, respaldo o desarrollo offline.

## Características

- **Descarga Completa**: Descarga todo el contenido de un sitio web, incluyendo HTML, imágenes, CSS, y JavaScript.
- **Fácil de Usar**: Interfaz sencilla para especificar la URL del sitio web que deseas descargar.
- **Configuración Personalizable**: Permite ajustar parámetros como la profundidad de la descarga y la inclusión de ciertos tipos de archivos.

## Tecnologías Utilizadas

- **Python**: Lenguaje de programación principal.
- **Requests**: Biblioteca para realizar solicitudes HTTP.
- **BeautifulSoup**: Biblioteca para analizar y extraer datos de archivos HTML y XML.
- **OS**: Módulo para interactuar con el sistema operativo.

## Requisitos Previos

- **Python 3.x**: Asegúrate de tener Python instalado en tu sistema.
- **Pip**: Administrador de paquetes de Python.

## Instalación

1. Clona el repositorio:
    ```bash
    git clone https://github.com/HERDACHI/HERDACHI.git
    ```
2. Navega al directorio del proyecto:
    ```bash
    cd HERDACHI/website_downloader
    ```
3. Instala las dependencias necesarias:
    ```bash
    pip install -r requirements.txt
    ```

## Uso

1. Ejecuta el script principal:
    ```bash
    python downloader.py
    ```
2. Introduce la URL del sitio web que deseas descargar cuando se te solicite.
3. El contenido descargado se guardará en un directorio local dentro del proyecto.


## Licencia

Este proyecto está licenciado bajo la Licencia MIT. Consulta el archivo LICENSE para más detalles.

