<template>
    <v-app id="app">
        <Navigation></Navigation>
        <v-container grid-list-lg>
            <router-view></router-view>
        </v-container>
    </v-app>
</template>

<script>
    import Navigation from "./components/Navigation";

    export default {
        name: 'LayoutsDemosBaselineFlipped',
        components: {
            Navigation
        },
        props: {
            source: String,
        },
        data: () => ({
            drawer: false,
        }),
        mounted () {},
    }
</script>

<style scoped>
    .container {
        margin-top: 48px;
        padding: 0px;
    }
</style>
