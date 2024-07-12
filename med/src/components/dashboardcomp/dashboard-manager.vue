<template>
    <div class="blue lighten-5 mt-n6 " style="min-height:1000px;">
        <div class="d-flex justify-center">
            <v-card class="d-flex justify-center mt-16 mb-16 " height="200" width="600">
                <h1 class="mt-16 mb-n16 pt16">TODAY SALES : {{ todaysales }} </h1>
            </v-card>
        </div>
        <div class="d-flex justify-center">
            <v-card class="d-flex justify-center mt-16 mb-16 " height="200" width="600">
                <h1 class="mt-16 mb-n16 pt16">Current Inventory Value : {{ finalval }} </h1>
            </v-card>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            todaydata: this.$store.state.billMaster, 
            stockdata: this.$store.state.stock,
            todaysales: 0,
            finalval: 0,
        }
    },
    created() {
        for (let i = 0; i < this.stockdata.length; i++) {
            let qty = parseInt(this.stockdata[i].quantity);
            let amt = parseInt(this.stockdata[i].unitprice);
            this.finalval = this.finalval + (qty * amt);
            this.todaysales = this.todaydata[this.todaydata.length - 1].netPrice;
            console.log(this.finalval);
        }

    }
}
</script>