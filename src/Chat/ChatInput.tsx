import styles from './ChatInput.module.css'

const ChatInput = () => {
    return (
        <div>
            <input 
                className={styles.ChatInputMessage}
                placeholder='Type something'
            />
            <button className={styles.ChatInputButton}>
                Submit
            </button>
        </div>
    );
};

export default ChatInput;