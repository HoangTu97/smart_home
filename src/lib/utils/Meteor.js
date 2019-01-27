import {
    getWeather
} from 'lib/networking/api'

export default class Meteor {
    constructor() {
        this.q = "London,uk"
        this.a = undefined
        this.id = undefined
        this.lat = undefined
        this.lon = undefined
        this.zip = undefined
    }

    async getWeather() {
        let data = await getWeather(this);
        console.log(data);
    }
}