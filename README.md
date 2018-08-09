# Go-Cake-Factory
Go coding test done for an interview under an hour

Goal <br/>
Create an app that allows us to hande cakes from a REST API.<br/>

Setup<br/>
When the app is started, the API must be able to send a list of cakes as follow :<br/>

[

 {

   "title":"Lemon cheesecake",

   "desc":"A cheesecake made of lemon",

   "rating":8,

   "image":"https://s3-eu-west-1.amazonaws.com/s3.mediafileserver.co.uk/carnation/WebFiles/RecipeImages/lemoncheesecake_lg.jpg"

 },

 {

   "title":"victoria sponge",

   "desc":"sponge with jam",

   "rating":3,

   "image":"http://www.bbcgoodfood.com/sites/bbcgoodfood.com/files/recipe_images/recipe-image-legacy-id--1001468_10.jpg"

 },

 {

   "title":"Carrot cake",

   "desc":"Bugs bunnys favourite",

   "rating":8,

   "image":"http://www.villageinn.com/i/pies/profile/carrotcake_main1.jpg"

 },

 {

   "title":"Banana cake",

   "desc":"Donkey kongs favourite",

   "rating":6,

   "image":"http://ukcdn.ar-cdn.com/recipes/xlarge/ff22df7f-dbcd-4a09-81f7-9c1d8395d936.jpg"

 },

 {

   "title":"Birthday cake",

   "desc":"a yearly treat",

   "rating":10,

   "image":"http://cornandco.com/wp-content/uploads/2014/05/birthday-cake-popcorn.jpg"

 }

]

The API must handle the following ROUTES : <br/>

GET /cakes : return a list of the cakes in JSON format, the cakes must be sorted by rank and alphabetically.

GET /cakes/:id : return the details of a cake in JSON format.

POST /cakes : Add a cake to the cakes list, the data will be sent as a JSON in the request body :

 {
    "title":"Lemon cheesecake",
    "desc":"A cheesecake made of lemon",
    "rank":8,
    "image":"https://s3-eu-west-1.amazonaws.com/s3.mediafileserver.co.uk/carnation/WebFiles/RecipeImages/lemoncheesecake_lg.jpg"
  }

DELETE /cakes/:id : delete a cake from the list

PUT /cakes/init : reset the list to the first 5 initial cakes

Sort:<br/>
By default, we want the cakes sorted by rank and alphabetically if they have the same score.
