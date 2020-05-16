import React, {useState, useEffect, useLayoutEffect} from 'react';
import {FlatList, Text, View, Image, TouchableHighlight} from 'react-native';
import styles from './styles';
import {ListItem, SearchBar} from 'react-native-elements';
import MenuImage from '../../components/MenuImage/MenuImage';
import {
  getCategoryName,
  getRecipesByRecipeName,
  getRecipesByCategoryName,
  getRecipesByIngredientName,
} from '../../data/MockDataAPI';

export default function SearchScreen({route, navigation}) {
  const [search, setSearch] = useState({value: '', data: []});

  handleSearch = (text) => {
    var recipeArray1 = getRecipesByRecipeName(text);
    var recipeArray2 = getRecipesByCategoryName(text);
    var recipeArray3 = getRecipesByIngredientName(text);
    var aux = recipeArray1.concat(recipeArray2);
    var recipeArray = [...new Set(aux)];
    if (text == '') {
      setSearch({
        value: text,
        data: [],
      });
    } else {
      setSearch({
        value: text,
        data: recipeArray,
      });
    }
  };

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

  useEffect(() => {
    navigation.setOptions({
      headerRight: () => (
        <MenuImage
          onPress={() => {
            navigation.openDrawer();
          }}
        />
      ),
      headerTitle: () => (
        <SearchBar
          containerStyle={{
            backgroundColor: 'transparent',
            borderBottomColor: 'transparent',
            borderTopColor: 'transparent',
            flex: 1,
          }}
          inputContainerStyle={{
            backgroundColor: '#EDEDED',
          }}
          inputStyle={{
            backgroundColor: '#EDEDED',
            borderRadius: 10,
            color: 'black',
          }}
          searchIcond
          clearIcon
          lightTheme
          round
          onChangeText={(text) => handleSearch(text)}
          onClear={() => handleSearch('')}
          placeholder="Search"
          value={search.value}
        />
      ),
    });
  });

  return (
    <View>
      <FlatList
        vertical
        showsVerticalScrollIndicator={false}
        numColumns={2}
        data={search.data}
        renderItem={renderRecipes}
        keyExtractor={(item) => `${item.recipeId}`}
      />
    </View>
  );
}
