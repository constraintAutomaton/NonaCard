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
    this.container.classList.add("container");
    this.container.innerHTML = `
    <div class="form">
      <input>
      <button>Exit</button>
    </div>
    <div class="results"></div>
    `;
    shadow.appendChild(linkElem);
    shadow.appendChild(this.container);
    this.container.querySelector("button").onclick = this.exitForm.bind(this);
    this.container.querySelector("input").oninput = this.showSearchResult.bind(
      this
    );
  }
  connectedCallback() {
    this.container.style.opacity = "0";
  }
  attributeChangedCallback(name, oldValue, newValue) {
    switch (name) {
      case "card": {
        // don't show if clicked again
        if (oldValue === newValue || newValue === "null") {
          this.container.style.opacity = "0";
          if (newValue !== "null") {
            this.exitForm();
          }
        } else {
          this.container.style.opacity = "1";
          this.emptyResult();
          this.container.querySelector("input").focus();
        }
        break;
      }
    }
  }
  async showSearchResult() {
    const query = this.container.querySelector("input").value;
    const resultSection = this.container.querySelector(".results");
    resultSection.innerHTML = "";
    if (query != "") {
      const data = await this.apiEngine.searchAnime(query);
      data.forEach(el => {
        const result = document.createElement("div");
        const title =
          el.title.english != null ? el.title.english : el.title.romaji;
        result.innerHTML = `<span class="result">${title}<span>`;
        // set the select anime in the card
        result.querySelector(".result").onclick = () => {
          const card = document.querySelector(`#${this.getAttribute("card")}`);
          el["card"] = this.getAttribute("card");
          card.setAttribute("data", JSON.stringify(el));
        };
        resultSection.appendChild(result);
      });
    } else {
      this.emptyResult();
    }
  }
  emptyResult() {
    const resultSection = this.container.querySelector(".results");
    resultSection.innerHTML = "";
    this.container.querySelector("input").value = "";
  }
  exitForm() {
    const card = document.querySelector(`#${this.getAttribute("card")}`);
    card.setAttribute("selected", "false");
    this.setAttribute("card", null);
  }
}

customElements.define("card-form", CardForm);
