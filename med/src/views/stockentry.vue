<template>
    <div>  
        <navb></navb>
        <stockentrycomp/>
    </div>
</template>

<script>
import navb from "../components/navbar.vue";
import stockentrycomp from "../components/stockentrycomp.vue";
export default {

    data() { 
        return {

        }
    },
    components : {
        stockentrycomp, navb
    }
}
</script>

