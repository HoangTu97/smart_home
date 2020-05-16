import React, {useState, useEffect, useLayoutEffect} from 'react';
import {
  FlatList,
  Text,
  View,
  Image,
  TouchableHighlight
} from 'react-native';
import styles from './styles';
import axios from 'axios'

export default function CategoriesScreen({navigation}) {
  const [categories, setCategories] = useState([])

  onPressCategory = (item) => {
    navigation.navigate('RecipesList', { cateId: item.id, title: item.name });
  }

  renderCategory = ({ item }) => (
    <TouchableHighlight underlayColor='rgba(73,182,77,1,0.9)' onPress={() => onPressCategory(item)}>
      <View style={styles.categoriesItemContainer}>
        <Image style={styles.categoriesPhoto} source={{ uri: item.image }} />
        <Text style={styles.categoriesName}>{item.name}</Text>
        <Text style={styles.categoriesInfo}>{item.numberRecipes} recipes</Text>
      </View>
    </TouchableHighlight>
  )

  useLayoutEffect(() => {
    navigation.setOptions({
      title: 'Categories'
    })
  }, [navigation])

  useEffect(() => {
    axios
      .get(
        `http://192.168.0.103:8080/api/public/category/getAll`,
      )
      .then((response) => response.data)
      .then((response) => {
        setCategories(response.data.items);
        console.log(response.data.items);
      })
      .catch(function (error) {
        console.log(error);
      });
  }, [])

  return (
    <View>
      <FlatList
        data={categories}
        renderItem={renderCategory}
        keyExtractor={item => `${item.id}`}
      />
    </View>
  )
}
