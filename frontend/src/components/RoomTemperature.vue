<template>
  <div>
    <CanvasJSChart :options="options" ref="chart" />
  </div>
</template>

<script>
import { dataStore } from "@/script";

export default {
  data() {
    return {
      options: {
        animationEnabled: true,
        panEnabled: true,
        zoomEnabled: true,
        title: {
          text: "Room Temperature Chart",
        },
        axisX: {
          title: "Time",
        },
        axisY: {
          title: "Temperature (in °C)",
        },
        data: [
          {
            type: "line",
            dataPoints: [],
          },
        ],
      },
    };
  },
  methods: {
    async getData() {
      try {
        const datas = dataStore.getDatas();
        datas.forEach((data) => {
          const time = data.created_at;
          const date = new Date(time);

          const padZero = (value) => (value < 10 ? "0" + value : value);
          const hour = padZero(date.getHours());
          const minute = padZero(date.getMinutes());
          const second = padZero(date.getSeconds());

          const timestamp = `${hour}:${minute}:${second}`;
          const dataInd = data["room_temperature"];

          const newData = { label: timestamp, y: dataInd };
          this.options.data[0].dataPoints.push(newData);

          if (this.options.data[0].dataPoints.length > Math.min()) {
            this.options.data[0].dataPoints.shift();
          }

          if (this.$refs.chart && this.$refs.chart.chart) {
            this.$refs.chart.chart.render();
          }
        });
      } catch (error) {
        console.error(error);
      }
    },
  },
  async beforeMount() {
    await this.getData();
  },
};
</script>
