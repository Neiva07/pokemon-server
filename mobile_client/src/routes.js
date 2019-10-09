import React from 'react';
import {TouchableOpacity, Image} from 'react-native';
import {createAppContainer, createSwitchNavigator} from 'react-navigation';
import {createBottomTabNavigator} from 'react-navigation-tabs';
import Icon from 'react-native-vector-icons/FontAwesome';
import Login from './pages/Login';
import Main from './pages/Main';
import FriendList from './pages/FriendList';
import Ranking from './pages/Ranking';
import Profile from './pages/Profile';
import Game from './pages/Game';
import pokeballIcon from '../assets/images/pokeball64px.png';
import pokedexIcon from '../assets/images/pokedex64px.png';
import crownIcon from '../assets/images/crown64px.png';
import playerIcon from '../assets/images/player64px.png';

const tabNavigator = createBottomTabNavigator(
  {
    Home: {
      screen: Main,
      navigationOptions: {
        tabBarLabel: 'Play',
        tabBarIcon: (
          <Image style={{width: 28, height: 28}} source={pokeballIcon} />
        ),
      },
    },
    FriendList: {
      screen: FriendList,
      navigationOptions: {
        tabBarLabel: 'Friend List',
        tabBarIcon: (
          <Image style={{width: 30, height: 30}} source={playerIcon} />
        ),
      },
    },
    Ranking: {
      screen: Game,
      navigationOptions: {
        tabBarLabel: 'Ranking',
        tabBarIcon: (
          <Image style={{width: 30, height: 30}} source={crownIcon} />
        ),
      },
    },
    Profile: {
      screen: Profile,
      navigationOptions: {
        tabBarLabel: 'Profile',
        tabBarIcon: (
          <Image style={{width: 30, height: 30}} source={pokedexIcon} />
        ),
      },
    },
  },
  {
    tabBarOptions: {
      activeBackgroundColor: '#477dcc',
      activeTintColor: 'white',
      inactiveTintColor: 'black',
      inactiveBackgroundColor: '#dede54',
    },
  },
);
const loginNavigator = createSwitchNavigator({Login, tabNavigator});
export default createAppContainer(loginNavigator);
