<template>
    <div>
        <div><p class="text-h3 d-flex justify-center">LOGIN HISTORY</p><br></div>
        <div class="d-flex justify-center" >
        <v-card height="680" width = "700" class="mt-n7">
            <v-card-title>
                <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line
                    hide-details></v-text-field>
            </v-card-title>
            <v-data-table :headers="headers" :items="loghistory" :search="search"></v-data-table>
        </v-card></div>
    </div>
</template>
<script>
export default {
    data() {
        return {
            loghistory : this.$store.state.loginHistory,
            
            search: '',
            headers: [
                { text: 'User', value: 'userid' },
                { text: 'Type', value: 'type' },
                { text: 'Date', value: 'date' },
                { text: 'Time', value: 'time' },
                
            ],
           
        }
    },
}
</script>