<template>
    <div>
        <navb></navb>
        <stockviewcomp/>
    </div>
</template>

<script>
import navb from "../components/navbar.vue";
import stockviewcomp from '../components/stockviewcomp.vue'; 
export default {
    data() {
        return {

        }
    },
    components: {
        stockviewcomp,navb,
    },
}
</script>