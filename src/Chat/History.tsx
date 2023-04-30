import styles from './History.module.css'
import plus from './img/plus.svg'
import chat from './img/chat.svg'
import rename from './img/rename.svg'
import deleting from './img/deleting.svg'
import { useState } from "react";
import AlgoInputInterface from './Interface/AlgoInputInterface'

const exampleHistory : String[] = ['contoh text overflow kaya gini !!!! !! ! ! ! !', 'hihi', 'huhu'];
const History = ({sendAlgoChange} : AlgoInputInterface) => {
    const [selectedHistory, setSelectedHistory] : [Number, any] = useState(0);
    const handleHistoryChange = (idx : Number) => {
        setSelectedHistory(idx);
    }
    
    const [history, setHistory] : [String[], any] = useState(exampleHistory);
    const newHistory = () => {
        setHistory([...history, new Date().toLocaleDateString() + " " + new Date().toTimeString()]);
    }

    
    function editHistory (idx : Number)  {
        return idx == selectedHistory ? (
            <a>
                <button style={{background:"transparent", borderColor:"transparent"}}>
                    <img src={rename} style={{zoom:"1750%"}}></img>
                </button>
                <button style={{background:"transparent", borderColor:"transparent"}}>
                    <img src={deleting} style={{zoom:"1750%"}}></img>
                </button>
            </a>
        ) : null;
    }

    return (
        <div className={styles.HistoryContainer}>
            <div className={styles.HistoryLog}>
                <button className={styles.HistoryButton}
                    onClick={newHistory}>
                    <img src={plus} style={{zoom:"1750%"}} alt="new"></img>
                    <label className={styles.HistoryName}>New chat</label>
                </button>
                {history.map((name, i) => (
                    <button className={styles.HistoryButton} 
                        key={i}
                        style={i == selectedHistory ? {backgroundColor:"#40414F"} : {borderColor:"transparent"}}
                        onClick={() => handleHistoryChange(i)}>
                        <img src={chat} style={{zoom:"1750%"}} alt="history"></img>
                        <label className={styles.HistoryName}>{name}</label>
                        {editHistory(i)}
                    </button>
                ))}
            </div>
            <hr style={{
                color: 'black',
                height: 1,
            }}/>
            <div className={styles.RadioButton}> 
                <input 
                    type="radio"
                    name="option" 
                    value="KMP" 
                    onChange={(event) => sendAlgoChange(event)}
                    defaultChecked/>KMP
            </div>
            <div className={styles.RadioButton}>
                <input 
                    type="radio"
                    name="option" 
                    value="BM" 
                    onChange={(event) => sendAlgoChange(event)}
                    />BM
            </div>
        </div>
    );
};

export default History;