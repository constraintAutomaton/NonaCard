import "./util/AnimeCard.js";

const cardsElement = Array.from(document.querySelectorAll(".card"));

cardsElement.forEach((card, i) => {
  const animeCard = document.createElement("anime-card");
  card.appendChild(animeCard);
});
document.querySelector("anime-card").data = { test: 33 };
console.log(document.querySelector("anime-card").data);
//debugger;
