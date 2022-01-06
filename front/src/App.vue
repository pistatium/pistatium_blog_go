<template>
    <v-app id="app">
        <Navigation></Navigation>

        <v-container grid-list-lg class="main container">
            <v-card loading v-if="this.$root.loading">
                <v-card-text class="loading-body">
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

            <footer>
                ©2022 pistatium.
            </footer>
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
        padding: 0;
        margin-top: 24px;
        font-family: "Helvetica Neue", Arial, "Hiragino Kaku Gothic ProN", "Hiragino Sans", Meiryo, sans-serif;
        font-size: 16x;
    }
    .main {
        max-width: 740px;
    }
    footer {
        margin: 12px 0;
        font-size: 14px;
        text-align: center;
        font-weight: lighter;
        color: #c3c3c3;
    }
    .v-alert {
        padding: 24px 0;
        margin: 24px 0 32px 0;
    }
    .loading-body {
        margin: 20px;
    }
</style>
