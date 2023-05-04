
interface NewUserMessage {
    chat_id : string;
    message : string;
    sender ?: string;
    algorithm : string;
}

export default NewUserMessage;