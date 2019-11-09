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
                    <v-card color="grey lighten-5">

                        <v-card-text>
                            2019-11-09

                            <p class="display-1 font-weight-black light-green--text text--darken-3">
                                App Engine + Go + Vue.js + Vuetify でつくる SPA ブログ
                            </p>

                            <div class="text--primary">
                                <p>
                                    ここに本文が入ります。ここに本文が入ります。ここに本文が入ります。ここに本文が入ります。ここに本文が入ります。ここに本文が入ります。
                                </p>
                            </div>
                        </v-card-text>
                        <v-card-actions>
                            <v-btn
                                    text
                                    color="green darken-3 accent-4"
                            >
                                &gt;&gt;&gt; 続きを読む
                            </v-btn>
                        </v-card-actions>
                    </v-card>
                </v-flex>
            </v-layout>
        </v-container>
    </v-app>
</template>

<script>
    export default {
        name: 'LayoutsDemosBaselineFlipped',
        props: {
            source: String,
        },
        data: () => ({
            drawer: false,
        }),
    }
</script>
