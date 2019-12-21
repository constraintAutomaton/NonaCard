import Card from "./util/Card.js"

const cardsElement = Array.from(document.querySelectorAll(".card"));
const cards = [];
for (let i in cardsElement){
    cards.push(new Card(cardsElement[i],i+1));
}
//debugger;