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
    this.container.querySelector("#info").onclick = this.ShowForm.bind(this); //this.search.bind(this);
  }
  connectedCallback() {
    this.setAttribute("data", JSON.stringify({}));
  }
  async ShowForm() {
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
}
customElements.define("anime-card", AnimeCard);
