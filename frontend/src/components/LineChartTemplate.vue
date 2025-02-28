<template>
  <div>
    <CanvasJSChart :options="options" ref="chart" />
  </div>
</template>

<script>
import { dataStore } from "@/script";

export default {
  props: {
    title: { type: String, required: true },
    axisXTitle: { type: String, required: true },
    axisYTitle: { type: String, required: true },
    indicatorType: { type: String, required: true },
  },
  data() {
    return {
      options: {
        animationEnabled: true,
        panEnabled: true,
        zoomEnabled: true,
        title: { text: this.title, fontSize: 22 },
        axisX: { title: this.axisXTitle },
        axisY: { title: this.axisYTitle },
        data: [{ type: "spline", dataPoints: [] }],
      },
    };
  },
  methods: {
    formatNewData(data) {
      const time = data.created_at;
      const date = new Date(time);

      const padZero = (value) => (value < 10 ? "0" + value : value);
      const hour = padZero(date.getHours());
      const minute = padZero(date.getMinutes());
      const second = padZero(date.getSeconds());

      const timestamp = `${hour}:${minute}:${second}`;
      const dataInd = data[this.indicatorType];

      const newData = { label: timestamp, y: dataInd };
      this.options.data[0].dataPoints.push(newData);
    },

    async getData() {
      try {
        const datas = dataStore.getDatas();
        datas.forEach((data) => {
          this.formatNewData(data);

          if (this.options.data[0].dataPoints.length > 10) {
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

    async updateData() {
      try {
        const datas = dataStore.getDatas();
        const newData = datas[datas.length - 1];
        this.formatNewData(newData);

        if (this.options.data[0].dataPoints.length > Math.min()) {
          this.options.data[0].dataPoints.shift();
        }

        if (this.$refs.chart && this.$refs.chart.chart) {
          this.$refs.chart.chart.render();
        }
      } catch (error) {
        console.error(error);
      }
    },
  },

  async beforeMount() {
    await this.getData();
  },

  mounted() {
    setInterval(() => {
      this.updateData();
    }, 5000);
  },
};
</script>
