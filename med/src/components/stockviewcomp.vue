<template>
    <div>
        <v-card>
            <v-card-title>
                Stock View
                <v-spacer></v-spacer>
                <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line
                    hide-details></v-text-field>
            </v-card-title>
            <v-data-table :headers="headers" :items="itemss" :search="search"></v-data-table>
        </v-card>
    </div>
</template>

<script>
export default {
    data() {
        return {
            search: '',
            itemss : this.$store.state.stock,
            headers: [
                
                { text: 'Medicine Name', value: 'medicineName' },
                { text: 'Brand', value: 'brand' },
                { text: 'Quantity', value: 'quantity' },
                { text: 'UnitPrice', value: 'unitprice' }
            ],
        }
    },
};
    
</script>