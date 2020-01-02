import "./CardForm.js";
/**
 * AnimeCard
 * Html tag that show anime card
 */
export default class AnimeCard extends HTMLElement {
  static get observedAttributes() {
    return ["data", "selected"];
  }
  constructor() {
    super();

    const shadow = this.attachShadow({ mode: "open" });
    const linkElem = document.createElement("link");
    linkElem.setAttribute("rel", "stylesheet");
    linkElem.setAttribute("href", "static/css/animeCard.css");
    this.container = document.createElement("div");
    this.container.classList.add("container");
    this.container.draggable = true;
    this.container.innerHTML = `
    <div id="info">
      <img class="option-card" alt="options" src="static/images/gear.svg"/>
    <div>
    `;
    shadow.appendChild(linkElem);
    shadow.appendChild(this.container);

    this.container.querySelector("#info").onclick = this.ShowForm.bind(this);

    this.container.ondragstart = this.dragstart_handler.bind(this);
    this.container.ondrop = this.drop_handler.bind(this);
    this.container.ondragover = this.dragover_handler.bind(this);
  }
  connectedCallback() {
    this.setAttribute("data", JSON.stringify({}));
  }
  async ShowForm() {
    const cards = Array.from(document.querySelectorAll("anime-card"));
    cards.forEach(card => {
      if (card.id != this.id) {
        card.setAttribute("selected", "false");
      }
    });
    this.container.classList.add("card-selected");
    const cardForm = document.querySelector("card-form");
    this.setAttribute("selected", "true");
    cardForm.setAttribute("card", this.id);
  }
  attributeChangedCallback(name, oldValue, newValue) {
    switch (name) {
      case "data": {
        if (newValue !== "{}") {
          this.changeCardImage();
        }
        break;
      }
      case "selected": {
        if (newValue === "false") {
          this.container.classList.remove("card-selected");
        }
      }
    }
  }
  changeCardImage() {
    const data = JSON.parse(this.getAttribute("data"));
    const image = data["coverImage"]["large"];
    const container = this.container;
    container.style.backgroundImage = `url(${image})`;
    container.style.boxShadow = `2px 2px 2px ${data["coverImage"]["color"]}`;
  }
  dragstart_handler(ev) {
    //ev.preventDefault();
    const data = this.getAttribute("data");
    ev.dataTransfer.setData("text/plain", data);
    ev.dataTransfer.dropEffect = "move";
  }
  dragover_handler(ev) {
    ev.preventDefault();
    ev.dataTransfer.dropEffect = "move";
  }
  drop_handler(ev) {
    ev.preventDefault();
    const data = ev.dataTransfer.getData("text/plain");

    const idTarget = ev.rangeParent.id;
    debugger;
    ev.target.parentNode.id = this.id;
    ev.target.parentNode.setAttribute("data", data);
    this.id = idTarget;
  }
}
customElements.define("anime-card", AnimeCard);
