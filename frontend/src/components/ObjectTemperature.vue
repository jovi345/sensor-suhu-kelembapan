<template>
  <h2>Object Temperature</h2>
  <div class="room">
    <LineChartTemplate
      title="Real-Time Object's Temperature Monitor"
      axis-x-title="Time"
      axis-y-title="Temperature (in °C)"
      indicator-type="object_temperature"
    />
    <DoughnutChartTemplate
      title="Current Object Temperature"
      data-type=" °C"
      indicator-type="object_temperature"
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
