import React, {useState, useEffect, useLayoutEffect} from 'react';
import {
  FlatList,
  Text,
  View,
  TouchableHighlight,
  Image,
} from 'react-native';
import styles from './styles';
import MenuImage from '../../components/MenuImage/MenuImage';

import axios from 'axios';

export default function HomeScreen({navigation}) {
  const [recipes, setRecipes] = useState([])

  renderRecipes = ({item}) => {
    return (
      <TouchableHighlight
        underlayColor="rgba(73,182,77,1,0.9)"
        onPress={() => onPressRecipe(item)}>
        <View style={styles.container}>
          <Image style={styles.photo} source={{uri: item.photo_url}} />
          <Text style={styles.title}>{item.title}</Text>
          <Text style={styles.category}>
            {item.categories[0]}
          </Text>
        </View>
      </TouchableHighlight>
    )
  }

  onPressRecipe = (item) => {
    navigation.navigate('Recipe', { item });
  }

  useLayoutEffect(() => {
    navigation.setOptions({
      title: 'Home',
      headerLeft: () => (
        <MenuImage
          onPress={() => {
            navigation.openDrawer();
          }}
        />
      ),
    })
  }, [navigation]);

  useEffect(() => {
    axios.get('http://192.168.0.103:8080/api/public/recipe/getAll')
      .then(response => response.data)
      .then(response => {
        setRecipes(response.data.items)
      })
      .catch(function (error) {
        console.log(error);
      })
  }, [])

  return (
    <View>
      <FlatList
        vertical
        showsVerticalScrollIndicator={false}
        numColumns={2}
        data={recipes}
        renderItem={renderRecipes}
        keyExtractor={(item) => `${item.id}`}
      />
    </View>
  )
}
