const express = require('express');
const mysql = require('mysql');
const bodyParser = require('body-parser');
const cors = require('cors');


const app = express();
const port = 3001;
app.use(cors());

// Configurar body-parser para que pueda leer JSON
app.use(bodyParser.json());

// Establecer conexión con la base de datos
const connection = mysql.createConnection({
    host: 'localhost',
    user: 'root',
    password: '16606016',
    database: 'db_chat'
});

// Agrega el encabezado 'Access-Control-Allow-Origin' a todas las respuestas de la API
app.use(function (req, res, next) {
    res.header('Access-Control-Allow-Origin', '*');
    res.header('Access-Control-Allow-Headers', 'Origin, X-Requested-With, Content-Type, Accept');
    next();
});

connection.connect((err) => {
    if (err) {
        console.error('Error al conectar con la base de datos:', err);
    } else {
        console.log('Conexión exitosa con la base de datos.');
    }
});



// Obtener todos los mensajes
app.get('/mensajes', (req, res) => {
    connection.query('SELECT * FROM mensajes', (error, results, fields) => {
        if (error) {
            console.error('Error al obtener mensajes:', error);
            res.status(500).send('Error al obtener mensajes.');
        } else {
            console.log('Mensajes obtenidos correctamente.');
            res.send(results);
        }
    });
});

// Obtener el número de mensajes existentes
app.get('/mensajes/count', (req, res) => {
    connection.query('SELECT COUNT(*) AS totalMensajes FROM mensajes', (error, results, fields) => {
        if (error) {
            console.error('Error al obtener el número de mensajes:', error);
            res.status(500).send('Error al obtener el número de mensajes.');
        } else {
            const totalMensajes = parseInt(results[0].totalMensajes);
            console.log(`Número total de mensajes: ${totalMensajes}`);
            res.send({totalMensajes});
        }
    });
});

// Insertar un nuevo mensaje
app.post('/mensajes', (req, res) => {
    const nuevoMensaje = req.body;
    connection.query(`INSERT INTO mensajes (sender, content) VALUES ('${nuevoMensaje.sender}', '${nuevoMensaje.content}')`, (error, results, fields) => {
        if (error) {
            console.error('Error al insertar mensaje:', error);
            res.status(500).send('Error al insertar mensaje.');
        } else {
            console.log('Mensaje insertado correctamente.');
            res.send('Mensaje insertado correctamente.');
        }
    });
});

// Insertar un nuevo usuario
app.post('/usuarios', (req, res) => {
    const nuevoUsuario = req.body;
    connection.query(`INSERT INTO usuarios (nick, conectado, ultimo_mensaje) VALUES ('${nuevoUsuario.nick}', ${nuevoUsuario.conectado}, '${nuevoUsuario.ultimoMensaje}')`, (error, results, fields) => {
        if (error) {
            console.error('Error al insertar usuario:', error);
            res.status(500).send('Error al insertar usuario.');
        } else {
            console.log('Usuario insertado correctamente.');
            res.send('Usuario insertado correctamente.');
        }
    });
});

//Mostrar los usuaros conectados
app.get('/usuarios', (req, res) => {
    connection.query('SELECT * FROM usuarios', (error, results, fields) => {
        if (error) {
            console.error('Error al obtener usuarios:', error);
            res.status(500).send('Error al obtener usuarios.');
        } else {
            console.log('Usuarios obtenidos correctamente.');
            res.send(results);
        }
    });
});

//Mostrar un usuario por su Nick
app.get('/usuarios/:nick', (req, res) => {
    const nick = req.params.nick;
    connection.query(`SELECT * FROM usuarios WHERE nick = '${nick}'`, (error, results, fields) => {
        if (error) {
            console.error('Error al obtener usuario:', error);
            res.status(500).send('Error al obtener usuario.');
        } else {
            console.log('Usuario obtenido correctamente.');
            res.send(results);
        }
    });
});

  //Eliminar un usuario por su nick
  app.delete('/usuario/:nick', (req, res) => {
    const nick = req.params.nick;
    connection.query(`DELETE FROM usuarios WHERE nick = '${nick}'`, (error, results, fields) => {
      if (error) {
        console.error('Error al eliminar usuario:', error);
        res.status(500).send('Error al eliminar usuario.');
      } else {
        console.log('Usuario eliminado correctamente.');
        res.send('Usuario eliminado correctamente.');
      }
    });
  });
  

app.listen(port, () => {
    console.log(`La aplicación está escuchando en el puerto ${port}.`);
});
