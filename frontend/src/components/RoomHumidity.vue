<template>
  <h2>Room Humidity</h2>
  <div class="room">
    <LineChartTemplate
      title="Real-Time Room's Humidity Level"
      axis-x-title="Time"
      axis-y-title="Humidity (in %)"
      indicator-type="humidity"
    />
    <DoughnutChartTemplate
      title="Current Humidity"
      data-type=" %"
      indicator-type="humidity"
      :maximum="100"
    />
  </div>
</template>

<script>
import DoughnutChartTemplate from "./DoughnutChartTemplate.vue";
import LineChartTemplate from "./LineChartTemplate.vue";

export default {
  components: {
    LineChartTemplate,
    DoughnutChartTemplate,
  },
};
</script>
