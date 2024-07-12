<template>
    <div style="min-height:1000px;">
        <div class="mt-16">
            <p class="text-h4 d-flex justify-center mt-16  ">MENU BAR</p>
        </div>
        <div class="d-flex justify-center">
            <v-card class="d-flex justify-center  " height="400" width="500">
                <div class="mt-16">
                    <!-- <p class="text-h5   mr-16">DASH BOARD</p> -->
                    <p class="mt-16 text-h4 ">Your today sales {{ todaysales }} </p>
                </div>
            </v-card>
        </div>
    </div>
</template>

<script>
export default {data() {
        return {
            today: this.$store.state.billMaster,
            todaysales: 0,
            current:0,
        };
    },
    created() {

        this.todaysales = this.today[this.today.length - 1].netPrice;
    },
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
            let amt = parseInt(this.stockdata[i].unitPrice);
            this.finalval += qty * amt;
            this.todaysales = this.todaydata[this.todaydata.length - 1].netPrice;
        }
        
    }
}

</script>