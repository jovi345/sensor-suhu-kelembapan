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
            dataPoints: [],
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

      if (this.options.data[0].dataPoints.length > 10) {
        this.options.data[0].dataPoints.shift();
      }

      if (this.$refs.chart && this.$refs.chart.chart) {
        this.$refs.chart.chart.render();
      }
    },

    async getData() {
      try {
        const response = await fetch("http://localhost:8080/api/v1/data/get");
        const datas = await response.json();
        datas.forEach((data) => {
          const time = data.created_at;
          const hour = new Date(time).getHours();
          const minute = new Date(time).getMinutes();
          const second = new Date(time).getSeconds();

          const timestamp = `${hour}:${minute}:${second}`;
          const firstTemp = data.first_temperature;

          const newData = { label: timestamp, y: firstTemp };
          this.options.data[0].dataPoints.push(newData);

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
  },
  mounted() {
    // setInterval(this.addRandomData, 30000);
    this.getData();
  },
};
</script>
