
interface NewUserMessage {
    chat_id : string;
    message : string;
    sender ?: Date;
    algorithm : string;
}

export default NewUserMessage;