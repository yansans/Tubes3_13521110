
interface UserMessageJSON {
    chat_id : string;
    chat_name : string;
    participants : string;
    message : string[];
}

interface UserMessage {
    status : Number;
    message : string;
    data : UserMessageJSON[];
}

export default UserMessage;