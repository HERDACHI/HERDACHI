import React, { useEffect, useRef } from 'react';

const MessageList = ({ messages, currentUser }) => {
  const messageListRef = useRef(null);

  useEffect(() => {
    messageListRef.current.scrollTop = messageListRef.current.scrollHeight;
  }, [messages]);

  return (
    <div className="message-list" ref={messageListRef}>
      {messages.map(({ id, sender, content, timestamp }) => {
        //Crear nombre de la clase Mensaje dinámico.
        let messageClass = "";
        (sender == currentUser) ? messageClass = 'current-user' : messageClass = 'message';
        return (
          <div key={id} className={messageClass}>
            <span className="sender">{sender}</span>
            <p className="message-content">{content}</p>
            <span className="timestamp">{new Date(timestamp).toLocaleString('es-ES')}</span>
          </div>
        );
      })}
    </div>
  );
};

export default MessageList;







