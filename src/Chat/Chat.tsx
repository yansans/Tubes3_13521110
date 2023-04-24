import styles from './Chat.module.css'
import person from './img/person.svg'
import logo from './img/logo.svg'

const exampleChat = ['hehe', 'hihi', 'huhu'];
const Chat = () => {
    return (
        <div className={styles.ChatColumn}>
            {exampleChat.map((msg, i) => (
                <div className={i%2 == 0 ? styles.ChatRight : styles.ChatLeft} key={i}>
                    <div style={{ display: 'flex', justifyContent: 'left'}}>
                        <img className={styles.ChatUserIcon} src={i%2 == 0 ? person : logo}></img>
                        <p style={{paddingLeft:"15px"}}>{msg}</p>
                    </div>
                </div>
            ))}
        </div>
    );
};

export default Chat;