import ChatInput from './ChatInput';
import Chat from './Chat'
import History from './History';
import styles from './Index.module.css'
import { useState } from "react";


const exampleChat : String[] = ['Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industrys standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum', 
'hihi', 'huhu', 'hihi', 'huhu', 'hihi', 'huhu', 'hihi', 'huhu', 'hihi', 'huhu'];

const ChatGPT = () => {
    // Chat.tsx
    const [chatMessages, setChatMessages] : [String[], any] = useState(exampleChat);
    const addMessage = (message : String) => {
        setChatMessages([...chatMessages, message]);
    };

    // History.tsx algo
    const [selectedAlgo, setSelectedAlgo] : [String, any] = useState("KMP");
    const handleAlgoChange = (event : React.ChangeEvent<HTMLInputElement>) => {
        const newVal = event.target.value
        setSelectedAlgo(event.target.value);
        addMessage("Algorithm use : " + newVal);
    };

    return (
        <div className={styles.IndexContainer}>
            <History sendAlgoChange={handleAlgoChange}/>
            <div className={styles.ChatContainer}>
                <Chat chatMessages = {chatMessages}/>
                <div className={styles.ChatInputContainer}>
                    <ChatInput sendNewMessage={addMessage}/>
                </div>
            </div>
        </div>
    );
};

export default ChatGPT;