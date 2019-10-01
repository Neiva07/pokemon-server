import React, {useEffect} from 'react';
import {TouchableOpacity, Text, SafeAreaView} from 'react-native';
import {
  GoogleSignin,
  GoogleSigninButton,
  statusCodes,
} from 'react-native-google-signin';
import AsyncStorage from '@react-native-community/async-storage';
import styles from './styles';

export default function Login({navigation}) {
  useEffect(() => {
    GoogleSignin.configure({
      webClientId:
        '467952177460-bulsjjsifel2jrfsuad97qcrjni44e1b.apps.googleusercontent.com', // client ID of type WEB for your server (needed to verify user ID and offline access)
    });
  }, []);

  signIn = async () => {
    try {
      await GoogleSignin.hasPlayServices();
      const userInfo = await GoogleSignin.signIn();
      const {idToken} = userInfo;
      console.log('------ID TOKEN------');
      console.log(idToken);

      await AsyncStorage.setItem('user', idToken);

      navigation.navigate('Main', {user: idToken});
    } catch (error) {
      console.error(error);
      if (error.code === statusCodes.SIGN_IN_CANCELLED) {
        // user cancelled the login flow
      } else if (error.code === statusCodes.IN_PROGRESS) {
        // operation (f.e. sign in) is in progress already
      } else if (error.code === statusCodes.PLAY_SERVICES_NOT_AVAILABLE) {
        // play services not available or outdated
      } else {
        // some other error happened
      }
    }
  };
  signOut = async () => {
    try {
      await GoogleSignin.revokeAccess();
      await GoogleSignin.signOut();
      // this.setState({user: null}); // Remember to remove the user from your app's state as well
    } catch (error) {
      console.error(error);
    }
  };
  return (
    <SafeAreaView style={styles.root}>
      <GoogleSigninButton
        style={{width: 192, height: 48}}
        size={GoogleSigninButton.Size.Wide}
        color={GoogleSigninButton.Color.Dark}
        onPress={this.signIn}
        disabled={false}
      />
      <TouchableOpacity onPress={() => navigation.navigate('Home')}>
        <Text>GO TO MAIN PAGE</Text>
      </TouchableOpacity>
    </SafeAreaView>
  );
}
