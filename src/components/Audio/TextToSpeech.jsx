import React, { Component } from "react";
import { Text, View, TouchableOpacity, StyleSheet, Button } from "react-native";

// extends librairy
import Tts from 'react-native-tts';

export default class Audio extends Component {
  constructor(props) {
      super(props);
      Tts.setDefaultLanguage('en-IE');
  }

  _onPressSpeech = () => {
    Tts.speak('Hello, world!');
  }

  render() {
    return (
      <View>
        <Text style={styles.transcript}>
            Transcript
        </Text>
        <Button style={styles.transcript}
        onPress={this._onPressSpeech.bind(this)}
        title="Start"></Button>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  transcript: {
    textAlign: 'center',
    color: '#B0171F',
    marginBottom: 1,
    top: '400%',
  },
});
