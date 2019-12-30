import Chartist from "./../../node_modules/chartist/dist/chartist";
import ApiInterface from "./ApiInterface";

const engine = new ApiInterface();
const chartOption =  { axisY: {
    onlyInteger: true,
}};
export const userData = new Map();
export const data = {
    // A labels array that can contain any sort of values
    labels: ["Mon", "aaa", "Wed", "Thu", "Fri"],
    // Our series array that contains series objects or in this case series data arrays
    series: [[5, 2, 4, 2, 0]]
  };
const getUserInfo = async () => {
  const userName = document.querySelector("#user").value;

  const userInfo = await engine.getUserInfo(userName);

  for (const el of Object.entries(userInfo)) {
    userData.set(el[0], el[1]);
  }
  formatChart();
};
const formatChart = ()=>{
    const label_serie = []
    for (const el of userData.get("statistics").scores){
        label_serie.push([el.score,el.count])
    }
    label_serie.sort((a,b)=>b[0]-a[0]);
    const labels = [];
    const series = [];
    for (const el of label_serie){
        labels.push(el[0]);
        series.push(el[1]);
    }
    
    data.labels = labels;
    data.series = [series];

    new Chartist.Bar('.ct-chart', data,chartOption);
}
document.querySelector("#login").onclick = getUserInfo;




