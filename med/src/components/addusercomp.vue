<template>
    <div class="mt-16">
        <div class="d-flex justify-center">
            <v-card width="800">
                <div class="d-flex justify-center" min-width="400">

                    <v-expansion-panels :disabled="disabled" multiple>
                        <v-expansion-panel>
                            <v-expansion-panel-header>Item</v-expansion-panel-header>
                            <v-expansion-panel-content>
                                <div>
                                    <div>
                                        <v-text-field v-model="useridnew" label="user_id"></v-text-field>
                                    </div>
                                    <div><v-text-field v-model="passwordnew" label="password"></v-text-field></div>
                                    <div>
                                        <div class="text-center">
                                            <v-select v-model="rolenew" :items="['biller', 'manager', 'admin', 'inventory']"
                                                label="role" ></v-select>
                                        </div>
                                    </div>
                                    <div>
                                        <VBtn color="success" @click="adduser">ADD</VBtn>
                                    </div>
                                </div>
                            </v-expansion-panel-content>
                        </v-expansion-panel>
                    </v-expansion-panels>
                </div>
            </v-card>
        </div>
    </div>
</template>
<script>
export default {
    data() {
        return {
            useridnew: '',
            passwordnew: '',
            rolenew: '',
            items: [
                { title: 'biller' },
                { title: 'admin' },
                { title: 'manager' },
                { title: 'inventory' },
            ],
            closeOnClick: true,
        };
    },
    components: {

    },
    methods: {
        adduser() {
            let addobj = { userid: this.useridnew, password: this.passwordnew, role: this.rolenew };
            this.$store.commit("MUTATE_ADDUSER", addobj);
            console.log(addobj)
            this.$store.commit("MUTATE_ROLE", this.role);
            console.log(this.addobj);
            this.$router.push("/");

        }
    }


}
</script>
