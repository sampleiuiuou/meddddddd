import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    login: [
      {
        userid: "suresh", password: "1", role: "biller"
      }, {
        userid: "ragul", password: "2", role: "manager"
      }, {
        userid: "krishna", password: "3", role: "admin"
      }, {
        userid: "yuvan", password: "4", role: "inventory"
      },
    ],

    loginHistory: [

    ],
    medicineMaster: [
      {
        medicineName: "Acebrophylline", brand: 'Sun Pharma'
      }, {
        medicineName: "Aceclofenac", brand: 'Aurobindo'
      }, {
        medicineName: "Acitretin", brand: 'Dr. Reddy’s Pharma'
      }, {
        medicineName: "Ambroxil", brand: 'lupin Pharma'
      }, {
        medicineName: "Aripiprazole", brand: 'cipla Pharma'
      },

    ],
    stock: [
      {
        medicineName: "Acebrophylline", brand: 'Sun Pharma', quantity: "100", unitprice: "10"
      }, {
        medicineName: "Aceclofenac", brand: 'Aurobindo', quantity: "100", unitprice: "30"
      }, {
        medicineName: "Acitretin", brand: 'Dr. Reddy’s Pharma', quantity: "100", unitprice: "40"
      }, {
        medicineName: "Ambroxil", brand: 'lupin Pharma', quantity: "100", unitprice: "70"
      }, {
        medicineName: "Aripiprazole", brand: 'cipla Pharma', quantity: "100", unitprice: "50"
      },
    ],
    billMaster: [
      // {
      //   billNo: "", billDate: "", billAmount: "", billGst: "", netPrice: "", userId: ""
      // }{
      {
        billNo: "21", billDate: "2024-01-13", billGst: "645", netPrice: "4764", userid: "suresh"
      }, {
        billNo: "30", billDate: "2024-01-13", billGst: "645", netPrice: "4764", userid: "suresh"
      }, {
        billNo: "4", billDate: "2024-01-15", billGst: "645", netPrice: "4764", userid: "suresh"
      }, {
        billNo: "34", billDate: "2024-01-17", billGst: "645", netPrice: "4764", userid: "suresh"
      }, {
        billNo: "24", billDate: "2024-01-18", billGst: "645", netPrice: "4764", userid: "suresh"
      }, {
        billNo: "23", billDate: "2024-01-19", billGst: "645", netPrice: "4764", userid: "suresh"
      }
    ],
    billDetails: [
      // {
      //   billNo: "", medicineName: "", quantity: "", unitPrice: "", amount: ""
      // },
      {
        billNo: "21", medicineName: "Acebrophylline", quantity: "100", unitPrice: "10",amount: "534"
      }, {
        billNo: "30", medicineName: "Aceclofenac", quantity: "100", unitPrice: "30",amount: "46"
      }, {
        billNo: "4", medicineName: "Acitretin", quantity: "100", unitPrice: "30",amount: "453"
      }, {
        billNo: "34", medicineName: "Ambroxil", quantity: "100", unitPrice: "70",amount: "454"
      }, {
        billNo: "24", medicineName: "Aripiprazole", quantity: "100", unitPrice: "50",amount: "32"
      }, {
        billNo: "23", medicineName: "Aripiprazole", quantity: "100", unitPrice: "50",amount: "123"
      }
    ],
    userrole: null,
  },
  mutations: {
    MUTATE_ROLE(state, val) {
      state.userrole = val;
    },
    MUTATE_ADDUSER(state, final) {
      state.login.push(final);
    },
    MUTATE_LOGHISTORY(state, obj) {
      state.loginHistory.push(obj);
    },
    MUTATE_ADD_STOCK(state, medobj) {
      state.medicineMaster.push(medobj);
    },
    UPDATE_DATA(state, updater) {
      state.stock[updater.i].medicineName = updater.medicine;

      state.stock[updater.i].brand = updater.brand;
      state.stock[updater.i].quantity = updater.quantity;
      state.stock[updater.i].unitprice = updater.unitPrice;
    },
    BILL_MASTER_UPDATE(state, billobj) {
      // state.billMaster.push(billobj);
      for (let i = 0; i < billobj.length; i++) {
        state.billMaster.push({
          billNo: billobj[i].billNo,
          billDate: billobj[i].billDate,
          billGst: billobj[i].billAmount,
          netPrice: billobj[i].netPrice,
          userid: billobj[i].userId
        })

      }
    },
    BILLDETAILS_UPDATE(state, newupdate) {
      for (let i = 0; i < newupdate.length; i++) {
        state.billDetails.push({
          billNo: newupdate[i].billNo,
          medicineName: newupdate[i].medicinename,
          quantity: newupdate[i].quantity,
          unitPrice: newupdate[i].unitPrice,
          amount:newupdate[i].amount
        })

      }
    },
    STOCK_QUANTITY_REDUCE(state, valuee) {
      for (let j = 0; j < valuee.length; j++) {
        if (valuee[j].medicinenamelist == state.stock[j].medicineName) {
          state.stock[j].quantity = state.stock[j].quantity - valuee[j].quantity;
        }

      }
    },
    MUTATE_ADD_STOCK_PUSH(state, value5) {
      state.stock.push(value5);
    }
  },
  actions: {},
  modules: {},
});
