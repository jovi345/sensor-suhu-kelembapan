<template>
  <div>
    <ChartTemplate
      title="Room Temperature Chart"
      axis-x-title="Time"
      axis-y-title="Humidity (in %)"
      indicator-type="room_temperature"
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
