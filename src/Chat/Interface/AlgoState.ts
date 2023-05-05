
interface AlgoState {
    sendAlgoOptionChange : (event : React.ChangeEvent<HTMLInputElement>) => void;
    sendAlgoExactChange : (isExact : boolean) => void;
}

export default AlgoState;