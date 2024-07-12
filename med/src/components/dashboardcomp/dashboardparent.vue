<template>
    <div>
        <div v-if="role == 'biller'">
            <billerdashboard></billerdashboard>
        </div>
        <div v-if="role == 'admin'">
            <admininventorydashboard></admininventorydashboard>
        </div>
        <div v-if="role== 'inventory'">
            <admininventorydashboard></admininventorydashboard>
        </div>
        <div v-if="role == 'manager'">
            <managerdashbpoard></managerdashbpoard>
        </div>
    </div>
</template>

<script>


import billerdashboard from "./dashboard-biller.vue";
import admininventorydashboard from "./dashboard-admin-inventory.vue";
import managerdashbpoard from "./dashboard-manager.vue";
export default {
    data() {
        return {
            role: '',
            
        }
    },
    created(){
        this.role=this.$store.state.userrole;
        console.log(this.role)
    },
    components: {
        billerdashboard,
        admininventorydashboard,
        managerdashbpoard
    }
}
</script>