import styles from './AlgoOption.module.css'
import AlgoState from "./Interface/AlgoState"

const AlgoOption = ({sendAlgoChange} : AlgoState) => {
    return (
        <>
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
        </>
    );
};

export default AlgoOption;
