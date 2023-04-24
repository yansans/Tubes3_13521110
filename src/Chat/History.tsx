import styles from './History.module.css'

const exampleHistory = ['contoh text overflow kaya gini !!!! !! ! ! ! !', 'hihi', 'huhu'];
const History = () => {
    return (
        <div className={styles.HistoryContainer}>
            <div className={styles.HistoryLog}>
                <button className={styles.HistoryName} style={{marginTop:"10px"}}>
                    <text>+ New chat</text>
                </button>
                {exampleHistory.map((msg, i) => (
                    <button className={styles.HistoryName} key={i}>
                        <text>{msg}</text>
                    </button>
                ))}
            </div>
            <hr style={{
                color: 'black',
                height: 1,
            }}/>
            <div className={styles.RadioButton}> 
                <input type="radio" name="option"/>KMP
            </div>
            <div className={styles.RadioButton}>
                <input type="radio" name="option"/>BM
            </div>
        </div>
    );
};

export default History;