
import React, { useState } from 'react';

const MessageInput = ({ onSendMessage, nick }) => {
  const [message, setMessage] = useState('');

  const handleInputChange = (e) => {
    setMessage(e.target.value);
  };

  const handleKeyPress = (event) => {
    if (event.key === 'Enter') {
      handleSubmitClick();
    }
  };

  const handleSubmitClick = () => {
    if (message.trim() !== '') {
      const now = new Date();
      const options = { year: 'numeric', month: 'numeric', day: 'numeric', hour: 'numeric', minute: 'numeric', second: 'numeric' };
      const newMessage = {
        id: Date.now(),
        sender: nick,
        content: message,
        timestamp: now.toLocaleString('es-ES', options),
      };
      onSendMessage(newMessage);
      setMessage('');
    }
  };

  const handleLogoutClick = () => {
    // Hacer una petición HTTP para borrar el nick del usuario de la base de datos
    fetch(`http://localhost:3001/usuario/${nick}`, {
      method: 'DELETE',
    })
      .then(() => {
        // Redirigir al usuario a la página de NickInput
        window.location.href = '/nickinput';
      })
      .catch((error) => {
        console.error(error);
      });
  };

  return (
    <div className="message-input">
      <input
        type="text"
        placeholder="Escribe un mensaje..."
        value={message}
        onChange={handleInputChange}
        onKeyPress={handleKeyPress}
      />
      <button onClick={handleSubmitClick}>Enviar</button>
      <button onClick={handleLogoutClick}>Cerrar sesión</button>
    </div>
  );
};

export default MessageInput;


