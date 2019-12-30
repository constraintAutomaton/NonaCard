import _ from "lodash";
import "./component/AnimeCard.js";
import "./component/CardForm.js";
//import { data } from "./util/user";
import { save, load } from "./util/util.js";


const cardsElement = Array.from(document.querySelectorAll(".card"));
cardsElement.forEach((card, i) => {
  const animeCard = document.createElement("anime-card");
  animeCard.id = `card-${i}`;
  card.appendChild(animeCard);
});
const form = document.createElement("card-form");
document.querySelector(".core").appendChild(form);
document.querySelector("#save").onclick = save;
document.querySelector("#load").onclick = load;

load();
