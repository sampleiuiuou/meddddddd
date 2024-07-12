<template>
    <div>
        <div>
            <p class="text-h5  d-flex ml-6 ">SALES REPORT</p><br>
            <v-card class="d-block justify-center ml-16 mr-n16 pa-6 " width="1000">
                <v-card class="d-flex " width="1000">

                    <template class="pa-4">
                        <v-menu ref="menu" v-model="menu" :close-on-content-click="false" :return-value.sync="fromdate"
                            transition="scale-transition" offset-y min-width="auto">
                            <template v-slot:activator="{ on, attrs }">
                                <v-text-field v-model="fromdate" label="From Date" prepend-icon="mdi-calendar" readonly
                                    v-bind="attrs" v-on="on"></v-text-field>
                            </template>
                            <v-date-picker v-model="fromdate" no-title scrollable>
                                <v-spacer></v-spacer>
                                <v-btn text color="primary" @click="menu = false">
                                    Cancel
                                </v-btn>
                                <v-btn text color="primary" @click="$refs.menu.save(fromdate)">
                                    OK
                                </v-btn>
                            </v-date-picker>
                        </v-menu> <v-menu ref="menu2" v-model="menu2" :close-on-content-click="false"
                            :return-value.sync="todate" transition="scale-transition" offset-y min-width="auto">
                            <template v-slot:activator="{ on, attrs }">
                                <v-text-field v-model="todate" label="To Date" prepend-icon="mdi-calendar" readonly
                                    v-bind="attrs" v-on="on"></v-text-field>
                            </template>
                            <v-date-picker v-model="todate" no-title scrollable>
                                <v-spacer></v-spacer>
                                <v-btn text color="primary" @click="menu2 = false">
                                    Cancel
                                </v-btn>
                                <v-btn text color="primary" @click="$refs.menu2.save(todate)">
                                    OK
                                </v-btn>
                            </v-date-picker>
                        </v-menu></template>

                    <v-btn width="70" class="ml-n16 mr-16 pl-3 pr-3 mt-5" @click="searchbutton()" color="primary">Search
                    </v-btn>

                </v-card>
                <v-card width="1000">
                    <v-card-title>
                        <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line
                            hide-details></v-text-field>
                    </v-card-title>
                    <v-data-table :headers="headers" :items="desserts" :search="search"></v-data-table>
                </v-card></v-card>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            // date: (new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10),
            menu: false,
            modal: false,
            menu2: false,
            search: '',
            fromdate: '',
            todate: '',
            billDetails: this.$store.state.billDetails,
            billMaster: this.$store.state.billMaster,
            headers: [
                {
                    text: 'Bill No',
                    align: 'start',
                    filterable: false,
                    value: 'billno',
                },
                { text: 'Bill Date', value: 'billdate' },
                { text: 'Medicine Name', value: 'medicinename' },
                { text: 'Qty', value: 'qty' },
                { text: ' Amount', value: 'amount' },
            ],
            desserts: []
        };
    }
    ,
    methods: {
        searchbutton() {
            this.billMaster= this.$store.state.billMaster
            for (let i = 0; i < this.billMaster.length; i++) {
                for (let j = 0; j < this.billDetails.length; j++) {
                    const bdate = new Date(this.billMaster[i].billDate);
                    const fdate = new Date(this.fromdate);
                    const tdate = new Date(this.todate);
                    if (bdate >= fdate && bdate <= tdate) {
                        if (this.billDetails[j].billNo == this.billMaster[i].billNo) {
                            this.desserts.push({ billno: this.billDetails[j].billNo, billdate: this.billMaster[i].billDate, medicinename: this.billDetails[j].medicineName, qty: this.billDetails[j].quantity, amount: this.billDetails[j].amount })
                        }
                    }

                }
            }
        }
    }

}
</script>