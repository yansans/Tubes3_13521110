import axios from 'axios';
import ChatInformation from './Interface/ChatInformation';
import UserMessage from './BackEndInterface/UserMessage';
import NewUserMessage from './BackEndInterface/NewUserMessage';
import NewChat from './BackEndInterface/NewChat';
import EditChat from './BackEndInterface/EditChat';

const apiUrl = 'http://localhost:6969';

export async function getData() : Promise<ChatInformation> {
    try {
        console.log("Get Data from API");
        const historyID : string[] = [];
        const history : string[] = [];
        const chat : string[][] = [];

        const response = await axios.get<UserMessage>(`${apiUrl}/app`);
        response.data.data.reverse();
        for(let historyInformation of response.data.data){
            historyID.push(historyInformation.chat_id);
            history.push(historyInformation.chat_name);
            chat.push(historyInformation.message);  
        }
        
        const chatInformation : ChatInformation = {
            historyID : historyID,
            history : history,
            chat : chat
        };
        return chatInformation;
    } catch (err) {
        alert(err);
        throw err;
    }
}

export async function sendMessage(historyID : string, message : string, algo : string) : Promise<string> {
    try {
        const messageToSend : NewUserMessage = {
            chat_id : historyID,
            message : message,
            algorithm : algo
        };
        const response = await axios.post(`${apiUrl}/app/chat`, messageToSend);
        const responseString : string = response.data.data;
        return responseString;
    } catch (err) {
        alert(err);
        throw err;
    }
}

export function deleteHistory(historyID : string) {
    try {
        axios.delete(`${apiUrl}/app/${historyID}`);
    } catch (err) {
        alert(err);
        throw err;
    }
}

export function renameHistory(historyID : string, newName : string) {
    try {
        const historyToRename : EditChat = {
            chat_id : historyID,
            chat_name : newName
        };
        axios.put(`${apiUrl}/app`, historyToRename);
    } catch (err) {
        alert(err);
        throw err;
    }
}

export async function newHistory(historyName : string) : Promise<string> {
    try {
        const historyToAdd : NewChat = {
            chat_name : historyName
        };
        const response = await axios.post(`${apiUrl}/app`, historyToAdd);
        const historyID : string = response.data.data;
        return historyID;
    } catch (err) {
        alert(err);
        throw err;
    }
}

