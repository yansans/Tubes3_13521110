import styles from './Chat.module.css'
import person from './img/person.svg'
import logo from './img/logo.svg'

const exampleChat = ['Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industrys standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum', 
'hihi', 'huhu', 'hihi', 'huhu', 'hihi', 'huhu', 'hihi', 'huhu', 'hihi', 'huhu'];
const Chat = () => {
    return (
        <div className={styles.ChatColumn}>
            {exampleChat.map((msg, i) => (
                <div className={i%2 == 0 ? styles.ChatRight : styles.ChatLeft} key={i}>
                    <div style={{ display: 'flex'}}>
                        <img className={styles.ChatUserIcon} src={i%2 == 0 ? person : logo}></img>
                        <p className={styles.ChatMessage}>{msg}</p>
                    </div>
                </div>
            ))}
        </div>
    );
};

export default Chat;