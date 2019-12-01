<template>
    <v-row no-gutters>

        <v-col cols="4">
            aaaaab
        </v-col>
        <v-col cols="8">
            <AdminLogin></AdminLogin>
        </v-col>
    </v-row>
</template>

<script>

    import AdminLogin from "../components/AdminLogin";

    export default {
        name: 'Admin',
        components: {
            AdminLogin,
        },
        data: () => ({
            isLogin: false,
        }),

    }
</script>