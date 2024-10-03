
import React, { useState } from 'react';


const saveUser = (nick) => {
  const newUser = {
    nick: nick,
    conectado: true,
    ultimo_mensaje: new Date()
  };

  return fetch('http://localhost:3001/usuarios', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(newUser)
  })
    .then(response => response.text())
    .then(data => {
      console.log(data);
    })
    .catch(error => {
      console.error('Error al guardar usuario:', error);
    });
};

const NickWindow = ({ onNickSubmit }) => {
  const [nick, setNick] = useState('');
  const [nickExists, setNickExists] = useState(false);

  const handleSubmitClick = () => {
    if (nick.trim() !== '') {
      fetch(`http://localhost:3001/usuarios/${nick}`)
        .then(response => response.json())
        .then(data => {
          if (data.length > 0) {
            setNickExists(true);
          } else {
            saveUser(nick)
              .then(() => {
                onNickSubmit(nick);
              });
          }
        })
        .catch(error => {
          console.error('Error al verificar si el nick existe:', error);
        });
    }
  };

  const handleNickChange = (e) => {
    setNick(e.target.value);
    setNickExists(false);
  };

  return (
    <div className="nick-window">
      <h3>Chat en React</h3>
      <input
        type="text"
        placeholder="Ingresa tu nick..."
        value={nick}
        onChange={handleNickChange}
      />
      {nickExists ? (
        <div className="error-message">El nick ya existe. Por favor, elige otro.</div>
      ) : (
        <button onClick={handleSubmitClick}>Comenzar chat</button>
      )}
    </div>
  );
};

export default NickWindow;



