import "./component/AnimeCard.js";
import "./component/CardForm.js";
import { save, load } from "./util/util.js";
import ApiInterface from "./util/ApiInterface.js";



const cardsElement = Array.from(document.querySelectorAll(".card"));
cardsElement.forEach((card, i) => {
  const animeCard = document.createElement("anime-card");
  animeCard.id = `card-${i}`;
  card.appendChild(animeCard);
});
const form = document.createElement("card-form");
document.querySelector(".intro").appendChild(form);
document.querySelector("#save").onclick = save;
document.querySelector("#load").onclick = load;
const engine = new ApiInterface();
const test = async ()=>{
  console.log(await engine.getUserInfo("constraintAutomaton"));

}
test();
load();
var data = {
  // A labels array that can contain any sort of values
  labels: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri'],
  // Our series array that contains series objects or in this case series data arrays
  series: [
    [5, 2, 4, 2, 0]
  ]
};

// Create a new line chart object where as first parameter we pass in a selector
// that is resolving to our chart container element. The Second parameter
// is the actual data object.
new Chartist.Line('.ct-chart', data);
