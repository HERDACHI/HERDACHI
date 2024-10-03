
// Importa el módulo MySQL
const mysql = require('mysql');

// Importa el módulo Express
const express = require('express');

// Crea una aplicación Express
const app = express();

// Crea una conexión a la base de datos MySQL
const connection = mysql.createConnection({
  host: 'localhost',
  user: 'root',
  password: '16606016',
  database: 'db_restaurante'
});

// Agrega el encabezado 'Access-Control-Allow-Origin' a todas las respuestas de la API
app.use(function(req, res, next) {
  res.header('Access-Control-Allow-Origin', '*');
  res.header('Access-Control-Allow-Headers', 'Origin, X-Requested-With, Content-Type, Accept');
  next();
});

// Define la ruta para la consulta a la base de datos
app.get('/mesoneros', (req, res) => {
  // Realiza la consulta a la base de datos
  connection.query("SELECT nombre, cargo, direccion, telefono FROM usuarios WHERE cargo='MESONERO'", (error, results) => {
    if (error) {
      throw error;
    }

    // Devuelve los datos en formato JSON
    res.json(results);
  });
});

app.get('/mesas', (req, res) => {
  // Realiza la consulta a la base de datos
  connection.query("SELECT nombre, capacidad, ubicacion FROM mesas", (error, results) => {
    if (error) {
      throw error;
    }

    // Devuelve los datos en formato JSON
    res.json(results);
  });
});

app.get('/productos', (req, res) => {
  // Realiza la consulta a la base de datos
  connection.query("SELECT nombre, precio, nota FROM productos", (error, results) => {
    if (error) {
      throw error;
    }

    // Devuelve los datos en formato JSON
    res.json(results);
  });
});

// Maneja la solicitud para procesar el pedido
app.post('/procesar-pedido', async (req, res) => {
  try {
    const { nombre, cantidad, precio, nota, subtotal } = req.body;

    // Inserta los datos en la tabla "pedido"
    await connection.execute(
      'INSERT INTO pedido (nombre, cantidad, precio, nota, subtotal) VALUES (?, ?, ?, ?, ?)',
      [nombre, cantidad, precio, nota, subtotal]
    );

    res.status(200).json({ message: 'Pedido procesado correctamente' });
  } catch (error) {
    console.error('Error al procesar el pedido:', error);
    res.status(500).json({ error: 'Error al procesar el pedido' });
  }
});

// Inicia el servidor
app.listen(3000, () => {
  console.log('Servidor iniciado en el puerto 3000');
});
