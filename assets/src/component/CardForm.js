import ApiInterface from "./../util/ApiInterface.js";
export default class CardForm extends HTMLElement {
  static get observedAttributes() {
    return ["card"];
  }

  constructor() {
    super();

    const shadow = this.attachShadow({ mode: "open" });
    const linkElem = document.createElement("link");
    linkElem.setAttribute("rel", "stylesheet");
    linkElem.setAttribute("href", "static/css/cardForm.css");
    this.apiEngine = new ApiInterface();
    this.container = document.createElement("div");
    this.container.innerHTML = `
    <input></input>
    <button>Exit</button>
    <div class="result"></div>
    `;
    shadow.appendChild(linkElem);
    shadow.appendChild(this.container);
    this.container.querySelector("button").onclick = this.exitForm.bind(this);
    this.container.querySelector("input").oninput = this.showSearchResult.bind(
      this
    );
  }
  connectedCallback() {
    this.container.style.display = "none";
  }
  attributeChangedCallback(name, oldValue, newValue) {
    switch (name) {
      case "card": {
        // don't show if clicked again
        if (oldValue === newValue || newValue === "null") {
          this.container.style.display = "none";
        } else {
          this.container.style.display = "inline";
          this.emptyResult();
        }
        break;
      }
    }
  }
  async showSearchResult() {
    const query = this.container.querySelector("input").value;
    const resultSection = this.container.querySelector(".result");
    resultSection.innerHTML = "";
    if (query != "") {
      const data = await this.apiEngine.searchAnime(query);
      data.forEach(el => {
        const result = document.createElement("div");
        const title =
          el.title.english != null ? el.title.english : el.title.romaji;
        result.innerHTML = `<span>${title}<span>`;
        // set the select anime in the card
        result.querySelector("span").onclick = () => {
          const card = document.querySelector(`#${this.getAttribute("card")}`);
          card.setAttribute("data", JSON.stringify(el));
        };
        resultSection.appendChild(result);
      });
    } else {
      this.emptyResult();
    }
  }
  emptyResult() {
    const resultSection = this.container.querySelector(".result");
    resultSection.innerHTML = "";
    this.container.querySelector("input").value = "";
  }
  exitForm() {
    this.setAttribute("card", null);
  }
}

customElements.define("card-form", CardForm);
