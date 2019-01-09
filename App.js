/**
 * Sample React Native App
 * https://github.com/facebook/react-native
 *
 * @format
 * @flow
 */

import React, {Component} from 'react';
import {Platform, StyleSheet, Text, View, TouchableOpacity} from 'react-native';
import { BleManager, LogLevel } from 'react-native-ble-plx';


export default class App extends Component {
  constructor() {
    super();
  }

  render() {
    return (
      <View style={{marginTop: 50}}>
        <TouchableOpacity onPress={() => {}}>
          <Text>Scan</Text>
        </TouchableOpacity>
        <Text>{this.state.info}</Text>
      </View>
    );
  }
}
