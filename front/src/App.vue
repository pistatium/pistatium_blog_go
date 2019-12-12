<template>
    <v-app id="app">
        <Navigation></Navigation>
        <v-container grid-list-lg class="main">
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

            <footer>
                ©2019 pistatium.
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
        margin-top: 48px;
        padding: 0px;
    }
    .main {
        max-width: 740px;
    }
    footer {
        margin: 12px;
        font-size: 14px;
        text-align: center;
        font-weight: lighter;
    }
    .v-alert {
        padding: 24px;
        margin: 24px 0 32px 0;
    }

    .ads-outer {
        margin: 10px 0;
    }
</style>

