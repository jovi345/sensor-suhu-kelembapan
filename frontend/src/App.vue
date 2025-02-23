<template>
  <NavigationBar />
  <div class="data">
    <ObjectTemperature />
    <RoomTemperature />
    <RoomHumidity />
  </div>
</template>

<script>
import NavigationBar from "./components/NavigationBar.vue";
import ObjectTemperature from "./components/ObjectTemperature.vue";
import RoomHumidity from "./components/RoomHumidity.vue";
import RoomTemperature from "./components/RoomTemperature.vue";
import { newData } from "./script";

export default {
  components: {
    NavigationBar,
    RoomTemperature,
    RoomHumidity,
    ObjectTemperature,
  },

  mounted() {
    setInterval(() => {
      newData();
    }, 5000);
  },
};
</script>

<style scoped>
.data {
  max-width: 1300px;
  margin: 0 auto;
}
</style>
