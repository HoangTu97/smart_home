import {Platform} from 'react-native';
import Voice from 'react-native-voice';
import Tts from 'react-native-tts';

import {getDialogFlow} from 'lib/networking/api';
import Meteor from './Meteor';

class BotAssist {
  constructor() {
    this._name = 'Lucas';
    this._timer = null;
    this._wait_time = 2000;
    this._lang = 'en-US';
    this._recorded_voice = [];

    Voice.onSpeechResults = this.onListening.bind(this);
    Voice.onSpeechEnd = this.endListening.bind(this);

    Tts.addEventListener('tts-finish', this.listen.bind(this));
  }

  get name() {
    return this._name;
  }

  sayHi() {
    return `Hi, my name is ${this.name}`;
  }

  async onListening(e) {
    this._recorded_voice = e.value;
    if (Platform.OS === 'ios') {
      if (this._timer !== null) {
        clearTimeout(this._timer);
      }
      this._timer = setTimeout(() => {
        this.stopListen();
      }, this._wait_time);
    }
  }

  async endListening(e) {
    if (Platform.OS === 'ios') {
      this._timer = null;
    }
    if (this._recorded_voice !== null && this._recorded_voice !== '') {
      console.log('begin');
      const response = await this.reply(this._recorded_voice);

      this.processResp(response);
    }
  }

  async listen() {
    await Voice.start(this._lang);
  }

  async stopListen() {
    Voice.stop();
  }

  async reply(msg) {
    return await getDialogFlow(msg);
  }

  speak(msg) {
    Tts.speak(msg);
  }

  processResp(response) {
    if (response.result.metadata.intentName === 'Weather') {
      // TODO: process intent
      let meteor = new Meteor();
      meteor.getWeather();
    } else {
      this.speak(response.result.fulfillment.speech);
    }
  }

  test() {
    let meteor = new Meteor();
    meteor.getMyWeather();
  }
}

export const Bot = new BotAssist();
