<template>
    <div>
        <navb/>
        <billentrycomp></billentrycomp>
    </div>
</template>

<script>
import navb from "../components/navbar.vue";
import billentrycomp from "../components/billentrycomp.vue"
export default {
    data() {
        return {

        };
    },
    components : {
        billentrycomp,navb
    }
}
</script>