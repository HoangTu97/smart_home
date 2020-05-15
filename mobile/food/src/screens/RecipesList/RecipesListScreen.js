import React, {useLayoutEffect, useEffect, useState} from 'react';
import {FlatList, Text, View, TouchableHighlight, Image} from 'react-native';
import styles from './styles';
import {getRecipes, getCategoryName} from '../../data/MockDataAPI';
import axios from 'axios';

export default function RecipesListScreen({route, navigation}) {
  const [recipes, setRecipes] = useState([]);

  onPressRecipe = (item) => {
    navigation.navigate('Recipe', {item});
  };

  renderRecipes = ({item}) => (
    <TouchableHighlight
      underlayColor="rgba(73,182,77,1,0.9)"
      onPress={() => onPressRecipe(item)}>
      <View style={styles.container}>
        <Image style={styles.photo} source={{uri: item.photo_url}} />
        <Text style={styles.title}>{item.title}</Text>
        <Text style={styles.category}>{getCategoryName(item.categoryId)}</Text>
      </View>
    </TouchableHighlight>
  );

  useLayoutEffect(() => {
    const {title} = route.params;
    navigation.setOptions({
      title: title,
    });
  }, [navigation]);

  useEffect(() => {
    const {cateId} = route.params;
    axios
      .get(
        `http://192.168.0.103:8080/api/public/recipe/getByCategory/${cateId}`,
      )
      .then((response) => response.data)
      .then((response) => {
        setRecipes(response.data.items)
        console.log(response.data.items)
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
        numColumns={2}
        data={recipes}
        renderItem={renderRecipes}
        keyExtractor={(item) => `${item.recipeId}`}
      />
    </View>
  );
}

// export default class RecipesListScreen extends React.Component {
//   static navigationOptions = ({ navigation }) => {
//     return {
//       title: navigation.getParam('title')
//     };
//   };

//   constructor(props) {
//     super(props);
//   }

//   onPressRecipe = item => {
//     this.props.navigation.navigate('Recipe', { item });
//   };

//   renderRecipes = ({ item }) => (
//     <TouchableHighlight underlayColor='rgba(73,182,77,1,0.9)' onPress={() => this.onPressRecipe(item)}>
//       <View style={styles.container}>
//         <Image style={styles.photo} source={{ uri: item.photo_url }} />
//         <Text style={styles.title}>{item.title}</Text>
//         <Text style={styles.category}>{getCategoryName(item.categoryId)}</Text>
//       </View>
//     </TouchableHighlight>
//   );

//   render() {
//     const { navigation } = this.props;
//     const item = navigation.getParam('category');
//     const recipesArray = getRecipes(item.id);
//     return (
//       <View>
//         <FlatList
//           vertical
//           showsVerticalScrollIndicator={false}
//           numColumns={2}
//           data={recipesArray}
//           renderItem={this.renderRecipes}
//           keyExtractor={item => `${item.recipeId}`}
//         />
//       </View>
//     );
//   }
// }
