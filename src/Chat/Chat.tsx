import styles from './Chat.module.css'

const exampleChat = ['hehe', 'hihi', 'huhu'];
const Chat = () => {
    return (
        <div className={styles.ChatColumn}>
            {exampleChat.map((msg, i) => (
                <div className={i%2 == 0 ? styles.ChatRight : styles.ChatLeft} key={i}>
                    <div style={{ display: 'flex', justifyContent: 'left'}}>
                        <p>{msg}</p>
                    </div>
                </div>
            ))}
        </div>
    );
};

export default Chat;