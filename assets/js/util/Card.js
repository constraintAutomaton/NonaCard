import ApiInterface from "./ApiInterface.js"
export default class Card{
    
    constructor(p_el,p_position){
        const _el = p_el;
        const _apiEngine = new ApiInterface();
        this.position = p_position;
        this.data = {}
        this.getEl = ()=>{ return _el; };
        this.getApiEngine = ()=>{return _apiEngine;};
        this.init();
    }
    async init(){
        const data= await this.getApiEngine().searchAnime("dragon maid");
        this.data = data[0];
        this.getEl().querySelector(".card__img").src= this.data["coverImage"]["medium"];
        debugger;
    }
}