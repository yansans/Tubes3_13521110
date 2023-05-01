import styles from './Chat.module.css'
import person from './img/person.svg'
import logo from './img/logo.svg'
import ChatInterface from './Interface/ChatInterface';
import { useRef, useEffect } from 'react';

const Chat = ({chatMessages} : ChatInterface) => {
    const chatEndRef = useRef<HTMLDivElement>(null);
    useEffect(() => {
        chatEndRef.current?.scrollIntoView({ behavior: "smooth" });
    }, [chatMessages]);

    return (
        <div className={styles.ChatColumn}>
            {chatMessages && chatMessages.map((msg, i) => (
                <div className={i%2 === 0 ? styles.ChatRight : styles.ChatLeft} key={i}>
                    <div style={{ display: 'flex'}}>
                        <img className={styles.ChatUserIcon} src={i%2 === 0 ? person : logo} alt=""></img>
                        <p className={styles.ChatMessage}>{msg}</p>
                    </div>
                </div>
            ))}
            <div ref={chatEndRef} />
        </div>
    );
};

export default Chat;