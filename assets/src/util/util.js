export const save = () => {
  const arrayCard = Array.from(document.querySelectorAll("anime-card"));
  const data = { card: {} };
  for (const el of arrayCard) {
    data["card"][el.id] = el.getAttribute("data");
  }
  const name3x3 = document.getElementById("3x3-name").value;
  const username = document.querySelector("#user").value;
  data["name3x3"] = name3x3;
  data["username"] = username;
  const myStorage = window.localStorage;
  myStorage.setItem("data", JSON.stringify(data));
};
export const load = () => {
  const myStorage = window.localStorage;

  const arrayCard = Array.from(document.querySelectorAll("anime-card"));
  const data = JSON.parse(myStorage.getItem("data"));
  if (data !== null) {
    arrayCard.forEach(el => {
      el.setAttribute("data", data["card"][el.id]);
    });
  }
  document.getElementById("3x3-name").value = data["name3x3"];
  document.querySelector("#user").value = data["username"];
};
