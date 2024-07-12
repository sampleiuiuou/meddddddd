<template>
    <div>
        <div>

        </div>
        <div>
            <div class="d-flex justify-center pa-10 ">
                <v-card class="pa-10 justify-center " min-height="800">
                    <div class="d-flex justify-center">
                        <v-card class=" justify-center" width="800">
                            <div class="d-flex justify-center" min-width="400">

                                <v-expansion-panels :disabled="disabled" multiple>
                                    <v-expansion-panel>
                                        <v-expansion-panel-header>Item</v-expansion-panel-header>
                                        <v-expansion-panel-content>
                                            <div class="d-flex">
                                                <div>
                                                    <div class="text-center mr-5 ">
                                                        <v-autocomplete v-model="mname" :items="medicines"
                                                            :search-input.sync="search2" class="mx-4" flat hide-no-data
                                                            hide-details label="Medicine Name"></v-autocomplete>
                                                    </div>
                                                </div>
                                                <div><v-text-field class="mr-5" v-model="qty"
                                                        label="Quantity"></v-text-field>
                                                </div>
                                                <div>
                                                </div>
                                                <div>
                                                    <VBtn color="primary" class="d-flex ml-16 mr-n16 mb-5 " width="100"
                                                        @click="addbutton()">ADD</VBtn>
                                                </div>
                                            </div>
                                        </v-expansion-panel-content>
                                    </v-expansion-panel>
                                </v-expansion-panels>
                            </div>
                        </v-card>
                    </div>
                    <div class="d-flex mt-10 ">
                        <div class="ml-2 mt-1 mr-3  "><v-row justify="center">
                                <v-dialog v-model="dialog" persistent max-width="500px">
                                    <template v-slot:activator="{ on, attrs }">
                                        <v-btn color="primary" dark v-bind="attrs" v-on="on">
                                            PREVIEW
                                        </v-btn>
                                    </template>
                                    <v-card>
                                        <v-card-title>
                                            <span class="text-h5">SK Pharmacy</span>
                                        </v-card-title>
                                        <v-card-text>
                                            <v-container class="d-flex  ">
                                                <p class="mr-16 ml-14 ">Medicine name</p>
                                                <p class="mr-16">Qty</p>
                                                <p class="mr-16">Amount</p>
                                            </v-container>
                                            <v-container class="mt-n4">
                                                <div class="d-flex  " v-for="(value, index) in desserts" :key="index">
                                                    <p class="mr-16 ml-14 ">{{ value.medicinenamelist }}</p>
                                                    <p class="mr-16 ml-3">{{ value.quantity }}</p>
                                                    <p class="mr-16 ml-5 ">{{ value.Amount }}</p>
                                                </div>

                                            </v-container>
                                        </v-card-text>
                                        <v-card-actions>
                                            <v-spacer></v-spacer>
                                            <v-btn color="blue darken-1" text @click="dialog = false">
                                                PRINT
                                            </v-btn>
                                        </v-card-actions>
                                    </v-card>
                                </v-dialog>
                            </v-row></div>
                        <v-btn color="primary" class="ml-3 mt-n2 " @click="savedata()">SAVE</v-btn>
                        <p class="ml-16">BILL NO : {{ billno }}</p>
                        <p class="ml-16">DATE : {{ dat }}</p>
                        <p class="ml-16">TOTAL : {{ totalamt }}</p>
                        <p class="ml-16">GST : {{ gst }}</p>
                        <p class="ml-16 pr-13 ">Net Payable : {{ payable }}</p>
                    </div>
                    <v-spacer></v-spacer>
                    <div class="mt-2 mr-n4 d-flex justify-end ">
                        <v-btn color="primary">Download</v-btn>
                    </div>
                    <div class="mt-5"><template>
                            <v-card elevation-1>
                                <v-card-title>
                                    <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line
                                        hide-details></v-text-field>
                                </v-card-title>
                                <v-data-table :headers="headers" :items="desserts" :search="search"></v-data-table>
                            </v-card>
                        </template></div>
                </v-card>

            </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            disabled: false,
            billername: this.$store.state.login,
            dat: '',
            gst: 0,
            payable: 0,
            totalamt: 0,
            date: '',
            amt: 0,
            counts: 0,
            billno: '',
            dialog: false,
            search2: '',
            search: '',
            brand: '',
            qty: '',
            medicines: [],
            medMaster: this.$store.state.medicineMaster,
            stockrarray: this.$store.state.stock,
            role: this.$store.state.userrole,
            desserts: [],
            mname: '',
            // tablequantity: [],
            billMaster: [],
            name: '',
            billdetails: [],
            // unitprices:[],
            headers: [
                {
                    text: 'Medicine Name',
                    align: 'start',
                    filterable: false,
                    value: 'medicinenamelist',
                },
                { text: 'Brand', value: 'brand' },
                { text: 'Quantity', value: 'quantity' },
                { text: 'Amount', value: 'Amount' },
            ],


        };
    },
    

    created() {
        // medicineName: {
        for (let i = 0; i < this.medMaster.length; i++) {
            console.log(this.medMaster[i]);
            this.medicines.push(this.medMaster[i].medicineName);

        }
        // }
    },
    methods: {
        addbutton() {
            for (let i = 0; i < this.stockrarray.length; i++) {
                if (this.stockrarray[i].medicineName == this.mname) {
                    this.brand = this.stockrarray[i].brand;
                    let total = parseInt(this.qty) * this.stockrarray[i].unitprice;
                    console.log(total);
                    this.amt = this.stockrarray[i].unitprice;
                    this.desserts.push({ medicinenamelist: this.mname, brand: this.brand, quantity: this.qty, Amount: total });
                    // this.tablequantity.push(this.stockrarray);
                }

            }

            if (this.counts < 1) {
                this.billno = Math.floor(Math.random() * 1000);
                this.counts += 1;
            }

            let date = new Date();
            this.dat = date.toLocaleDateString('en-ca');
            this.totalamt += parseInt(this.qty) * this.amt;
            this.gst =18;
            this.payable = this.totalamt + (this.totalamt*(this.gst/100));
            this.qty = '';
            this.mname = '';


        },
        savedata() {
            for (let i = 0; i < this.billername.length; i++) {
                if (this.role == this.billername[i].role) {
                    this.name = this.billername[i].userid;
                }
            }
            this.billMaster.push({ billNo: this.billno, billDate: this.dat, billAmount: this.totalamt, billGst: this.gst, netPrice: this.payable, userId: this.name })
            this.$store.commit("BILL_MASTER_UPDATE", this.billMaster);
            //   console.log(this.billMaster); 
            for (let j = 0; j < this.stockrarray.length; j++) {
                for (let i = 0; i < this.desserts.length; i++) {
                    if (this.desserts[i].medicinenamelist == this.stockrarray[j].medicineName) {
                        this.billdetails.push({ billNo: this.billno, medicinename: this.desserts[i].medicinenamelist, quantity: this.desserts[i].quantity, unitPrice: this.stockrarray[j].unitprice, amount: this.desserts[i].Amount })
                    }
                }
            }
            console.log(this.billdetails); console.log(this.billMaster);
            this.$store.commit("BILLDETAILS_UPDATE", this.billdetails);
            this.$store.commit("STOCK_QUANTITY_REDUCE", this.desserts);
            this.$router.push("/");

        }
    }
}
</script>