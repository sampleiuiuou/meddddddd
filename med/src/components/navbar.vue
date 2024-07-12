<template>
  <div>
    <v-app-bar app color="light-blue lighten-2" flat white_text>
      <v-avatar :color="$vuetify.breakpoint.smAndDown ? 'grey darken-1' : 'transparent'" size="32">MED </v-avatar>

      <v-tabs v-if="(role == 'biller')" centered class="ml-n9" color="white--text">
        <!-- Biller -->
        <v-tab @click="dashboardfunction()">
          Dashboard
        </v-tab>
        <v-tab @click="stockviewfunction()">
          stock-view
        </v-tab>
        <v-tab @click="billentryfunction()">
          Bill entry
        </v-tab>
      </v-tabs>
      <v-tabs v-if="(role == 'inventory')" centered class="ml-n9" color="white--text">
        <!-- Inventory -->
        <v-tab @click="dashboardfunction()">
          Dashboard
        </v-tab>
        <v-tab @click="stockviewfunction()">
          stock-view
        </v-tab>
        <v-tab @click="stockentryfunction()">
          Stock entry
        </v-tab>

      </v-tabs> <v-tabs v-if="role == 'manager'" centered class="ml-n9" color="white--text">
        <!-- Manager -->
        <v-tab @click="dashboardfunction()">
          Dashboard
        </v-tab>
        <v-tab @click="stockviewfunction()">
          stock-view
        </v-tab>
        <v-tab @click="stockentryfunction()">
          Stock Entry
        </v-tab>
        <v-tab @click="salesreportfunction()">
          Sales Report
        </v-tab>

      </v-tabs> <v-tabs v-if="role == 'admin'" centered class="ml-n9" color="white--text">
        <!-- System Admin -->
        <v-tab @click="dashboardfunction()">
          Dashboard
        </v-tab>
        <v-tab @click="loginhistoryfunction()">
          Login History
        </v-tab>
        <v-tab @click="adduserfunction()">
          Add User
        </v-tab>

      </v-tabs>
      <v-btn @click="logout()">LOGOUT</v-btn>
    </v-app-bar>

    <v-container>
      <v-row>
        <v-col cols="12" sm="2">
        </v-col>

        <v-col cols="12" sm="8">
          <!-- min-height="70vh" -->
          <v-sheet rounded="lg">
            <!--  -->
          </v-sheet>
        </v-col>

        <v-col cols="12" sm="2">
          <!-- min-height="268" -->
          <v-sheet rounded="lg">
            <!--  -->
          </v-sheet>
        </v-col>
      </v-row>
    </v-container>

  </div>
</template>
  
<script>
export default {
  data() {
    return {
      role: this.$store.state.userrole,

    }
  },
  methods: {
    dashboardfunction() {
      this.$router.push("/dashboard");
    },

    stockviewfunction() {
      this.$router.push("/stockview");
    },
    stockentryfunction() {
      this.$router.push("/stockentry");
    },
    salesreportfunction() {
      this.$router.push("/salesreport");
    },
    billentryfunction() {
      this.$router.push("/billentry");
    },
    adduserfunction() {
      this.$router.push("/adduser");
    },
    loginhistoryfunction() {
      this.$router.push("/loginhistory");
    },
    logout() {
      let d = new Date();
      let date = d.toLocaleDateString("in-ng")
      let time = d.toLocaleTimeString();
      // let val = this.role;
      let obj = { userid: this.userid, password: this.password, date: date, time: time, type: 'logout' };
      this.$store.commit("MUTATE_LOGHISTORY", obj);
      this.$router.push("/");
    }
  }
}
</script>
  
  
  