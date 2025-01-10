<template>
  <div>
    <CanvasJSChart :options="options" ref="chart" />
  </div>
</template>

<script>
export default {
  data() {
    return {
      options: {
        animationEnabled: true,
        title: {
          text: "Real-Time Line Chart",
        },
        axisY: {
          title: "Values",
        },
        data: [
          {
            type: "line",
            dataPoints: [
              { label: "apple", y: 10 },
              { label: "orange", y: 15 },
              { label: "banana", y: 25 },
              { label: "mango", y: 30 },
              { label: "grape", y: 28 },
            ],
          },
        ],
      },
    };
  },
  methods: {
    addRandomData() {
      const randomValue = Math.floor(Math.random() * 50);
      const newLabel = `Item ${this.options.data[0].dataPoints.length + 1}`;
      const newData = { label: newLabel, y: randomValue };

      this.options.data[0].dataPoints.push(newData);

      if (this.options.data[0].dataPoints.length > 10000) {
        this.options.data[0].dataPoints.shift();
      }

      if (this.$refs.chart && this.$refs.chart.chart) {
        this.$refs.chart.chart.render();
      }
    },

    async getData() {
      try {
        const response = await fetch("http://localhost:8080/api/data/get");
        console.log(response);
      } catch (error) {
        console.error(error);
      }
    },
  },
  mounted() {
    setInterval(this.addRandomData, 3000);
    setInterval(this.getData, 3000);
  },
};
</script>
