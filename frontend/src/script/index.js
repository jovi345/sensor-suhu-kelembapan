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

export const addNewData = async () => {
  try {
    const length = dataStore.getDatas().length;
    const newData = {
      id: length + 1,
      room_temperature: Number(Math.floor(Math.random() * 80)),
      object_temperature: Number(Math.floor(Math.random() * 80)),
      humidity: Number(Math.floor(Math.random() * 80)),
      created_at: new Date().toISOString(),
    };
    dataStore.datas.push(newData);
  } catch (error) {
    console.error(error);
  }
};
