import styles from './ChatInput.module.css'
import submit from './img/submit.svg'
import InputInterface from './Interface/ChatInputInterface';
import React, { useState } from "react";

const ChatInput = ({sendNewMessage} : InputInterface) => {
    const [message, setMessage] = useState('');
    const sendMessage = () => {
        if(message.length !== 0){
            sendNewMessage(message);
            setMessage('');
        }
    };
    const enterSubmit = (event : React.KeyboardEvent<HTMLInputElement>) => {
        if (event.key === 'Enter') {
            sendMessage();
        }
    }
    
    return (
        <div className={styles.ChatInputFull}>
            <input 
                className={styles.ChatInputMessage}
                placeholder='Send a message...'
                onChange={(event) => setMessage(event.target.value)}
                value={message}
                onKeyDown={(event) => enterSubmit(event)}
            />
            <button className={styles.ChatInputButton} onClick={sendMessage}>
                <img src={submit} alt="submit" style={{zoom:"1500%"}}></img>
            </button>
        </div>
    );
};

export default ChatInput;