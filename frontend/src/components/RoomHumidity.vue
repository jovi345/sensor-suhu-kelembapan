<template>
  <div>
    <ChartTemplate
      title="Room Humidity Chart"
      axis-x-title="Time"
      axis-y-title="Humidity (in %)"
      indicator-type="humidity"
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
