<template>
  <h2>Room Temperature</h2>
  <div class="room">
    <LineChartTemplate
      title="Real-Time Room's Temperature Monitor"
      axis-x-title="Time"
      axis-y-title="Temperature (in °C)"
      indicator-type="room_temperature"
    />
    <DoughnutChartTemplate
      title="Current Room Temperature"
      data-type=" °C"
      indicator-type="room_temperature"
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
