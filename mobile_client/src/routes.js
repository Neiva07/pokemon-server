import React from 'react';
import {createAppContainer, createSwitchNavigator} from 'react-navigation';
import {createBottomTabNavigator} from 'react-navigation-tabs';
import Icon from 'react-native-vector-icons/FontAwesome';
import Login from './pages/Login';
import Main from './pages/Main';
import FriendList from './pages/FriendList';
import Ranking from './pages/Ranking';
import Profile from './pages/Profile';
const tabNavigator = createBottomTabNavigator({
  Home: {
    screen: Main,
    navigationOptions: {
      tabBarLabel: 'Play',
      tabBarIcon: props => <Icon name="play" size={28} color="#7a7774" />,
    },
  },
  FriendList: {
    screen: FriendList,
    navigationOptions: {
      tabBarLabel: 'Friend List',
      tabBarIcon: <Icon name="users" size={28} color="#7a7774" />,
    },
  },
  Ranking: {
    screen: Ranking,
    navigationOptions: {
      tabBarLabel: 'Ranking',
      tabBarIcon: <Icon name="trophy" size={28} color="#7a7774" />,
    },
  },
  Profile: {
    screen: Profile,
    navigationOptions: {
      tabBarLabel: 'Profile',
      tabBarIcon: <Icon name="home" size={28} color="#7a7774" />,
    },
  },
});
const loginNavigator = createSwitchNavigator({Login, tabNavigator});
export default createAppContainer(loginNavigator);
