<template>
    <div>
        <v-card color="">
            <div class="text--primary entry-body" v-html="markdown"></div>
            <h1 class="entry-title light-green--text text--darken-3">Portfolio</h1>

        </v-card>
    </div>
</template>

<script>
    export default {
        name: 'Portfolio',
        data: () => ({
            drawer: null,
        }),
        mounted() {
            document.title = "Portfolio - Pistatium"
        }
    }
</script>

<style scoped>

    .v-card {
        margin: 12px 0 24px 0;
        padding: 0 12px 12px 12px;
        border-top: 4px solid #7cb342;
    }

    .entry {
        font-size: 108%;
        line-height: 200%;
    }

    @media screen and (max-width: 480px) {
        .v-card .entry {
            padding-left: 0;
            padding-right: 0;
        }
    }

    .entry-title {
        font-weight: 600;
        line-height: 120%;
    }
</style>
