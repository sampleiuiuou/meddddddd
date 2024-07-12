import Vue from "vue";
import VueRouter from "vue-router";
import login from "../views/login.vue";
import dashboard from "../views/dashboard.vue";
import loginhistory from "../views/loginhistory.vue";
import stockview from "../views/stockview.vue"
import adduser from "../views/adduser.vue";
import stockentry from "../views/stockentry.vue";
import billentry from "../views/billentry.vue";
import salesreport from "../views/salesreport.vue";


Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "login",
    component: login,
  },
  {
    path: "/dashboard",
    name: "dashboard",
    component: dashboard,

  },
  {
    path: "/loginhistory",
    name: "loginhistory",
    component: loginhistory,
  },
  {
    path: "/stockview",
    name: "stockview",
    component: stockview,
  },
  {
    path: "/adduser",
    name: "adduser",
    component: adduser,
  },
  {
    path: "/stockentry",
    name: "stockentry",
    component: stockentry
  },
  {
    path: "/billentry",
    name: "billentryview",
    component: billentry,
  },
  {
    path: "/salesreport",
    name: "salesreportview",
    component: salesreport,
  }

];

const router = new VueRouter({
  routes,
});

export default router;
