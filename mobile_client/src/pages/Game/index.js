import React, {useState} from 'react';
import {
  SafeAreaView,
  Text,
  View,
  TouchableOpacity,
  ImageBackground,
  Image,
} from 'react-native';
import LinearGradient from 'react-native-linear-gradient';
import {Brightness} from 'react-native-color-matrix-image-filters';
import CountdownCircle from 'react-native-countdown-circle';
import backgroundImage from '../../../assets/images/pokeball-background.jpg';
import PikachuImage from '../../../assets/images/Pikachu.png';
import styles from './styles';

function Game() {
  const [gameIsOn, setGameIsOn] = useState(false);
  return (
    <SafeAreaView style={{flex: 1}}>
      <LinearGradient
        style={styles.root}
        colors={['#e2a253', '#f3d361']}
        useAngle
        angle={135}
        angleCenter={{x: 0.5, y: 0.5}}>
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
          <Brightness amount={0}>
            <Image style={styles.pokemonImage} source={PikachuImage} />
          </Brightness>
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
      </LinearGradient>
    </SafeAreaView>
  );
}
export default Game;
