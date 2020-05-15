import React, {useLayoutEffect, useState, useEffect} from 'react';
import {FlatList, Text, View, Image, TouchableHighlight} from 'react-native';
import styles from './styles';
import {getIngredientName, getAllIngredients} from '../../data/MockDataAPI';
import axios from 'axios';

export default function IngredientsDetailsScreen({route, navigation}) {
  const [ingredients, setIngredients] = useState([]);

  onPressIngredient = (item) => {
    navigation.navigate('Ingredient', {
      ingredientId: item.id,
      name: item.name,
    });
  };

  renderIngredient = ({item}) => (
    <TouchableHighlight
      underlayColor="rgba(73,182,77,1,0.9)"
      onPress={() => onPressIngredient(item)}>
      <View style={styles.container}>
        <Image style={styles.photo} source={{uri: item.image}} />
        <Text style={styles.title}>{item.name}</Text>
        <Text style={{color: 'grey'}}>{item.quantity}</Text>
      </View>
    </TouchableHighlight>
  );

  useLayoutEffect(() => {
    const {title} = route.params;
    navigation.setOptions({
      title: title,
      headerTitleStyle: {
        fontSize: 16,
      },
    });
  }, [navigation]);

  useEffect(() => {
    const {recipeId} = route.params;
    axios
      .get(
        `http://192.168.0.103:8080/api/public/ingredient/searchIngredientsByRecipeId?recipeId=${recipeId}`,
      )
      .then((response) => response.data)
      .then((response) => {
        setIngredients(response.data.items);
        console.log(response.data.items);
      })
      .catch(function (error) {
        console.log(error);
      });
  }, []);

  return (
    <View>
      <FlatList
        vertical
        showsVerticalScrollIndicator={false}
        numColumns={3}
        data={ingredients}
        renderItem={renderIngredient}
        keyExtractor={(item) => `${item.recipeId}`}
      />
    </View>
  );
}

// export default class IngredientsDetailsScreen extends React.Component {
//   static navigationOptions = ({ navigation }) => {
//     return {
//       title: navigation.getParam('title'),
//       headerTitleStyle: {
//         fontSize: 16
//       }
//     };
//   };

//   constructor(props) {
//     super(props);
//   }

//   onPressIngredient = item => {
//     let name = getIngredientName(item.ingredientId);
//     let ingredient = item.ingredientId;
//     this.props.navigation.navigate('Ingredient', { ingredient, name });
//   };

//   renderIngredient = ({ item }) => (
//     <TouchableHighlight underlayColor='rgba(73,182,77,1,0.9)' onPress={() => this.onPressIngredient(item[0])}>
//       <View style={styles.container}>
//         <Image style={styles.photo} source={{ uri: item[0].photo_url }} />
//         <Text style={styles.title}>{item[0].name}</Text>
//         <Text style={{ color: 'grey' }}>{item[1]}</Text>
//       </View>
//     </TouchableHighlight>
//   );

//   render() {
//     const { navigation } = this.props;
//     const item = navigation.getParam('ingredients');
//     const ingredientsArray = getAllIngredients(item);

//     return (
//       <View>
//         <FlatList
//           vertical
//           showsVerticalScrollIndicator={false}
//           numColumns={3}
//           data={ingredientsArray}
//           renderItem={this.renderIngredient}
//           keyExtractor={item => `${item.recipeId}`}
//         />
//       </View>
//     );
//   }
// }
