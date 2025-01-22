export const dataStore = {
  datas: [],
  setDatas(datas) {
    this.datas = datas;
  },
  getDatas() {
    return this.datas;
  },
};

export const getAllData = async () => {
  try {
    const response = await fetch("http://localhost:8080/api/v1/data/get");
    const datas = await response.json();
    return datas;
  } catch (error) {
    console.error(error);
  }
};
