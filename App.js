import React, {Component} from 'react';
import {Text, View, TouchableOpacity} from 'react-native';

import Audio from 'components/Audio/speech.jsx'

export default class App extends Component {
  render() {
    return (
      <View style={{marginTop: 50}}>
        <TouchableOpacity onPress={() => {}}>
          <Text>Scan</Text>
        </TouchableOpacity>
        <Text>abc</Text>
        <Audio/>
      </View>
    );
  }
}
