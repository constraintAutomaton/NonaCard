import "./CardForm.js";
/**
 * AnimeCard
 * Html tag that show anime card
 */
export default class AnimeCard extends HTMLElement {
  static get observedAttributes() {
    return ["data"];
  }
  constructor() {
    super();

    const shadow = this.attachShadow({ mode: "open" });
    const linkElem = document.createElement("link");
    linkElem.setAttribute("rel", "stylesheet");
    linkElem.setAttribute("href", "static/css/animeCard.css");
    this.container = document.createElement("div");
    this.container.innerHTML = `
    <div id="info">
      <button>OK</button>
    <div>
    `;
    shadow.appendChild(linkElem);
    shadow.appendChild(this.container);
    this.container.querySelector("button").onclick = this.ShowForm.bind(this); //this.search.bind(this);
  }
  connectedCallback() {
    this.setAttribute("data", JSON.stringify({}));
  }
  async ShowForm() {
    const cardForm = document.querySelector("card-form");
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
    }
  }
  changeCardImage() {
    const data = JSON.parse(this.getAttribute("data"));
    const image = data["coverImage"]["large"];
    const container = this.container;
    container.style.background = `url(${image}) no-repeat center`;
  }
}
customElements.define("anime-card", AnimeCard);
