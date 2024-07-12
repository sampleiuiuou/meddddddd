<template>
  <div>
    <navbarr />
    <dashboardparent></dashboardparent>
  </div>
</template>

<script>
import navbarr from "../components/navbar.vue";
import dashboardparent from "../components/dashboardcomp/dashboardparent.vue";
// import billerdashboard from "../components/dashboardcomp/dashboard-biller.vue";
// import admininventorydashboard from "../components/dashboardcomp/dashboard-admin-inventory.vue";
// import  managerdashbpoard from "../components/dashboardcomp/dashboard-manager.vue";
export default {
  data() {
    return {

    }
  },
  components: {
    // billerdashboard,
    // admininventorydashboard,
    // managerdashbpoard,
    navbarr, dashboardparent
  }
}
</script>
    
    
    