import styles from './ChatInput.module.css'
import submit from './img/submit.svg'

const ChatInput = () => {
    return (
        <div className={styles.ChatInputFull}>
            <input 
                className={styles.ChatInputMessage}
                placeholder='Send a message...'
            />
            <button className={styles.ChatInputButton}>
                <img src={submit}></img>
            </button>
        </div>
    );
};

export default ChatInput;