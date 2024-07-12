<template>
    <div>
        <navb></navb>
        <salesreportcomp></salesreportcomp>
    </div>
</template>

<script>
import navb from "../components/navbar.vue";
import salesreportcomp from "../components/salesreportcomp.vue";
export default {
    data() {
        return {

        };
    },
    components: {
        salesreportcomp,navb,
    }
}
</script>