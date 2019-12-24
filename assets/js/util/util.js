export const save = () => {
  const arrayCard = Array.from(document.querySelectorAll("anime-card"));
  const data = {};
  for (const el of arrayCard) {
    data[el.id] = el.getAttribute("data");
  }
  const myStorage = window.localStorage;
  myStorage.setItem("data", JSON.stringify(data));
};
export const load = () => {
  const myStorage = window.localStorage;

  const arrayCard = Array.from(document.querySelectorAll("anime-card"));
  const data = JSON.parse(myStorage.getItem("data"));
  arrayCard.forEach(el => {
    el.setAttribute("data", data[el.id]);
  });
};
