import React, {useState} from 'react';
import {
  SafeAreaView,
  Text,
  View,
  TouchableOpacity,
  ImageBackground,
} from 'react-native';
import CountdownCircle from 'react-native-countdown-circle';
import backgroundImage from '../../../assets/images/pokeball-background.jpg';
import styles from './styles';

function Game() {
  const [gameIsOn, setGameIsOn] = useState(false);
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
        <View style={{marginTop: 10}}>
          <CountdownCircle
            style={{marginTop: 2}}
            seconds={10}
            radius={25}
            borderWidth={7}
            color="#3169c4"
            bgColor="#f7c52d"
            shadowColor="#fff"
            textStyle={{fontSize: 24, color: 'white'}}
            onTimeElapsed={() => console.log('Game Ended!')}
          />
        </View>

        <View>
          <Text>POKEMON IMAGE GOES HERE</Text>
        </View>
        <View style={styles.buttonContainer}>
          <TouchableOpacity
            onPress={() => setGameIsOn(true)}
            style={[{backgroundColor: 'rgb(202,65,34)'}, styles.button]}>
            <Text style={styles.buttonText}>Pokemon 1</Text>
          </TouchableOpacity>
          <TouchableOpacity
            onPress={() => setGameIsOn(true)}
            style={[{backgroundColor: 'rgb(123,63,129)'}, styles.button]}>
            <Text style={styles.buttonText}>Pokemon 2</Text>
          </TouchableOpacity>
          <TouchableOpacity
            onPress={() => setGameIsOn(true)}
            style={[{backgroundColor: 'rgb(130,230,250)'}, styles.button]}>
            <Text style={styles.buttonText}>Pokemon 3</Text>
          </TouchableOpacity>
          <TouchableOpacity
            onPress={() => setGameIsOn(true)}
            style={[{backgroundColor: 'rgb(100,163,64)'}, styles.button]}>
            <Text style={styles.buttonText}>Pokemon 4</Text>
          </TouchableOpacity>
        </View>
      </ImageBackground>
    </SafeAreaView>
  );
}
export default Game;
