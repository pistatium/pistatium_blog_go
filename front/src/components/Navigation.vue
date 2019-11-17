<template>
    <div class="navigation">
<!--        <v-navigation-drawer v-model="drawer" app right dark>-->
<!--            <v-list dense>-->
<!--                <v-list-item link>-->
<!--                    <v-list-item-action>-->
<!--                        <v-icon>mdi-home</v-icon>-->
<!--                    </v-list-item-action>-->

<!--                    <v-list-item-content>-->
<!--                        <v-list-item-title>Home</v-list-item-title>-->
<!--                    </v-list-item-content>-->
<!--                </v-list-item>-->

<!--                <v-list-item link>-->
<!--                    <v-list-item-action>-->
<!--                        <v-icon>mdi-contact-mail</v-icon>-->
<!--                    </v-list-item-action>-->

<!--                    <v-list-item-content>-->
<!--                        <v-list-item-title>Contact</v-list-item-title>-->
<!--                    </v-list-item-content>-->
<!--                </v-list-item>-->
<!--            </v-list>-->
<!--        </v-navigation-drawer>-->

        <v-app-bar
                app
                color="light-green darken-1"
                dark
        >
            <v-toolbar-title href="/"><router-link tag="h1" to="/">Pistatium</router-link></v-toolbar-title>

            <v-spacer/>

<!--            <v-app-bar-nav-icon @click.stop="drawer = !drawer"/>-->
        </v-app-bar>
    </div>
</template>

<script>
    export default {
        name: "Navigation",
        data: () => ({
            drawer: false,
        }),
    }
</script>

<style scoped>
    h1 {
        cursor: pointer;
    }
    .navigation {
        margin-bottom: 32px;
    }
</style>
