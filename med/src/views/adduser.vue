<template>
    <div>
        <navb/>
        <addusercomp/>
    </div>
</template>

<script>
import navb from "../components/navbar.vue";
import addusercomp from "../components/addusercomp";
export default {
    data() {
        return {

        };
    },
    components : {
        addusercomp, navb
    }
}
</script>