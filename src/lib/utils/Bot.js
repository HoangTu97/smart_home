import {
    Platform
} from 'react-native'
import Voice from 'react-native-voice'
import Tts from 'react-native-tts'

import {
    getDialogFlow
} from 'lib/networking/api'

class BotAssist {
    constructor() {
        this._name = "Peter"
        this._timer = null
        this._recorded_voice = []

        Voice.onSpeechResults = this.onListening.bind(this)
        Voice.onSpeechEnd = this.endListening.bind(this)

        Tts.addEventListener('tts-finish', this.listen.bind(this))
    }

    get name() {
        return this._name
    }

    sayHi() {
        return `Hi, my name is ${this.name}`
    }

    async onListening(e) {
        this._recorded_voice = e.value
        _this = this
        if (Platform.OS === 'ios') {
            if (this._timer !== null) {
                clearTimeout(this._timer)
            }
            this._timer = setTimeout(() => {
                _this.stopListen()
            }, 2000)
        }
    }

    async endListening(e) {
        if (Platform.OS === 'ios') {
            this._timer = null;
        }
        if (this._recorded_voice != null && this._recorded_voice != "") {
            const response = await this.reply(this._recorded_voice)

            this.processResp(response)
        }
    }

    async listen() {
        await Voice.start('en-US')
    }

    async stopListen() {
        Voice.stop()
    }

    async reply(msg) {
        return await getDialogFlow(msg)
    }

    speak(msg) {
        Tts.speak(msg)
    }

    processResp(response) {
        if (response.result.metadata.intentName == "") {
            // TODO: process intent
        }
        else {
            console.log(response.result.fulfillment.speech)
            // this.speak(response.result.fulfillment.speech)
        }
    }
}

export const Bot = new BotAssist();