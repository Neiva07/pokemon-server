import React, {useEffect, useState} from 'react';
import axios from 'axios';
import {
  SafeAreaView,
  View,
  Text,
  TouchableOpacity,
  ImageBackground,
} from 'react-native';
import styles from './styles';
import backgroundImage from '../../../assets/images/background-mage.jpg';
import Axios from 'axios';
export default function Main({navigation}) {
  const user = navigation.getParam('user');
  console.log(user);
  const [userName, setUserName] = useState('');
  useEffect(() => {
    const GOOGLE_URL = `https://oauth2.googleapis.com/tokeninfo?id_token=${user}`;
    axios.get(GOOGLE_URL).then(res => {
      console.log(res.data.given_name);
      setUserName(res.data.given_name);
    });
  }, []);
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
        <Text style={styles.welcomeText}>Welcome, {`${userName}!`}</Text>
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
