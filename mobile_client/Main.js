import React from 'react';
import {SafeAreaView, Text, StyleSheet} from 'react-native';
const styles = StyleSheet.create({
  root: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
});
export default function Main({navigation}) {
  const user = navigation.getParam('user');
  console.log(user);
  return (
    <SafeAreaView style={styles.root}>
      <Text>Welcome!</Text>
    </SafeAreaView>
  );
}
