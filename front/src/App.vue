<template>
    <v-app id="app">
        <Navigation></Navigation>
        <v-container grid-list-lg>
            <v-card loading v-if="this.$root.loading">
                <v-card-text>
                    Loading...
                </v-card-text>
            </v-card>

            <v-alert
                    dense
                    outlined
                    type="error"
                    v-if="this.$root.error "
            >
                {{ this.$root.error }}<br>
            </v-alert>
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
        mounted() {
        },
    }
</script>

<style scoped>
    .container {
        margin-top: 48px;
        padding: 0px;
    }
</style>
