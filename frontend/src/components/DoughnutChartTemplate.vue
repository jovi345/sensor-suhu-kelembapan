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
    indicatorType: { type: String, required: true },
    dataType: { type: String, required: true },
    maximum: { type: Number, required: true },
  },
  data() {
    return {
      options: {
        animationEnabled: true,
        title: { text: this.title, fontSize: 22 },
        toolTip: { enabled: false },
        data: [
          {
            type: "doughnut",
            innerRadius: "70%",
            startAngle: 90,
            dataPoints: [],
          },
        ],
        subtitles: [
          {
            text: this.dataType,
            verticalAlign: "center",
            fontSize: 24,
            dockInsidePlotArea: true,
          },
        ],
      },
    };
  },
  methods: {
    async getData() {
      try {
        const datas = dataStore.getDatas();
        const lastData = datas[datas.length - 1];
        const indicatorType = lastData[this.indicatorType];

        const newData = [
          { y: indicatorType, name: "Humidity", color: "#285c99" },
          { y: this.maximum - indicatorType, name: "", color: "#e0e0e0" },
        ];
        this.options.data[0].dataPoints = newData;
        this.options.subtitles[0].text = indicatorType + this.dataType;

        if (this.$refs.chart && this.$refs.chart.chart) {
          this.$refs.chart.chart.render();
        }
      } catch (error) {
        console.error(error);
      }
    },
  },
  beforeMount() {
    this.getData();
  },

  mounted() {
    setInterval(() => {
      this.getData();
    }, 3000);
  },
};
</script>
