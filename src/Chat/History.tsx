import styles from './History.module.css'
import plus from './img/plus.svg'
import chat from './img/chat.svg'
import rename from './img/rename.svg'
import deleting from './img/deleting.svg'
import check from './img/check.svg'
import HistoryState from './Interface/HistoryState'
import AlgoOption from './AlgoOption'
import { useState, useRef } from 'react'

const History = ({selectedHistory,
                    history,
                    sendAlgoOptionChange,
                    sendAlgoExactChange,
                    handleHistoryChange,
                    newHistory,
                    renameHistory,
                    deleteHistory
                } : HistoryState) => {
    const renameRef = useRef<HTMLInputElement>(null);
    const [renaming, setRenaming] : [boolean, any] = useState(false);
    const [renameName, setRenameName] : [string, any] = useState("");

    const enableRenaming = (name : string) => {
        setRenameName(name);
        setRenaming(true);
    }

    const renameSubmit = () => {
        renameHistory(renameName);
        setRenaming(false);
    };

    const enterSubmit = (event : React.KeyboardEvent<HTMLInputElement>) => {
        if (event.key === 'Enter') {
            renameSubmit();
        }
    }

    const historyName = (name : string, idx : number) => {
        return (
            renaming && idx === selectedHistory ? (
                <input className={styles.HistoryName}
                    type="text"
                    value={renameName}
                    ref={renameRef}
                    onChange={(event) => setRenameName(event.target.value)}
                    onFocus={() => console.log("test")}
                    onBlur={renameSubmit}
                    onKeyDown={(event) => enterSubmit(event)} 
                    />
            ) : (
                <label className={styles.HistoryName}>{name}</label>
            )
        )
    }
    
    return (
        <div className={styles.HistoryContainer}>
            <div className={styles.HistoryLog}>
                <button className={styles.HistoryButton}
                    onClick={newHistory}>
                    <img src={plus} style={{zoom:"1750%"}} alt="new"></img>
                    <label className={styles.HistoryName}>New chat</label>
                </button>
                {history.map((name, idx) => (
                    <button className={styles.HistoryButton} 
                        key={idx}
                        style={idx === selectedHistory ? {backgroundColor:"#40414F"} : {borderColor:"transparent"}}
                        onClick={() => handleHistoryChange(idx)}>
                        <img src={chat} style={{zoom:"1750%"}} alt="history"></img>
                        {historyName(name, idx)}
                        {(idx === selectedHistory) ? 
                        <>
                            <button className={styles.EditButton} onClick={deleteHistory}>
                                <img src={deleting} style={{zoom:"1750%"}} alt="delete"/>
                            </button>
                            {renaming ? 
                                <button className={styles.EditButton} onClick={renameSubmit}>
                                    <img src={check} style={{zoom:"7%"}} alt="renamesubmit"/>
                                </button> : 
                                <button className={styles.EditButton} onClick={() => enableRenaming(name)}>
                                    <img src={rename} style={{zoom:"1750%"}} alt="rename"/>
                                </button>
                            }
                                
                        </> : ""}
                    </button>
                ))}
            </div>
            <hr style={{
                color: 'black',
                height: 1,
            }}/>
            <AlgoOption
                sendAlgoOptionChange={sendAlgoOptionChange}
                sendAlgoExactChange={sendAlgoExactChange}
                />
        </div>
    );
};

export default History;