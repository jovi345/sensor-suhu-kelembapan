<template>
  <div class="room">
    <RoomTemperature />
    <RoomHumidity />
  </div>
  <div>
    <ObjectTemperature />
  </div>
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

<style scoped>
.room {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 32px;
  margin-bottom: 64px;
}
</style>
