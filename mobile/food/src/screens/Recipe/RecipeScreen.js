import React, {useState, useEffect, useLayoutEffect} from 'react';
import {
  FlatList,
  ScrollView,
  Text,
  View,
  TouchableOpacity,
  Image,
  Dimensions,
  TouchableHighlight,
} from 'react-native';
import styles from './styles';
import Carousel, {Pagination} from 'react-native-snap-carousel';
import BackButton from '../../components/BackButton/BackButton';
import ViewIngredientsButton from '../../components/ViewIngredientsButton/ViewIngredientsButton';
import axios from 'axios';

const {width: viewportWidth} = Dimensions.get('window');

let slider1Ref = null;

export default function RecipeScreen({route, navigation}) {
  const [activeSlide, setActiveSlide] = useState(0);
  const [recipe, setRecipe] = useState({photos: [], categories: []});

  renderImage = ({item}) => (
    <TouchableHighlight>
      <View style={styles.imageContainer}>
        <Image style={styles.image} source={{uri: item}} />
      </View>
    </TouchableHighlight>
  );

  onPressIngredient = (item) => {
    var name = getIngredientName(item);
    let ingredient = item;
    navigation.navigate('Ingredient', {ingredient, name});
  };

  useLayoutEffect(() => {
    navigation.setOptions({
      headerTransparent: 'true',
      headerLeft: () => (
        <BackButton
          onPress={() => {
            navigation.goBack();
          }}
        />
      ),
    });
  }, [navigation]);

  useEffect(() => {
    const {item} = route.params;
    axios
      .get(`http://192.168.0.103:8080/api/public/recipe/detail/${item.id}`)
      .then((response) => response.data)
      .then((response) => {
        setRecipe(response.data)
        console.log(response.data)
      })
      .catch(function (error) {
        console.log(error);
      });
  }, []);

  return (
    <ScrollView style={styles.container}>
      <View style={styles.carouselContainer}>
        <View style={styles.carousel}>
          <Carousel
            ref={(c) => {
              slider1Ref = c;
            }}
            data={recipe.photos}
            renderItem={renderImage}
            sliderWidth={viewportWidth}
            itemWidth={viewportWidth}
            inactiveSlideScale={1}
            inactiveSlideOpacity={1}
            firstItem={0}
            loop={false}
            autoplay={false}
            autoplayDelay={500}
            autoplayInterval={3000}
            onSnapToItem={(index) => setActiveSlide(index)}
          />
          <Pagination
            dotsLength={recipe.photos.length}
            activeDotIndex={activeSlide}
            containerStyle={styles.paginationContainer}
            dotColor="rgba(255, 255, 255, 0.92)"
            dotStyle={styles.paginationDot}
            inactiveDotColor="white"
            inactiveDotOpacity={0.4}
            inactiveDotScale={0.6}
            carouselRef={slider1Ref}
            tappableDots={!!slider1Ref}
          />
        </View>
      </View>
      <View style={styles.infoRecipeContainer}>
        <Text style={styles.infoRecipeName}>{recipe.title}</Text>
        <View style={styles.infoContainer}>
          {recipe.categories.map((category) => (
            <TouchableHighlight
              onPress={() =>
                navigation.navigate('RecipesList', {
                  cateId: category.id,
                  title: category.name,
                })
              }>
              <Text style={styles.category}>{category.name}</Text>
            </TouchableHighlight>
          ))}
        </View>

        <View style={styles.infoContainer}>
          <Image
            style={styles.infoPhoto}
            source={require('../../../assets/icons/time.png')}
          />
          <Text style={styles.infoRecipe}>{recipe.duration} minutes </Text>
        </View>

        <View style={styles.infoContainer}>
          <ViewIngredientsButton
            onPress={() => {
              let title = 'Ingredients for ' + recipe.title;
              navigation.navigate('IngredientsDetails', {
                recipeId: recipe.id,
                title,
              });
            }}
          />
        </View>
        <View style={styles.infoContainer}>
          <Text style={styles.infoDescriptionRecipe}>{recipe.description}</Text>
        </View>
      </View>
    </ScrollView>
  );
}

/*cooking steps
<View style={styles.infoContainer}>
  <Image style={styles.infoPhoto} source={require('../../../assets/icons/info.png')} />
  <Text style={styles.infoRecipe}>Cooking Steps</Text>
</View>
<Text style={styles.infoDescriptionRecipe}>{item.description}</Text>
*/
