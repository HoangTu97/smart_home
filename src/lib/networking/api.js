import axios from 'axios';
import config from 'config/google-api.json';

export async function getDialogFlow(msg) {
    const ACCESS_TOKEN = config.dialogflow;

    return await axios.post(`https://api.dialogflow.com/v1/query?v=20170712`,
        JSON.stringify({
            query: msg,
            lang: 'EN',
            sessionId: 'somerandomthing'
        })
    , {
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json; charset=utf-8',
            'Authorization': `Bearer ${ACCESS_TOKEN}`,
        }
    }).then((res) => {
        return res.data;
    })
}