import ChatInput from './ChatInput';
import ChatLog from './Chat'
import History from './History';
import styles from './Index.module.css'


const Chat = () => {
    return (
        <div className={styles.IndexContainer}>
            <History/>
            <div className={styles.ChatContainer}>
                <ChatLog/>
                <div className={styles.ChatInputContainer}>
                    <ChatInput/>
                </div>
            </div>
        </div>
    );
};

export default Chat;