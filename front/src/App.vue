<template>
    <v-app id="main">
        <v-navigation-drawer v-model="drawer" app right dark>
            <v-list dense>
                <v-list-item link>
                    <v-list-item-action>
                        <v-icon>mdi-home</v-icon>
                    </v-list-item-action>

                    <v-list-item-content>
                        <v-list-item-title>Home</v-list-item-title>
                    </v-list-item-content>
                </v-list-item>

                <v-list-item link>
                    <v-list-item-action>
                        <v-icon>mdi-contact-mail</v-icon>
                    </v-list-item-action>

                    <v-list-item-content>
                        <v-list-item-title>Contact</v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
            </v-list>
        </v-navigation-drawer>

        <v-app-bar
                app
                color="light-green darken-1"
                dark
        >
            <v-toolbar-title><h1>Pistatium</h1></v-toolbar-title>

            <v-spacer/>

            <v-app-bar-nav-icon @click.stop="drawer = !drawer"/>
        </v-app-bar>

        <v-container grid-list-lg style="margin-top: 64px;">
            <v-layout row wrap>
                <v-flex xs12>
                    <Entry v-for="entry in entries" entry="entry"></Entry>
                </v-flex>
            </v-layout>
        </v-container>
    </v-app>
</template>

<script>
    import Entry from "./components/Entry";
    export default {
        name: 'LayoutsDemosBaselineFlipped',
        components: {Entry},
        props: {
            source: String,
        },
        data: () => ({
            drawer: false,
            entries: [],
        }),
    }
</script>
