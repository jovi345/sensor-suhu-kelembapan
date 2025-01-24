<template>
  <ObjectTemperature />
  <RoomTemperature />
  <RoomHumidity />
</template>

<script>
import ObjectTemperature from "./components/ObjectTemperature.vue";
import RoomHumidity from "./components/RoomHumidity.vue";
import RoomTemperature from "./components/RoomTemperature.vue";
import { addNewData } from "./script";

export default {
  components: {
    RoomTemperature,
    RoomHumidity,
    ObjectTemperature,
  },

  mounted() {
    setInterval(() => {
      addNewData();
    }, 3000);
  },
};
</script>
