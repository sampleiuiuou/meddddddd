<template>
  <div>
      <div class="d-flex justify-center mt-16 ">
          <VCard width="900">
              <div class="d-flex">
                  <div>
                      <p class="ml-5 pt-5 mb-n5 ">Refill Stock</p>
                  </div>
                  <VSpacer />
                  <span>
                      <v-container class="mt-4 mr-5">
                          <v-row justify="center">
                              <v-dialog v-model="dialog" persistent max-width="600px">
                                  <template v-slot:activator="{ on, attrs }">
                                      <v-btn color="primary" dark v-bind="attrs" v-on="on">
                                          + ADD
                                      </v-btn>
                                  </template>
                                  <v-card>
                                      <v-card-title>
                                          <span class="text-h5">ADD STOCK</span>
                                      </v-card-title>
                                      <v-card-text>
                                          <v-container>
                                              <v-row>
                                                  <v-col cols="12" sm="6" md="4">
                                                      <v-text-field label="Medicine Name " v-model="addname"
                                                          required></v-text-field>
                                                  </v-col>


                                                  <v-col cols="12" sm="6" md="4">
                                                      <v-text-field label="Brand " v-model="addbrand"
                                                          required></v-text-field>
                                                  </v-col>


                                              </v-row>
                                          </v-container>
                                      </v-card-text>
                                      <v-card-actions>
                                          <v-spacer></v-spacer>

                                          <v-btn color="blue darken-1"  v-show="hideadd"  @click="savemed(); dialog = false " text >
                                                ADD
                                            </v-btn>
                                      </v-card-actions>
                                  </v-card>
                              </v-dialog>
                          </v-row>
                      </v-container></span>

              </div>
              <div>
                  <div class="pa-16  ">
                      <v-select v-model="mName" :items="listitem" label="Medicine Name    "></v-select>

                      <VTextField label="Brand" v-model="mbrand" />

                      <VTextField label="Quantity" v-model="qty" />
                      <VTextField label="UnitPrice" v-model="unitPrice" />
                      <v-btn color="primary" v-show="shw" @click="update()">UPDATE</v-btn>
                  </div>
              </div>
          </VCard>
      </div>
      <popup />
  </div>
</template>

<script>
export default {

  data() {
      return {
          dialog: false,
          mName: '',
          addname: '',
          addbrand: '',
          mbrand: '',
          listitem: [],
          hideadd: false,
          qty: '',
          unitPrice: '',
          storeData: this.$store.state.stock,
          shw: false,

      }
  },
  components: {

  },
  created() {
      for (let i = 0; i < this.storeData.length; i++) {
          this.listitem.push(this.storeData[i].medicineName);

      }
  },
  watch: {
      mName() {
          for (let i = 0; i < this.storeData.length; i++) {
              if (this.mName == this.storeData[i].medicineName) {
                  this.mbrand = this.storeData[i].brand;
                  break;
              }

          }
      },
      qty() {
          if (this.qty != 0 && this.mName != '' && this.mbrand != '') {
              this.shw = true;
          }
          else {
              this.shw = false;
          }
      },
      unitPrice() {
          if (this.unitPrice != 0 && this.mName != '' && this.mbrand != '') {
              this.shw = true;
          }

          else {
              this.shw = false;
          }
      },
      addbrand() {
          if (this.addbrand != '' && this.addname != '') {
              this.hideadd = true;
          }
          else {
              this.hideadd = false;
          }
      },addname() {
          if (this.addbrand != '' && this.addname != '') {
              this.hideadd = true;
          }
          else {
              this.hideadd = false;
          }
      }

  },
  methods: {
      savemed() {
          this.dialog=true
          let medobj = { medicineName: this.addname, Brand: this.addbrand };
          this.$store.commit("MUTATE_ADD_STOCK", medobj);
          console.log(medobj);
      
      },
      update() {
          for (let i = 0; i < this.storeData.length; i++) {
              if (this.storeData[i].medicineName == this.mName) {
                  let qty = parseInt(this.storeData[i].qty) + this.quantity;
                  let upt = this.unitPrice;
                  let updater = { medicine: this.mName, brand: this.brand, quantity: qty, unitPrice: upt, i: i };
                  this.$store.commit("UPDATE_DATA", updater);
              }
          }
      }
  }
}
</script>