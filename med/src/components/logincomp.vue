<template>
    <div class="blue " style="min-height:750px;">
        <div class="d-flex justify-center  ">
            <v-card width="550" class="pa-16 mt-16 "><br>
                <div class="mt-n10">
                    <p class="d-flex justify-center text-h4 ">Log In</p>
                    <v-form>
                        <VTextField class="mb-n3" v-model="userid" label="  Userid" />
                        <br>
                        <VTextField v-model="password" label="  Password" />
                        <div class="d-flex justify-center">
                            <v-btn color="primary" class=" pa-4  mt-2" width="100" @click="login()"
                                v-show="bt">Login</v-btn>
                            <!-- {{ role }} -->
                        </div>
                    </v-form>
                </div>
            </v-card>
        </div>
        <!-- <v-btn @click="loghis" >loginhistory</v-btn> -->
    </div>
</template>

<script>
export default {
    data() {
        return {
            storelogindata: this.$store.state.login,
            userid: '',
            password: '',
            role: 'sd',
            bt: false,

        }
    },
    watch: {
        userid() {

            console.log(this.userid + this.password);

            for (let i = 0; i < this.storelogindata.length; i++) {
                console.log(this.storelogindata[i].userid + " " + this.storelogindata[i].password);
                if (this.userid === this.storelogindata[i].userid && this.password === this.storelogindata[i].password) {
                    this.role = this.storelogindata[i].role;
                    this.bt = true;
                }
            }
        },
        password() {
            console.log(this.userid + this.password);
            for (let i = 0; i < this.storelogindata.length; i++) {
                console.log(this.storelogindata[i].userid + " " + this.storelogindata[i].password);
                if (this.userid === this.storelogindata[i].userid && this.password === this.storelogindata[i].password) {
                    this.role = this.storelogindata[i].role;
                    this.bt = true;
                }
            }
        }
    },
    methods: {
        login() {
            let d = new Date();
            let time = d.toLocaleTimeString();
            let date = d.toLocaleDateString('ig-ng')
            let val = this.role;
            let obj = { userid: this.userid, password: this.password, date: date, time: time, type: 'login' };
            this.$store.commit("MUTATE_ROLE", val);
            this.$store.commit("MUTATE_LOGHISTORY", obj);
            this.$router.push("/dashboard");

        }
    }
}
</script>