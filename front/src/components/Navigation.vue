<template>
    <div class="navigation">
        <v-navigation-drawer v-model="drawer" app right dark class="dark-green darken-4">
            <v-list dense>

                <v-subheader>Contents</v-subheader>
                <v-list-item link to="/">
                    <v-list-item-action>
                        <v-icon>mdi-home</v-icon>
                    </v-list-item-action>

                    <v-list-item-content>
                        <v-list-item-title>Top</v-list-item-title>
                        <v-list-item-subtitle>Back to blog top.</v-list-item-subtitle>
                    </v-list-item-content>
                </v-list-item>

<!--                <v-list-item link to="/portfolio">-->
<!--                    <v-list-item-action>-->
<!--                        <v-icon>mdi-notebook-outline</v-icon>-->
<!--                    </v-list-item-action>-->
<!--                    <v-list-item-content>-->
<!--                        <v-list-item-title>Portfolio</v-list-item-title>-->
<!--                        <v-list-item-subtitle>About works.</v-list-item-subtitle>-->
<!--                    </v-list-item-content>-->
<!--                </v-list-item>-->

                <v-divider></v-divider>
                <v-subheader>Links</v-subheader>

                <v-list-item link href="https://qiita.com/kimihiro_n" target="_blank" rel="noopener">
                    <v-list-item-action>
                        <v-icon>mdi-post</v-icon>
                    </v-list-item-action>
                    <v-list-item-content>
                        <v-list-item-title>Qiita</v-list-item-title>
                    </v-list-item-content>
                </v-list-item>

                <v-list-item link href="https://scrapbox.io/kimihiro-n/" target="_blank" rel="noopener">
                    <v-list-item-action>
                        <v-icon>mdi-note</v-icon>
                    </v-list-item-action>
                    <v-list-item-content>
                        <v-list-item-title>Scrapbox</v-list-item-title>
                    </v-list-item-content>
                </v-list-item>

                <v-list-item link href="https://play.google.com/store/apps/dev?id=8846644054781474049&hl=ja"
                             target="_blank" rel="noopener">
                    <v-list-item-action>
                        <v-icon>mdi-android</v-icon>
                    </v-list-item-action>
                    <v-list-item-content>
                        <v-list-item-title>Android Apps</v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
                <v-subheader>Author</v-subheader>
                <v-list-item>
                    <v-list-item-avatar>
                        <v-img src="/img/profile.png"
                               style="border-radius: 32px; border: #ffffff 1px solid;background: #ffffff"></v-img>
                    </v-list-item-avatar>
                    <v-list-item-content>
                        <v-list-item-title>kimhiro_n</v-list-item-title>
                        <v-list-item-subtitle>
                            <a href="https://twitter.com/kimihiro_n" target="_blank" rel="noopener">Twitter</a> /
                            <a href="https://github.com/pistatium" target="_blank" rel="noopener">GitHub</a>
                           </v-list-item-subtitle>

                    </v-list-item-content>
                </v-list-item>
            </v-list>
        </v-navigation-drawer>

        <v-app-bar
                app
                color="light-green darken-1"
                dark
        >

            <v-toolbar-title href="/">
                <router-link tag="h1" to="/">
                    <span v-if="! this.$root.isLogin">Pistatium blog</span>
                    <span v-if="this.$root.isLogin">Note</span>
                </router-link>
            </v-toolbar-title>

            <v-spacer/>
            Menu
            <v-app-bar-nav-icon @click.stop="drawer = !drawer"/>
        </v-app-bar>
    </div>
</template>

<script>
    export default {
        name: "Navigation",
        data: () => ({
            drawer: null,
        }),
    }
</script>

<style scoped>
    h1 {
        cursor: pointer;
        font-size: 24px;
    }

    .navigation {
        margin-bottom: 16px;
    }

    .site-icon {
        float: left;
        margin-right: 6px;
    }

</style>
