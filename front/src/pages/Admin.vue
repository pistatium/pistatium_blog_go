<template>
    <v-row no-gutters>

        <v-col cols="4" v-if="this.$root.isLogin">
            <AdminEntries></AdminEntries>
        </v-col>
        <v-col cols="8" v-if="!this.$root.isLogin">
            <AdminLogin></AdminLogin>
        </v-col>
        <v-col cols="8" v-if="this.$root.isLogin">
            ...
        </v-col>

    </v-row>
</template>

<script>

    import AdminLogin from "../components/AdminLogin";
    import AdminEntries from "../components/AdminEntries";

    export default {
        name: 'Admin',
        components: {
            AdminLogin,
            AdminEntries,
        },
        data: () => ({
            isLogin: false,
        }),

    }
</script>