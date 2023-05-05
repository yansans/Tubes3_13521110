import ChatInput from './ChatInput';
import Chat from './Chat'
import History from './History';
import styles from './Index.module.css'
import { useState, useEffect } from "react";
import * as API from './API';

const ChatGPT = () => {
    // Chat.tsx
    const [chatMessages, setChatMessages] : [string[][], any] = useState([]);
    const [chatMutex, setChatMutex] : [boolean, any] = useState(false);
    const addMessage = (message : string) => {
        if(history.length <= selectedHistory) {
            alert("No chat selected, please choose or make chat before sending a message");
            return;
        }

        if(chatMutex) {
            alert("Please wait for bot to respond before send new message");
            return;
        }

        setChatMutex(true);
        const chatCopy = [...chatMessages];
        chatCopy[selectedHistory] = [...chatCopy[selectedHistory], message];
        setChatMessages(chatCopy);
        
        API.sendMessage(historyID[selectedHistory], message, selectedAlgo + selectedExact).then((botResponse) => {
            const chatCopy2 = [...chatCopy];
            chatCopy2[selectedHistory] = [...chatCopy2[selectedHistory], botResponse];
            setChatMessages(chatCopy2);
            setChatMutex(false);
        });
    };

    // History.tsx history
    const [selectedHistory, setSelectedHistory] : [number, any] = useState(0);
    const handleHistoryChange = (idx : number) => {
        setSelectedHistory(idx);
    }
    
    const [historyID, setHistoryID] : [string[], any] = useState([]);
    const [history, setHistory] : [string[], any] = useState([]);
    const newHistory = () => {
        const name : string = new Date().toLocaleDateString() + " " + new Date().toTimeString();

        setHistory([name, ...history]);
        setChatMessages([[], ...chatMessages]);

        API.newHistory(name).then((newHistoryID) => {
            setHistoryID([newHistoryID, ...historyID]);
        })

        setSelectedHistory(0);
    }

    const deleteHistory = () => {
        setHistoryID([...historyID.slice(0, selectedHistory), ...historyID.slice(selectedHistory+1)]);
        setHistory([...history.slice(0, selectedHistory), ...history.slice(selectedHistory+1)]);
        setChatMessages([...chatMessages.slice(0, selectedHistory), ...chatMessages.slice(selectedHistory+1)]);
        API.deleteHistory(historyID[selectedHistory]);
    }

    const renameHistory = (name : string) => {
        history[selectedHistory] = name;
        API.renameHistory(historyID[selectedHistory], name);
    }

    // History.tsx algo
    const [selectedAlgo, setSelectedAlgo] : [string, any] = useState("KMP");
    const [selectedExact, setSelectedExact] : [string, any] = useState("");
    const handleAlgoOptionChange = (event : React.ChangeEvent<HTMLInputElement>) => {
        setSelectedAlgo(event.target.value);
    };
    const handleAlgoExactChange = (isExact : boolean) => {
        setSelectedExact(isExact ? "EXACT" : "");
    };

    // API
    useEffect(() => {
        getChat();
    }, []);

    const getChat = () => {
        API.getData().then((ChatInformation) => {
            setHistoryID(ChatInformation.historyID); 
            setHistory(ChatInformation.history); 
            setChatMessages(ChatInformation.chat); 
        })
    }

    // main HTML element
    return (
        <div className={styles.IndexContainer}>
            <History 
                selectedHistory={selectedHistory}
                history={history}
                sendAlgoOptionChange={handleAlgoOptionChange}
                sendAlgoExactChange={handleAlgoExactChange}
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