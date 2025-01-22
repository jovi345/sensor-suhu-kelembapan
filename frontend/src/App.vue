<template>
  <RoomTemperature />
  <RoomHumidity />
  <ObjectTemperature />
</template>

<script>
import ObjectTemperature from "./components/ObjectTemperature.vue";
import RoomHumidity from "./components/RoomHumidity.vue";
import RoomTemperature from "./components/RoomTemperature.vue";

export default {
  components: {
    RoomTemperature,
    RoomHumidity,
    ObjectTemperature,
  },
};
</script>
