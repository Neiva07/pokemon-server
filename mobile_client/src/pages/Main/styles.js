import {StyleSheet} from 'react-native';
export default StyleSheet.create({
  root: {
    backgroundColor: 'white',
    flex: 1,
    justifyContent: 'space-between',
    alignItems: 'center',
    paddingTop: 10,
  },
  welcomeText: {
    fontSize: 30,
    paddingTop: 15,
    color: 'white',
  },
  buttonContainer: {},
  button: {
    marginVertical: 20,
    borderColor: 'rgba(0,0,0,0.5)',
    borderStyle: 'solid',
    borderWidth: 1,
    borderRadius: 5,
    backgroundColor: '#ed8032',
    paddingHorizontal: 40,
    height: 45,
    justifyContent: 'center',
    alignItems: 'center',
    elevation: 5,
  },
  buttonText: {
    color: 'white',
    fontSize: 25,
  },
});
