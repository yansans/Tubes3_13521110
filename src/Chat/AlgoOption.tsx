import styles from './AlgoOption.module.css'
import AlgoState from "./Interface/AlgoState"
import { useState } from "react";

const AlgoOption = ({sendAlgoOptionChange, sendAlgoExactChange} : AlgoState) => {
    const [exact, setExact] : [boolean, any] = useState(false);
    const handleExactChange = () => {
        sendAlgoExactChange(!exact);
        setExact(!exact);
    }
    return (
        <>
            <div className={styles.RadioButton}> 
            <input 
                type="radio"
                name="option" 
                value="KMP" 
                onChange={(event) => sendAlgoOptionChange(event)}
                defaultChecked/>KMP
            </div>
            <div className={styles.RadioButton}>
            <input 
                type="radio"
                name="option" 
                value="BM" 
                onChange={(event) => sendAlgoOptionChange(event)}
                />BM
            </div>
            <div className={styles.RadioButton}>
                <input 
                    type="checkbox"
                    checked={exact}
                    onChange={handleExactChange}/>
                <label/>Exact
            </div>
        </>
    );
};

export default AlgoOption;
