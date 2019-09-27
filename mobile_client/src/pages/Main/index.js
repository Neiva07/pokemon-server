import React from 'react';
import {SafeAreaView, Text, StyleSheet} from 'react-native';
import styles from './styles';

export default function Main({navigation}) {
  const user = navigation.getParam('user');
  console.log(user);
  return (
    <SafeAreaView style={styles.root}>
      <Text>Welcome!</Text>
    </SafeAreaView>
  );
}
