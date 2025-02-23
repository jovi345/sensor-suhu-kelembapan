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

export const newData = async () => {
  try {
    const response = await getAllData();
    const lastData = response[response.length - 1];
    const newData = {
      id: lastData.id,
      room_temperature: lastData.room_temperature,
      object_temperature: lastData.object_temperature,
      humidity: lastData.humidity,
      created_at: lastData.created_at,
    };
    dataStore.datas.push(newData);
  } catch (error) {
    console.error(error);
  }
};
