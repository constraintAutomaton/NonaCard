export default class SearchComponent extends HTMLElement {
  constructor() {
    super();
    const shadow = this.attachShadow({ mode: "open" });
    const linkElem = document.createElement("link");
    linkElem.setAttribute("rel", "stylesheet");
    linkElem.setAttribute("href", "static/css/main.css");

    const container = document.createElement("div");
    const textInput = document.createElement("input");
    const confirmButton = document.createElement("button");
    const test = document.createElement("h1");
    test.innerText="TEST";

    
    shadow.appendChild(linkElem);
    container.appendChild(textInput);
    container.appendChild(confirmButton);
    container.appendChild(test);
    shadow.appendChild(container);
  }
}
customElements.define("search-component", SearchComponent);

