import {StyleSheet} from 'react-native';
export default StyleSheet.create({
  root: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'space-between',
    paddingTop: 10,
  },
  buttonContainer: {
    width: '100%',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'space-between',
  },
  button: {
    marginVertical: 20,
    borderColor: 'rgba(0,0,0,0.5)',
    borderStyle: 'solid',
    borderWidth: 1,
    borderRadius: 5,
    paddingHorizontal: 40,
    height: 45,
    justifyContent: 'center',
    alignItems: 'center',
    elevation: 5,
    width: '80%',
  },
  buttonText: {
    color: 'white',
    fontSize: 25,
  },
});
