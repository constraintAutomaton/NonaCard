export default class ApiInterface {
  constructor() {
    const _rootUrl = `${window.location.href}/api/v1/`;
    this.getRootUrl = () => {
      return _rootUrl;
    };
  }
  async searchAnime(p_query) {
    const url = `${this.getRootUrl()}search/anime?q=${p_query}`;
    const rep = await fetch(url);
    if (rep.ok) {
      return await rep.json();
    } else {
      console.error(rep.status);
      return false;
    }
  }
  async getUserInfo(p_user) {
    const url = `${this.getRootUrl()}/user/${p_user}`;
    const rep = await fetch(url);
    if (rep.ok) {
      return await rep.json();
    } else {
      console.error(rep.status);
      return false;
    }
  }
}
