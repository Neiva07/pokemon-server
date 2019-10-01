import React from 'react';
import {
  SafeAreaView,
  View,
  Text,
  TouchableOpacity,
  ImageBackground,
} from 'react-native';
import styles from './styles';
import backgroundImage from '../../../assets/images/background-mage.jpg';
export default function Main({navigation}) {
  const user = navigation.getParam('user');
  console.log(user);
  return (
    <SafeAreaView>
      <ImageBackground
        source={backgroundImage}
        style={{
          width: '100%',
          height: '100%',
          alignItems: 'center',
          justifyContent: 'space-between',
        }}>
        <Text style={styles.welcomeText}>Welcome, Matheus!</Text>
        <View style={styles.buttonContainer}>
          <TouchableOpacity
            style={styles.button}
            onPress={() => navigation.navigate('Game')}>
            <Text style={styles.buttonText}>Play with a Friend</Text>
          </TouchableOpacity>
          <TouchableOpacity
            style={styles.button}
            onPress={() => navigation.navigate('Game')}>
            <Text style={styles.buttonText}>Find Random Trainer</Text>
          </TouchableOpacity>
        </View>
        <View></View>
      </ImageBackground>
    </SafeAreaView>
  );
}
