import axios from 'axios'
import google_key from 'config/google-api.json'
import weather_key from 'config/openWeatherMap-api.json'


export async function getDialogFlow(msg) {
    const ACCESS_TOKEN = google_key.dialogflow

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
        return res.data
    })
}

export async function getWeather(data) {
    const WEATHER_KEY = weather_key.api_key
    
    let query = Object.keys(data)
    .filter((key) => { return data[key] !== undefined && key[0] !== '_' })
    .map((key) => {
        return encodeURIComponent(key) + '=' + encodeURIComponent(data[key])
    }).join('&')

    console.log(query)
    
    return await axios.get(`http://api.openweathermap.org/data/2.5/weather?appid=${WEATHER_KEY}&${query}`)
    .then((res) => {
        return res.data
    })
}