import React, { Component } from "react";
import { Text, View, StyleSheet, Button, Platform } from "react-native";

import { Bot } from 'lib/utils/Bot';


export default class Audio extends Component {
  constructor(props) {
    super(props);
    this.timer = null;
    this.state = { 
        showText: null, 
        results: []
    };

    Bot.speak(Bot.sayHi());
  }

  render() {
    return (
      <View>
        <Text style={styles.transcript}>
            Transcript
        </Text>
        {this.state.results.map((result, index) => 
            <Text style={styles.transcript}>{result}</Text>
        )}
        <Button style={styles.transcript}
        onPress={Bot.listen}
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
