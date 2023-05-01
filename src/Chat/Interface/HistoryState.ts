
interface HistoryState {
    selectedHistory : number;
    history : string[];
    sendAlgoChange : (event : React.ChangeEvent<HTMLInputElement>) => void;
    handleHistoryChange : (idx : number) => void;
    newHistory : () => void;
    deleteHistory : () => void;
    renameHistory : (name : string) => void;
}

export default HistoryState;