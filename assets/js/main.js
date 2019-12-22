import "./util/AnimeCard.js";

const cardsElement = Array.from(document.querySelectorAll(".card"));

cardsElement.forEach((card, i) => {
  const animeCard = document.createElement("anime-card");
  card.appendChild(animeCard);
});
console.log(document.querySelector("anime-card").data);
//debugger;
