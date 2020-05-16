import React, {useLayoutEffect, useState, useEffect} from 'react';
import {FlatList, Text, View, Image, TouchableHighlight} from 'react-native';
import styles from './styles';
import axios from 'axios';

export default function IngredientsDetailsScreen({route, navigation}) {
  const [ingredients, setIngredients] = useState([]);

  onPressIngredient = (item) => {
    navigation.navigate('Ingredient', {
      ingredient: item,
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
