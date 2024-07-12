<template>
    <div>
        <navb></navb>
        <loginhistory/>
    
    </div>
</template>

<script>
import navb from "../components/navbar.vue";
import loginhistory from "../components/loginhistorycomp.vue";
export default {
    name : "loginhistorycomponent",
    data() {
        return {

        }
    },
    components: {
        loginhistory,
        navb,
    },
    // },
    // methods : {

    // }
}
</script>