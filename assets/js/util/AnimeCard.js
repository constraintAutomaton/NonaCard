import ApiInterface from "./ApiInterface.js";
export default class AnimeCard extends HTMLElement {
  constructor() {
    super();
    this.apiEngine = new ApiInterface();
    this.data = new Map();
    const shadow = this.attachShadow({ mode: "open" });
    const linkElem = document.createElement("link");
    linkElem.setAttribute("rel", "stylesheet");
    linkElem.setAttribute("href", "static/css/main.css");

    this.container = document.createElement("div");
    this.container.innerHTML = `
    <input></input>
    <button>OK</button>
    `;
    shadow.appendChild(linkElem);
    shadow.appendChild(this.container);

    this.container.querySelector("button").onclick = this.search.bind(this);
  }
  async search() {
    console.log(this);
    const query = this.container.querySelector("input").value;
    const data = await this.apiEngine.searchAnime(query);
    this.data = new Map(Object.entries(data[0]));
    const image = this.data.get("coverImage")["large"];
    this.parentElement.style.background = `url(${image}) no-repeat center`;
    console.log(this.data);
  }
}
customElements.define("anime-card", AnimeCard);
