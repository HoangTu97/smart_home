import React, {useState, useEffect, useLayoutEffect} from 'react';
import {
  FlatList,
  ScrollView,
  Text,
  View,
  Image,
  TouchableHighlight
} from 'react-native';
import styles from './styles';
import axios from 'axios'

export default function IngredientScreen({ route, navigation }) {
  const [recipes, setRecipes] = useState([])
  const {ingredient} = route.params

  onPressRecipe = item => {
    navigation.navigate('Recipe', { item });
  }

  renderRecipes = ({ item }) => (
    <TouchableHighlight underlayColor='rgba(73,182,77,1,0.9)' onPress={() => onPressRecipe(item)}>
      <TouchableHighlight underlayColor='rgba(73,182,77,1,0.9)' onPress={() => onPressRecipe(item)}>
        <View style={styles.container}>
          <Image style={styles.photo} source={{ uri: item.photo_url }} />
          <Text style={styles.title}>{item.title}</Text>
          <Text style={styles.category}>{item.categories}</Text>
        </View>
      </TouchableHighlight>
    </TouchableHighlight>
  )

  useLayoutEffect(() => {
    const {name} = route.params
    navigation.setOptions({
      title: name
    })
  }, [navigation])

  useEffect(() => {
    axios
      .get(
        `http://192.168.0.103:8080/api/public/recipe/getByIngredient/${ingredient.id}`,
      )
      .then((response) => response.data)
      .then((response) => {
        setRecipes(response.data.items);
        console.log(response.data.items);
      })
      .catch(function (error) {
        console.log(error);
      });
  }, [])

  return (
    <ScrollView style={styles.mainContainer}>
      <View style={{ borderBottomWidth: 0.4, marginBottom: 10, borderBottomColor: 'grey' }}>
        <Image style={styles.photoIngredient} source={{ uri: '' + ingredient.image }} />
      </View>
      <Text style={styles.ingredientInfo}>Recipes with {ingredient.name}:</Text>
      <View>
        <FlatList
          vertical
          showsVerticalScrollIndicator={false}
          numColumns={2}
          data={recipes}
          renderItem={renderRecipes}
          keyExtractor={item => `${item.recipeId}`}
        />
      </View>
    </ScrollView>
  );
}
