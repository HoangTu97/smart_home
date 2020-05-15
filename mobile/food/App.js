/**
 * Sample React Native App
 * https://github.com/facebook/react-native
 *
 * @format
 * @flow strict-local
 */

import 'react-native-gesture-handler';

import React from 'react';
import { LogBox } from 'react-native';

import AppContainer from './src/navigations/AppNavigation';

export default function App() {
  return (<AppContainer />)
}

LogBox.ignoreAllLogs(false);
