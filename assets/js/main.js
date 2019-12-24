import "./component/AnimeCard.js";
import "./component/CardForm.js";
import { save, load } from "./util/util.js";

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
load();
