import {StyleSheet} from 'react-native';
export default StyleSheet.create({
  root: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'space-between',
  },
  buttonContainer: {
    width: '100%',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'space-between',
  },
  button: {
    marginVertical: 10,
    borderColor: 'rgba(0,0,0,0.5)',
    borderStyle: 'solid',
    borderWidth: 1,
    borderRadius: 5,
    paddingHorizontal: 40,
    height: 50,
    justifyContent: 'center',
    alignItems: 'center',
    elevation: 5,
    width: '60%',
  },
  buttonText: {
    color: 'white',
    fontSize: 25,
  },
  pokemonImage: {
    resizeMode: 'stretch',
    width: 200,
    height: 200,
  },
});
