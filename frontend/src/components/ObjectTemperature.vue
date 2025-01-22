<template>
  <div>
    <ChartTemplate
      title="Object Temperature Chart"
      axis-x-title="Time"
      axis-y-title="Temperature (in °C)"
      indicator-type="object_temperature"
    />
  </div>
</template>

<script>
import ChartTemplate from "./ChartTemplate.vue";

export default {
  components: {
    ChartTemplate,
  },
};
</script>
