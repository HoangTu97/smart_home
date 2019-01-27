import React, {Component} from 'react';
import {Text, View, TouchableOpacity} from 'react-native';

import Audio from 'components/Audio/Audio.jsx'

export default class App extends Component {
  render() {
    return (
      <View style={{marginTop: 50}}>
        <Audio/>
      </View>
    );
  }
}
