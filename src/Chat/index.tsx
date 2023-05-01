import ChatInput from './ChatInput';
import Chat from './Chat'
import History from './History';
import styles from './Index.module.css'
import { useState } from "react";

const exampleChat : string[][] = [
    ['Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industrys standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum', 'hihi', 'huhu', 'hihi', 'huhu', 'hihi', 'huhu', 'hihi', 'huhu', 'hihi', 'huhu'],
    ['aaaa', 'iiii', 'uuuu', 'uuuuu'],
    ['uweeeee']];
const exampleHistory : string[] = ['contoh text overflow kaya gini !!!! !! ! ! ! !', 'hihi', 'huhu'];

const ChatGPT = () => {
    // Chat.tsx
    const [chatMessages, setChatMessages] : [string[][], any] = useState(exampleChat);
    const addMessage = (message : string) => {
        const chatCopy = [...chatMessages]
        chatCopy[selectedHistory] = [...chatCopy[selectedHistory], message];
        setChatMessages(chatCopy);
    };

    // History.tsx history
    const [selectedHistory, setSelectedHistory] : [number, any] = useState(0);
    const handleHistoryChange = (idx : number) => {
        setSelectedHistory(idx);
    }
    
    const [history, setHistory] : [string[], any] = useState(exampleHistory);
    const newHistory = () => {
        setHistory([new Date().toLocaleDateString() + " " + new Date().toTimeString(), ...history]);
        setChatMessages([[], ...chatMessages]);
        setSelectedHistory(0);
    }

    const deleteHistory = () => {
        setHistory([...history.slice(0, selectedHistory), ...history.slice(selectedHistory+1)]);
        setChatMessages([...chatMessages.slice(0, selectedHistory), ...chatMessages.slice(selectedHistory+1)]);
    }

    const renameHistory = (name : string) => {
        history[selectedHistory] = name;
    }

    // History.tsx algo
    const [selectedAlgo, setSelectedAlgo] : [string, any] = useState("KMP");
    const handleAlgoChange = (event : React.ChangeEvent<HTMLInputElement>) => {
        const newVal = event.target.value
        setSelectedAlgo(event.target.value);
        addMessage("Algorithm use : " + newVal);
    };

    // main HTML element
    return (
        <div className={styles.IndexContainer}>
            <History 
                selectedHistory={selectedHistory}
                history={history}
                sendAlgoChange={handleAlgoChange}
                handleHistoryChange={handleHistoryChange}
                newHistory={newHistory}
                deleteHistory={deleteHistory}
                renameHistory={renameHistory}
                />
            <div className={styles.ChatContainer}>
                <Chat chatMessages={chatMessages[selectedHistory]}/>
                <div className={styles.ChatInputContainer}>
                    <ChatInput sendNewMessage={addMessage}/>
                </div>
            </div>
        </div>
    );
};

export default ChatGPT;