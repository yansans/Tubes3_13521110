
interface HistoryState {
    selectedHistory : number;
    history : string[];
    sendAlgoOptionChange : (event : React.ChangeEvent<HTMLInputElement>) => void;
    sendAlgoExactChange : (isExact : boolean) => void;
    handleHistoryChange : (idx : number) => void;
    newHistory : () => void;
    deleteHistory : () => void;
    renameHistory : (name : string) => void;
}

export default HistoryState;