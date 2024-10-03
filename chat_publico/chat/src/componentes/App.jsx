import '../App.css';
import React, { useState, useEffect } from 'react';
import MessageList from './MessageList';
import MessageInput from './MessageInput';
import NickWindow from './NickWindow';

const App = () => {
  const [messages, setMessages] = useState([]);
  const [nick, setNick] = useState('');
  const [showNickInput, setShowNickInput] = useState(true);
  const [currentUser, setCurrentUser] = useState(""); // nombre del usuario actual

  useEffect(() => {
    fetch('http://localhost:3001/mensajes', {
      method: 'GET'
    })
      .then(response => response.json())
      .then(data => {
        setMessages(data);
        console.log('Mensajes cargados correctamente desde la base de datos.');
      })
      .catch(error => {
        console.error('Error al cargar mensajes desde la base de datos:', error);
      });
  }, []);

  useEffect(() => {
    // Obtener el número inicial de mensajes
    let initialMessageCount = 0;

    fetch('http://localhost:3001/mensajes/count')
      .then(response => response.json())
      .then(data => {
        initialMessageCount = parseInt(data); // Parseo a entero
        console.log(`Número inicial de mensajes: ${initialMessageCount}`);
      })
      .catch(error => {
        console.error('Error al obtener el número inicial de mensajes:', error);
      });

    // Intervalo para actualizar los mensajes
    const intervalId = setInterval(() => {
      fetch('http://localhost:3001/mensajes', {
        method: 'GET'
      })
        .then(response => response.json())
        .then(data => {
          const currentMessageCount = data.length; // Obtener el número actual de mensajes
          if (currentMessageCount !== initialMessageCount) {
            setMessages(data);
            console.log('Mensajes actualizados correctamente desde la base de datos.');
            initialMessageCount = currentMessageCount; // Actualizar el número inicial
          }
        })
        .catch(error => {
          console.error('Error al actualizar mensajes desde la base de datos:', error);
        });
    }, 3000);

    // Limpieza del intervalo al desmontar el componente
    return () => {
      clearInterval(intervalId);
    };
  }, []);

  const handleSendMessage = (newMessage) => {
    console.log("newMessage:", newMessage)
    setMessages([...messages, newMessage]);
    fetch('http://localhost:3001/mensajes', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(newMessage)
    })
      .then(response => {
        if (!response.ok) {
          throw new Error('Error al guardar el mensaje en la base de datos.');
        }
        console.log('Mensaje guardado correctamente en la base de datos.');
      })
      .catch(error => {
        console.error(error);
      });
  };

  const handleNickSubmit = (nick) => {
    setNick(nick);
    setShowNickInput(false);
    setCurrentUser(nick);
  };

  return (
    <div className="chat-container">
      {showNickInput ? (
        <NickWindow onNickSubmit={handleNickSubmit} />
      ) : (
        <>
          <MessageList messages={messages} currentUser={currentUser} />
          <MessageInput onSendMessage={handleSendMessage} nick={nick} />
        </>
      )}
    </div>

  );

};

export default App;
