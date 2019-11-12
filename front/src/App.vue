<template>
    <v-app id="app">
        <Navigation></Navigation>
        <v-container grid-list-lg style="margin-top: 64px;">
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
